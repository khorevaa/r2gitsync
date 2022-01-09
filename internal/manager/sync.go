package manager

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/khorevaa/r2gitsync/internal/log"
	"github.com/khorevaa/r2gitsync/internal/manager/types"
	v8 "github.com/v8platform/api"
	"github.com/v8platform/designer"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"time"
)

type syncJob struct {
	name     string
	workdir  string
	tempDir  string
	repo     *designer.Repository
	infobase *v8.Infobase
	versions types.RepositoryVersionsList
	authors  types.RepositoryAuthorsList

	increment      bool
	currentVersion int
	minVersion     int
	maxVersion     int

	flow     Flow
	log      log.Logger
	endpoint types.V8Endpoint
	domainEmail string

	limitVersion int
	options      []interface{}

	skipFiles []string
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
		extention:  j.repo.Extension,
	}

	taskFlow := j.flow

	taskFlow.StartSyncProcess(j.endpoint, j.workdir)
	defer taskFlow.FinishSyncProcess(j.endpoint, j.workdir, &err)

	j.currentVersion, err = taskFlow.ReadVersionFile(j.endpoint, j.workdir, VERSION_FILE)

	if err != nil {
		return err
	}

	err = taskFlow.GetRepositoryAuthors(j.endpoint, j.workdir, AUTHORS_FILE, &j.authors)

	if err != nil {
		return err
	}

	err = taskFlow.GetRepositoryVersions(j.endpoint, j.workdir, j.currentVersion, j.maxVersion, &j.versions)
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
		zap.Int("currentVersion", j.currentVersion),
		zap.Int("maxVersion", j.maxVersion),
		zap.Int("nextVersion", nextVersion),
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

func (j *syncJob) syncVersionFiles(rVersion types.RepositoryVersion) (err error) {

	tempDir, err := ioutil.TempDir(j.tempDir, fmt.Sprintf("v%d", rVersion.Number()))

	if err != nil {
		return err
	}

	flowTask := j.flow

	j.log.Infow(fmt.Sprintf("Start process sources for version %d", rVersion.Number()),
		zap.Int("number", rVersion.Number()))
	startTime := time.Now()

	flowTask.StartSyncVersion(j.endpoint, j.workdir, tempDir, rVersion.Number())

	defer func() {

		flowTask.FinishSyncVersion(j.endpoint, j.workdir, tempDir, rVersion.Number(), &err)

		_ = os.RemoveAll(tempDir)

		j.log.Infow(fmt.Sprintf("Finished process sources for version %d", rVersion.Number()),
			zap.Int("number", rVersion.Number()),
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
		err = flowTask.ClearWorkDir(j.endpoint, j.workdir, tempDir, j.skipFiles)

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

func (j *syncJob) getRepositoryAuthor(name string) types.RepositoryAuthor {

	author, ok := j.authors[name]

	if !ok {

		author = NewAuthor(name, fmt.Sprintf("%s@%s", name, j.domainEmail))

		j.authors[name] = author
	}

	return author

}
