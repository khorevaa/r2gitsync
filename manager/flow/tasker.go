package flow

import (
	"github.com/khorevaa/r2gitsync/plugin/Subscription"
	"github.com/v8platform/designer/repository"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

var _ Flow = (*tasker)(nil)

type tasker struct {
}

func (t tasker) ConfigureRepositoryVersions(v8end V8Endpoint, versions []RepositoryVersion, NBegin, NNext, NMax *int64) (err error) {

	if len(versions) > 0 {
		maxVersion := versions[len(versions)-1].Number()
		*NMax = maxVersion
	}

	return
}

func (t tasker) StartSyncVersion(v8end V8Endpoint, workdir string, tempdir string, number int64) error {
	return nil
}

func (t tasker) StartSyncProcess(v8end V8Endpoint, dir string) {
	return
}

func (t tasker) GetRepositoryVersions(v8end V8Endpoint, dir string, nBegin int64) (versions []RepositoryVersion, err error) {

	reportFile, err := ioutil.TempFile(os.TempDir(), "v8_rep_history")
	if err != nil {
		return
	}
	reportFile.Close()
	report := reportFile.Name()

	defer os.Remove(report)

	RepositoryReportOptions := repository.RepositoryReportOptions{
		File:      report,
		Extension: v8end.Extention(),
		NBegin:    nBegin,
	}.GroupByComment().WithRepository(*v8end.Repository())

	err = run(*v8end.Infobase(), RepositoryReportOptions, v8end.Options())

	if err != nil {
		return
	}

	versions, err = parseRepositoryReport(report)

	if err != nil {
		return
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Number() < versions[j].Number()
	})

	//if len(r.Versions) > 0 {
	//	r.MaxVersion = r.Versions[len(r.Versions)-1].Number()
	//}

	return
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
