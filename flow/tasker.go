package flow

import (
	"github.com/khorevaa/r2gitsync/plugin/Subscription"
	"github.com/v8platform/designer/repository"
	"time"
)

type tasker struct {
}

func (t tasker) StartSyncVersions(v8end V8Endpoint, list []RepositoryVersion, currentVersion int64, nextVersion *int64, maxVersion *int64) {
	return
}

func (t tasker) StartSyncVersion(v8end V8Endpoint, workdir string, tempdir string, number int64) error {
	return nil
}

func (t tasker) StartSyncProcess(v8end V8Endpoint, dir string) {
	return
}

func (t tasker) GetRepositoryVersions(v8end V8Endpoint, dir string) ([]RepositoryVersion, error) {

	// TODO

	return nil, nil
}

func (t tasker) GetRepositoryAuthors(v8end V8Endpoint, dir string) ([]RepositoryAuthor, error) {
	// TODO

	return nil, nil
}

func (t tasker) UpdateCfg(v8end V8Endpoint, workDir string, number int64) (err error) {

	RepositoryUpdateCfgOptions := repository.RepositoryUpdateCfgOptions{
		Version:   number,
		Force:     true,
		Extension: v8end.Extention(),
	}.WithRepository(*v8end.Repository())

	err = run(*v8end.Infobase(), RepositoryUpdateCfgOptions, v8end.Options()...)

	return
}

func (t tasker) FinishSyncVersion(endpoint V8Endpoint, workdir string, tempdir string, number int64, err *error) {
	return
}

func (t tasker) DumpConfigToFiles(endpoint V8Endpoint, dir string, dir2 string, number int64) error {
	//TODO
	return nil
}

func (t tasker) FinishSyncProcess(endpoint V8Endpoint, dir string) {
	return
}

func (t tasker) ClearWorkDir(endpoint V8Endpoint, dir string, dir2 string) error {

	//TODO
	return nil
}

func (t tasker) MoveToWorkDir(endpoint V8Endpoint, dir string, dir2 string) error {
	//TODO
	return nil
}

func (t tasker) WriteVersionFile(endpoint V8Endpoint, dir string, number int64) error {
	//TODO
	return nil
}

func (t tasker) CommitFiles(endpoint V8Endpoint, dir string, a string, d time.Time, comment string) error {
	//TODO
	return nil
}

func (t tasker) WithSubscribes(sm *Subscription.SubscribeManager) Flow {

	return subscribeTasker{
		tasker: t,
		pm:     sm,
	}

}
