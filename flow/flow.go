package flow

import (
	"github.com/khorevaa/r2gitsync/plugin/Subscription"
	"github.com/v8platform/designer/repository"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	v8 "github.com/v8platform/v8"
	"time"
)

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *repository.Repository
	Extention() string
	Options() []interface{}
}

type RepositoryAuthor interface {
	Name() string
	Email() string
	Desc() string
}

type RepositoryVersion interface {
	Version() string
	Author() string
	Date() time.Time
	Comment() string
	Number() int64
}

type Flow interface {
	StartSyncVersions(v8end V8Endpoint, list []RepositoryVersion, currentVersion int64, nextVersion *int64, maxVersion *int64)

	StartSyncVersion(v8end V8Endpoint, workdir string, tempdir string, number int64) error
	FinishSyncVersion(v8end V8Endpoint, workdir string, tempdir string, number int64, err *error)

	StartSyncProcess(v8end V8Endpoint, dir string)
	FinishSyncProcess(v8end V8Endpoint, dir string)

	UpdateCfg(v8end V8Endpoint, workDir string, number int64) (err error)
	DumpConfigToFiles(v8end V8Endpoint, dir string, dir2 string, number int64) error

	ClearWorkDir(v8end V8Endpoint, dir string, tempDir string) error
	MoveToWorkDir(v8end V8Endpoint, dir string, tempDir string) error
	WriteVersionFile(v8end V8Endpoint, dir string, number int64) error
	CommitFiles(v8end V8Endpoint, dir string, author string, d time.Time, comment string) error

	GetRepositoryVersions(v8end V8Endpoint, dir string) ([]RepositoryVersion, error)
	GetRepositoryAuthors(v8end V8Endpoint, dir string) ([]RepositoryAuthor, error)
}

func Tasker() Flow {
	return tasker{}
}

func WithSubscribes(sm *Subscription.SubscribeManager) Flow {
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
