package flow

import (
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/manager/types"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"time"
)

var _ Flow = (*subscribeTasker)(nil)

type subscribeTasker struct {
	tasker
	log log.Logger
	pm  *subscription.SubscribeManager
}

func (t subscribeTasker) StartSyncVersion(v8end types.V8Endpoint, workdir string, tempdir string, number int64) {
	t.pm.SyncVersion.Start(v8end, workdir, tempdir, number)
}

func (t subscribeTasker) FinishSyncVersion(v8end types.V8Endpoint, workdir string, tempdir string, number int64, err *error) {
	t.pm.SyncVersion.Finish(v8end, workdir, tempdir, number, err)
}

func (t subscribeTasker) StartSyncProcess(v8end types.V8Endpoint, dir string) {
	t.pm.SyncProcess.Start(v8end, dir)
}

func (t subscribeTasker) FinishSyncProcess(v8end types.V8Endpoint, dir string, err *error) {
	t.pm.SyncProcess.Finish(v8end, dir, err)
}

func (t subscribeTasker) DumpConfigToFiles(v8end types.V8Endpoint, dir string, temp string, number int64, update bool) error {

	h := t.pm.DumpConfigToFiles

	err := h.Before(v8end, dir, temp, number, &update)

	if err != nil {
		return err
	}

	stdHandler := true

	err = h.On(v8end, dir, temp, number, &update, &stdHandler)

	if err != nil {
		return err
	}

	if stdHandler {
		err = t.tasker.DumpConfigToFiles(v8end, dir, temp, number, update)
	}

	err = h.After(v8end, dir, temp, number, update)

	return err
}

func (t subscribeTasker) ClearWorkDir(v8end types.V8Endpoint, dir string, temp string) error {
	h := t.pm.ClearWorkdir

	err := h.Before(v8end, dir, temp)

	if err != nil {
		return err
	}

	stdHandler := true

	err = h.On(v8end, dir, temp, &stdHandler)

	if err != nil {
		return err
	}

	if stdHandler {
		err = t.tasker.ClearWorkDir(v8end, dir, temp)
	}

	err = h.After(v8end, dir, temp)
	return err
}

func (t subscribeTasker) MoveToWorkDir(v8end types.V8Endpoint, dir string, temp string) error {
	h := t.pm.MoveToWorkdir

	err := h.Before(v8end, dir, temp)

	if err != nil {
		return err
	}

	stdHandler := true

	err = h.On(v8end, dir, temp, &stdHandler)

	if err != nil {
		return err
	}

	if stdHandler {
		err = t.tasker.MoveToWorkDir(v8end, dir, temp)
	}

	err = h.After(v8end, dir, temp)

	return err
}

func (t subscribeTasker) WriteVersionFile(v8end types.V8Endpoint, dir string, number int64, filename string) error {

	h := t.pm.WriteVersionFile

	err := h.Before(v8end, dir, number, filename)

	if err != nil {
		return err
	}

	stdHandler := true

	err = h.On(v8end, dir, number, filename, &stdHandler)

	if err != nil {
		return err
	}

	if stdHandler {
		err = t.tasker.WriteVersionFile(v8end, dir, number, filename)
	}

	err = h.After(v8end, dir, number, filename)

	return err

}

func (t subscribeTasker) ReadVersionFile(v8end types.V8Endpoint, dir string, filename string) (int64, error) {
	h := t.pm.ReadVersionFile

	err := h.Before(v8end, dir, filename)

	if err != nil {
		return 0, err
	}

	stdHandler := true

	number, err := h.On(v8end, dir, filename, &stdHandler)

	if err != nil {
		return 0, err
	}

	if stdHandler {
		number, err = t.tasker.ReadVersionFile(v8end, dir, filename)
	}

	err = h.After(v8end, dir, filename, &number)

	if err != nil {
		return 0, err
	}

	return number, nil
}

func (t subscribeTasker) CommitFiles(v8end types.V8Endpoint, dir string, author types.RepositoryAuthor, when time.Time, comment string) error {

	h := t.pm.CommitFiles

	err := h.Before(v8end, dir, author, when, comment)

	if err != nil {
		return err
	}

	stdHandler := true

	err = h.On(v8end, dir, author, &when, &comment, &stdHandler)

	if err != nil {
		return err
	}

	if stdHandler {
		err = t.tasker.CommitFiles(v8end, dir, author, when, comment)
	}

	err = h.After(v8end, dir, author, when, comment)

	return err
}

func (t subscribeTasker) GetRepositoryVersions(v8end types.V8Endpoint, dir string, NBegin int64) (rv []types.RepositoryVersion, err error) {
	h := t.pm.GetRepositoryHistory

	err = h.Before(v8end, dir, NBegin)

	if err != nil {
		return
	}

	stdHandler := true

	rv, err = h.On(v8end, dir, NBegin, &stdHandler)

	if err != nil {
		return
	}

	if stdHandler {
		rv, err = t.tasker.GetRepositoryVersions(v8end, dir, NBegin)
	}

	err = h.After(v8end, dir, NBegin, &rv)

	return
}

func (t subscribeTasker) ConfigureRepositoryVersions(v8end types.V8Endpoint, versions *[]types.RepositoryVersion, NCurrent, NNext, NMax *int64) (err error) {
	h := t.pm.ConfigureRepositoryVersions

	err = h.On(v8end, versions, NCurrent, NNext, NMax)

	if err != nil {
		return
	}

	return
}

func (t subscribeTasker) GetRepositoryAuthors(v8end types.V8Endpoint, dir string, filenme string) (ra map[string]types.RepositoryAuthor, err error) {
	h := t.pm.GetRepositoryAuthors

	err = h.Before(v8end, dir, filenme)

	if err != nil {
		return
	}

	stdHandler := true

	ra, err = h.On(v8end, dir, filenme, &stdHandler)

	if err != nil {
		return
	}

	if stdHandler {
		ra, err = t.tasker.GetRepositoryAuthors(v8end, dir, filenme)
	}

	err = h.After(v8end, dir, &ra)

	return
}

func (t subscribeTasker) UpdateCfg(v8end types.V8Endpoint, workDir string, number int64) (err error) {

	UpdateCfg := t.pm.UpdateCfg

	err = UpdateCfg.Before(v8end, workDir, number)

	if err != nil {
		return
	}

	stdHandler := true

	err = UpdateCfg.On(v8end, workDir, number, &stdHandler)

	if err != nil {
		return
	}

	if stdHandler {
		err = t.tasker.UpdateCfg(v8end, workDir, number)
	}

	err = UpdateCfg.After(v8end, workDir, number)

	return nil
}
