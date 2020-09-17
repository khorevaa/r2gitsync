package flow

import (
	"github.com/khorevaa/r2gitsync/manager/types"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	v8 "github.com/v8platform/v8"
	"time"
)

type RepositoryVersions []types.RepositoryVersion

type Flow interface {
	StartSyncVersion(v8end types.V8Endpoint, workdir string, tempdir string, number int64)
	FinishSyncVersion(v8end types.V8Endpoint, workdir string, tempdir string, number int64, err *error)

	StartSyncProcess(v8end types.V8Endpoint, dir string)
	FinishSyncProcess(v8end types.V8Endpoint, dir string, err *error)

	UpdateCfg(v8end types.V8Endpoint, workDir string, number int64) (err error)
	DumpConfigToFiles(v8end types.V8Endpoint, dir string, temp string, number int64, update bool) error

	ClearWorkDir(v8end types.V8Endpoint, dir string, tempDir string) error
	MoveToWorkDir(v8end types.V8Endpoint, dir string, tempDir string) error
	WriteVersionFile(v8end types.V8Endpoint, dir string, number int64, filename string) error
	ReadVersionFile(end types.V8Endpoint, dir string, filename string) (int64, error)
	CommitFiles(v8end types.V8Endpoint, dir string, author types.RepositoryAuthor, when time.Time, comment string) error

	GetRepositoryVersions(v8end types.V8Endpoint, dir string, NBegin int64) ([]types.RepositoryVersion, error)
	ConfigureRepositoryVersions(v8end types.V8Endpoint, versions *[]types.RepositoryVersion, NCurrent, NNext, NMax *int64) (err error)
	GetRepositoryAuthors(v8end types.V8Endpoint, dir string, filenme string) (map[string]types.RepositoryAuthor, error)
}

func Tasker() Flow {
	return tasker{}
}

func WithSubscribes(sm *subscription.SubscribeManager) Flow {
	return tasker{}.WithSubscribes(sm)
}

func run(where runner.Infobase, what runner.Command, opts ...interface{}) error {

	err := v8.Run(where, what, opts...,
	//	v8.WithTempDir(opts.tempDir), // TODO Сделать для запуска временный катиалог
	)

	errorContext := errors.GetErrorContext(err)

	out, ok := errorContext["message"]
	if ok {
		err = errors.Internal.Wrap(err, out)
	}

	//TODO Сделать несколько попыток при отсутсвии лицензиии

	return err

}
