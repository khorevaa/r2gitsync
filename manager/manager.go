package manager

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/manager/flow"
	"github.com/khorevaa/r2gitsync/manager/types"
	"github.com/lithammer/shortuuid/v3"
	"github.com/v8platform/designer/repository"
	"github.com/v8platform/v8"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type SyncRepository struct {
	repository.Repository
	Name     string
	Workdir  string
	Versions []types.RepositoryVersion
	Authors  map[string]types.RepositoryAuthor

	Extention      string
	increment      bool
	CurrentVersion int64 `xml:"VERSION"`
	MinVersion     int64
	MaxVersion     int64
	endpoint       types.V8Endpoint
	flow           flow.Flow
	log            log.Logger
}

func (r *SyncRepository) Auth(user, passowrd string) {

	r.User = user
	r.Password = passowrd

}

func (r *SyncRepository) ReadCurrentVersion() (err error) {

	r.CurrentVersion, err = r.flow.ReadVersionFile(r.endpoint, r.Workdir, VERSION_FILE)

	return
}

func (r *SyncRepository) sync(opts *Options) (err error) {

	if opts.logger != nil {
		r.log = opts.logger.Named("manager")
	}
	if r.log == nil {
		r.log = log.NewLogger().Named("manager")
	}

	if len(r.Name) == 0 {
		r.Name = shortuuid.New()
	}

	r.log.Infow("Start sync with repository",
		zap.String("name", r.Name),
		zap.String("path", r.Repository.Path),
	)

	r.endpoint = &v8Endpoint{
		infobase:   &opts.infobase,
		repository: &r.Repository,
		options:    opts.Options(),
		extention:  r.Extention,
	}

	r.log.Infow("Using infobase for sync",
		zap.String("path", opts.infobase.ConnectionString()))

	r.flow = flow.Tasker(r.log)

	if opts.plugins != nil {

		r.log.Info("Using plugins for sync")
		r.flow = flow.WithSubscribes(r.log.With(zap.String("tasker", "subscriber")), opts.plugins)

	}

	r.increment = !opts.disableIncrement
	r.log.Infow("Using increment dump config to files", zap.Bool("increment", r.increment))

	taskFlow := r.flow

	taskFlow.StartSyncProcess(r.endpoint, r.Workdir)
	defer taskFlow.FinishSyncProcess(r.endpoint, r.Workdir, &err)

	err = r.prepare(opts)

	if err != nil {
		return err
	}

	if len(r.Versions) == 0 {
		r.log.Warn("No versions to sync")
		return nil
	}

	nextVersion := r.Versions[0].Number()
	err = taskFlow.ConfigureRepositoryVersions(r.endpoint, &r.Versions, &r.CurrentVersion, &nextVersion, &r.MaxVersion)

	r.log.Infow("Sync version number",
		zap.Int64("currentVersion", r.CurrentVersion),
		zap.Int64("mexVersion", r.MaxVersion),
		zap.Int64("nextVersion", nextVersion),
	)

	if err != nil {
		return err
	}

	for _, rVersion := range r.Versions {

		if r.MaxVersion != 0 &&
			rVersion.Number() > r.MaxVersion {
			break
		}

		if nextVersion > rVersion.Number() {
			continue
		}

		err = r.syncVersionFiles(rVersion, opts)

		if err != nil {
			return err
		}

	}

	return nil
}

func (r *SyncRepository) WriteVersionFile(CurrentVersion int64) error {

	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<VERSION>%d</VERSION>`, CurrentVersion)

	filename := filepath.Join(r.Workdir, VERSION_FILE)
	err := ioutil.WriteFile(filename, []byte(data), 0644)

	return err

}

func (r *SyncRepository) prepare(opts *Options) error {

	if !opts.infobaseCreated {

		CreateFileInfobase := v8.CreateFileInfobase(opts.infobase.Path())

		err := Run(opts.infobase, CreateFileInfobase, opts)

		if err != nil {
			return err
		}

		opts.infobaseCreated = true
	}

	r.CurrentVersion = 0

	err := r.ReadCurrentVersion()

	if err != nil {
		return err
	}

	err = r.GetRepositoryAuthors(opts)

	if err != nil {
		return err
	}

	r.MaxVersion = 0

	err = r.GetRepositoryHistory(opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *SyncRepository) Sync(opts ...Option) error {

	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	return r.sync(options)

}

func Sync(r SyncRepository, opts ...Option) error {

	return r.Sync(opts...)

}

func (r *SyncRepository) syncVersionFiles(rVersion types.RepositoryVersion, opts *Options) (err error) {

	tempDir, err := ioutil.TempDir(opts.tempDir, fmt.Sprintf("v%d", rVersion.Number()))

	if err != nil {
		return err
	}

	flowTask := r.flow

	r.log.Infow(fmt.Sprintf("Start process sources for version %d", rVersion.Number()),
		zap.Int64("number", rVersion.Number()))
	startTime := time.Now()

	flowTask.StartSyncVersion(r.endpoint, r.Workdir, tempDir, rVersion.Number())

	defer func() {

		flowTask.FinishSyncVersion(r.endpoint, r.Workdir, tempDir, rVersion.Number(), &err)

		_ = os.RemoveAll(tempDir)

		r.log.Infow(fmt.Sprintf("Finished process sources for version %d", rVersion.Number()),
			zap.Int64("number", rVersion.Number()),
			zap.Float64("duration", time.Since(startTime).Seconds()),
			zap.Error(err))
	}()

	err = flowTask.UpdateCfg(r.endpoint, r.Workdir, rVersion.Number())

	if err != nil {
		return err
	}

	err = flowTask.DumpConfigToFiles(r.endpoint, r.Workdir, tempDir, rVersion.Number(), r.increment)

	if err != nil {
		return err
	}

	err = flowTask.ClearWorkDir(r.endpoint, r.Workdir, tempDir)

	if err != nil {
		return err
	}

	err = flowTask.MoveToWorkDir(r.endpoint, r.Workdir, tempDir)

	if err != nil {
		return err
	}

	err = flowTask.WriteVersionFile(r.endpoint, r.Workdir, rVersion.Number(), VERSION_FILE)

	if err != nil {
		return err
	}

	err = flowTask.CommitFiles(r.endpoint, r.Workdir, r.getRepositoryAuthor(rVersion.Author(), opts), rVersion.Date(), rVersion.Comment())

	if err != nil {

		errV := flowTask.WriteVersionFile(r.endpoint, r.Workdir, rVersion.Number(), VERSION_FILE)
		if errV != nil {
			return multierror.Append(err, errV)
		}
		return err
	}

	r.CurrentVersion = rVersion.Number()

	return

}

func (r SyncRepository) getRepositoryAuthor(name string, opts *Options) types.RepositoryAuthor {

	author, ok := r.Authors[name]

	if !ok {

		author = flow.NewAuthor(name, fmt.Sprintf("%s@%s", name, opts.DomainEmail()))

		r.Authors[name] = author
	}

	return author

}

func (r *SyncRepository) GetRepositoryHistory(opts *Options) (err error) {

	r.Versions, err = r.flow.GetRepositoryVersions(r.endpoint, r.Workdir, r.CurrentVersion)

	return

}

func (r *SyncRepository) GetRepositoryAuthors(opts *Options) (err error) {

	r.Authors, _ = r.flow.GetRepositoryAuthors(r.endpoint, r.Workdir, AUTHORS_FILE)

	return

}
