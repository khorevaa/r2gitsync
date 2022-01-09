package manager

import (
	"github.com/khorevaa/r2gitsync/internal/log"
	"github.com/khorevaa/r2gitsync/internal/manager/types"
	"github.com/khorevaa/r2gitsync/pkg/plugin/subscription"
	"time"
)

const ConfigDumpInfoFileName = "ConfigDumpInfo.xml"

type Flow struct {
	SubscribeManager *subscription.SubscribeManager
	Logger           log.Logger
}

func (t Flow) StartSyncVersion(v8end types.V8Endpoint, workdir string, tempdir string, number int) {
	_ = t.Task(StartSyncVersion{v8end, workdir, tempdir, number})
}

func (t Flow) FinishSyncVersion(v8end types.V8Endpoint, workdir string, tempdir string, number int, err *error) {
	_ = t.Task(FinishSyncVersion{v8end, workdir, tempdir, number, err})
}

func (t Flow) StartSyncProcess(v8end types.V8Endpoint, dir string) {
	_ = t.Task(StartSyncProcess{v8end, dir})
}

func (t Flow) FinishSyncProcess(v8end types.V8Endpoint, dir string, err *error) {
	_ = t.Task(FinishSyncProcess{v8end, dir, err})
}

func (t Flow) DumpConfigToFiles(v8end types.V8Endpoint, dir string, temp string, number int, update bool) (bool, error) {

	isIncremented := new(bool)

	err := t.Task(DumpConfigToFiles{v8end, dir,
		temp, number,
		update, isIncremented})
	return *isIncremented, err
}

func (t Flow) ClearWorkDir(v8end types.V8Endpoint, dir string, temp string, skipFiles []string) error {

	return t.Task(ClearWorkdir{v8end, dir,
		temp, skipFiles})
}

func (t Flow) MoveToWorkDir(v8end types.V8Endpoint, dir string, temp string) error {

	return t.Task(MoveToWorkdir{v8end, dir, temp})
}

func (t Flow) WriteVersionFile(v8end types.V8Endpoint, dir string, number int, filename string) error {

	return t.Task(WriteVersionFile{v8end, dir, filename, number})

}

func (t Flow) ReadVersionFile(v8end types.V8Endpoint, dir string, filename string) (int, error) {

	number := new(int)

	err := t.Task(ReadVersionFile{v8end, dir,
		filename, number})
	return *number, err

}

func (t Flow) CommitFiles(v8end types.V8Endpoint, dir string, author types.RepositoryAuthor, when time.Time, comment string) error {

	return t.Task(CommitFiles{v8end, dir, author, when, comment})

}

func (t Flow) GetRepositoryVersions(v8end types.V8Endpoint, dir string, NBegin, NEnd int, list *types.RepositoryVersionsList) error {

	err := t.Task(GetRepositoryVersions{
		V8end:    v8end,
		Workdir:  dir,
		NBegin:   NBegin,
		NEnd:     NEnd,
		Versions: list,
	})

	return err
}

func (t Flow) ConfigureRepositoryVersions(v8end types.V8Endpoint, versions *types.RepositoryVersionsList, NCurrent, NNext, NMax *int) (err error) {

	return t.Task(ConfigureRepositoryVersions{v8end, versions,
		NCurrent, NNext, NMax})
}

func (t Flow) GetRepositoryAuthors(v8end types.V8Endpoint, dir string, filename string, authors *types.RepositoryAuthorsList) error {

	return t.Task(GetRepositoryAuthors{v8end, dir, filename, authors})
}

func (t Flow) UpdateCfg(v8end types.V8Endpoint, workDir string, number int) error {

	return t.Task(UpdateCfg{v8end, workDir, number})
}

func (f Flow) Task(task Task) error {
	return DoTask(task, f.SubscribeManager)
}

type Task interface {
	Action(useStdHandler bool) error
	Before(pm *subscription.SubscribeManager) error
	On(pm *subscription.SubscribeManager, useStdHandler *bool) error
	After(pm *subscription.SubscribeManager) error
}

func DoTask(t Task, subscribeManager *subscription.SubscribeManager) (err error) {

	stdHandler := true

	if subscribeManager == nil {
		return t.Action(true)
	}

	if err = t.Before(subscribeManager); err != nil {
		return err
	}

	if err = t.On(subscribeManager, &stdHandler); err != nil {
		return err
	}

	if stdHandler {
		return t.Action(stdHandler)
	}

	if err = t.After(subscribeManager); err != nil {
		return err
	}

	return

}
