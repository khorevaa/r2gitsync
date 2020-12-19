package manager

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/manager/flow"
	"github.com/khorevaa/r2gitsync/manager/types"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"github.com/lithammer/shortuuid/v3"
	v8 "github.com/v8platform/api"
	"github.com/v8platform/designer"
	"github.com/v8platform/runner"
	"go.uber.org/zap"
	"io/ioutil"
	"os"

	"time"
)

func Sync(r SyncRepository, options Options) error {

	return r.Sync(options)

}

func Init(r SyncRepository, opts Options) error {

	//return r.Init(opts)

	return nil
}

type syncJob struct {
	name     string
	workdir  string
	tempDir  string
	repo     *designer.Repository
	infobase *v8.Infobase
	versions []types.RepositoryVersion
	authors  map[string]types.RepositoryAuthor

	increment      bool
	currentVersion int64
	minVersion     int64
	maxVersion     int64

	flow        flow.Flow
	log         log.Logger
	endpoint    types.V8Endpoint
	domainEmail string

	jobs         []func()
	limitVersion int64
	options      []interface{}
}

func (j *syncJob) Run() error {

	return j.start()

}

func (j *syncJob) start() (err error) {

	j.log.Infow("Start sync with repository",
		zap.String("name", j.name),
		zap.String("path", j.repo.Path),
	)

	j.log.Infow("Using infobase for sync",
		zap.String("path", j.infobase.Connect.String()))

	if j.increment {
		j.log.Infow("Using increment dump config to files")
	}

	j.endpoint = &v8Endpoint{
		infobase:   j.infobase,
		repository: j.repo,
		options:    j.options,
	}

	taskFlow := j.flow

	taskFlow.StartSyncProcess(j.endpoint, j.workdir)
	defer taskFlow.FinishSyncProcess(j.endpoint, j.workdir, &err)

	err = j.readCurrentVersion()

	if err != nil {
		return err
	}

	err = j.getRepositoryAuthors()

	if err != nil {
		return err
	}

	err = j.getRepositoryHistory()
	if err != nil {
		return err
	}

	if len(j.versions) == 0 {
		j.log.Warn("No versions to sync")
		return nil
	}

	nextVersion := j.versions[0].Number()
	err = taskFlow.ConfigureRepositoryVersions(j.endpoint, &j.versions, &j.currentVersion, &nextVersion, &j.maxVersion)

	j.log.Infow("Sync version number",
		zap.Int64("currentVersion", j.currentVersion),
		zap.Int64("maxVersion", j.maxVersion),
		zap.Int64("nextVersion", nextVersion),
	)

	if err != nil {
		return err
	}

	for _, rVersion := range j.versions {

		if j.maxVersion != 0 &&
			rVersion.Number() > j.maxVersion {
			break
		}

		if nextVersion > rVersion.Number() {
			continue
		}

		err = j.syncVersionFiles(rVersion)

		if err != nil {
			return err
		}

	}

	return nil

}

type SyncRepository struct {
	designer.Repository
	Name    string
	Workdir string
}

func (r *SyncRepository) Sync(options Options) error {

	jobLogger := log.Logger(log.NullLogger)

	if options.Logger != nil {
		jobLogger = options.Logger.Named("manager")
	}

	ib, err := r.getInfobase(jobLogger, options)
	if err != nil {
		return err
	}

	if len(r.Name) == 0 {
		r.Name = shortuuid.New()
	}

	jobName := fmt.Sprintf("Sync job for repository: %s (%s)", r.Name, r.Path)

	job := syncJob{
		name:         jobName,
		workdir:      r.Workdir,
		tempDir:      options.TempDir,
		repo:         &r.Repository,
		infobase:     ib,
		increment:    !options.DisableIncrement,
		minVersion:   options.MinVersion,
		maxVersion:   options.MaxVersion,
		limitVersion: options.LimitVersions,
		flow:         getTasker(jobLogger, options.Plugins),
		options:      options.Options(),
		log:          jobLogger,
		domainEmail:  options.DomainEmail,
	}

	return job.Run()

}

func (j *syncJob) readCurrentVersion() (err error) {

	j.currentVersion, err = j.flow.ReadVersionFile(j.endpoint, j.workdir, VERSION_FILE)

	return
}

func (j *syncJob) getRepositoryHistory() (err error) {

	j.versions, err = j.flow.GetRepositoryVersions(j.endpoint, j.workdir, j.currentVersion)

	return

}

func (j *syncJob) getRepositoryAuthors() (err error) {

	j.authors, err = j.flow.GetRepositoryAuthors(j.endpoint, j.workdir, AUTHORS_FILE)

	return
}

func (j *syncJob) syncVersionFiles(rVersion types.RepositoryVersion) (err error) {

	tempDir, err := ioutil.TempDir(j.tempDir, fmt.Sprintf("v%d", rVersion.Number()))

	if err != nil {
		return err
	}

	flowTask := j.flow

	j.log.Infow(fmt.Sprintf("Start process sources for version %d", rVersion.Number()),
		zap.Int64("number", rVersion.Number()))
	startTime := time.Now()

	flowTask.StartSyncVersion(j.endpoint, j.workdir, tempDir, rVersion.Number())

	defer func() {

		flowTask.FinishSyncVersion(j.endpoint, j.workdir, tempDir, rVersion.Number(), &err)

		_ = os.RemoveAll(tempDir)

		j.log.Infow(fmt.Sprintf("Finished process sources for version %d", rVersion.Number()),
			zap.Int64("number", rVersion.Number()),
			zap.Float64("duration", time.Since(startTime).Seconds()),
			zap.Error(err))
	}()

	err = flowTask.UpdateCfg(j.endpoint, j.workdir, rVersion.Number())

	if err != nil {
		return err
	}

	update, err := flowTask.DumpConfigToFiles(j.endpoint, j.workdir, tempDir, rVersion.Number(), j.increment)

	if err != nil {
		return err
	}

	if !update {
		err = flowTask.ClearWorkDir(j.endpoint, j.workdir, tempDir)

		if err != nil {
			return err
		}
	}

	err = flowTask.MoveToWorkDir(j.endpoint, j.workdir, tempDir)

	if err != nil {
		return err
	}

	err = flowTask.WriteVersionFile(j.endpoint, j.workdir, rVersion.Number(), VERSION_FILE)

	if err != nil {
		return err
	}

	err = flowTask.CommitFiles(j.endpoint, j.workdir,
		j.getRepositoryAuthor(rVersion.Author()), rVersion.Date(), rVersion.Comment())

	if err != nil {

		errV := flowTask.WriteVersionFile(j.endpoint, j.workdir, rVersion.Number(), VERSION_FILE)
		if errV != nil {
			return multierror.Append(err, errV)
		}
		return err
	}

	j.currentVersion = rVersion.Number()

	return

}

func getTasker(logger log.Logger, sm *subscription.SubscribeManager) flow.Flow {

	if sm != nil {

		return flow.WithSubscribes(logger.With(zap.String("tasker", "subscriber")), sm)

	}

	return flow.Tasker(logger)

}

func (r *SyncRepository) getInfobase(logger log.Logger, opts Options) (*v8.Infobase, error) {

	if len(opts.InfobaseConnect) > 0 {
		return opts.Infobase()
	}

	logger.Debug("Creating temp infobase")

	CreateFileInfobase := v8.CreateFileInfobase(v8.NewTempDir(opts.TempDir, "temp_ib"))

	ib, err := v8.CreateInfobase(CreateFileInfobase, opts.Options())

	if err != nil {
		return nil, err
	}

	if len(r.Extension) > 0 {

		tempExtension, err := restoreTempExtension()
		if err != nil {
			return nil, err
		}

		LoadExtensionCfg := v8.LoadExtensionCfg(tempExtension, r.Extension)
		err = v8.Run(ib, LoadExtensionCfg, opts.Options())

		if err != nil {
			return nil, err
		}
		logger.Debug("Empty extension loaded into infobase")
	}

	return ib, nil

}

func (j *syncJob) getRepositoryAuthor(name string) types.RepositoryAuthor {

	author, ok := j.authors[name]

	if !ok {

		author = flow.NewAuthor(name, fmt.Sprintf("%s@%s", name, j.domainEmail))

		j.authors[name] = author
	}

	return author

}

func restoreTempExtension() (string, error) {
	tempFile, err := ioutil.TempFile("", ".cfe")
	defer tempFile.Close()
	if err != nil {
		return "", err
	}

	bytes, err := Asset("tempExtension.cfe")
	_, err = tempFile.Write(bytes)

	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

func toInterface(options []runner.Option) []interface{} {

	var opts []interface{}

	for _, o := range options {

		opts = append(opts, o)

	}

	return opts

}
