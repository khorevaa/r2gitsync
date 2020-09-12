package flow

import (
	"fmt"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"github.com/v8platform/designer"
	"github.com/v8platform/designer/repository"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var _ Flow = (*tasker)(nil)

const ConfigDumpInfoFileName = "ConfigDumpInfo.xml"

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
	}.
		GroupByComment().
		WithRepository(*v8end.Repository())

	err = run(*v8end.Infobase(), RepositoryReportOptions, v8end.Options()...)

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

func (t tasker) GetRepositoryAuthors(v8end V8Endpoint, dir string, filename string) (authors map[string]RepositoryAuthor, err error) {

	authors = make(map[string]RepositoryAuthor)

	file := path.Join(dir, filename)
	if ok, _ := Exists(file); !ok {

		return

	}

	bytesRead, _ := ioutil.ReadFile(file)
	file_content := string(bytesRead)
	lines := strings.Split(file_content, "\n")

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if len(line) == 0 || strings.HasPrefix(line, "//") {
			continue
		}

		data := strings.Split(line, "=")

		authors[data[0]] = NewAuthor(decodeAuthor([]byte(data[1])))

	}

	return

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

func (t tasker) DumpConfigToFiles(endpoint V8Endpoint, update bool, dir string, tempdir string, number int64) error {

	var configDumpInfoFile = ""

	if update {
		configDumpInfoFile = path.Join(dir, ConfigDumpInfoFileName)

		if ok, _ := Exists(configDumpInfoFile); !ok {
			update = false
			configDumpInfoFile = ""
		}
	}

	DumpConfigToFilesOptions := designer.DumpConfigToFilesOptions{
		Dir:                      tempdir,
		Force:                    true,
		Update:                   update,
		Extension:                endpoint.Extention(),
		ConfigDumpInfoForChanges: configDumpInfoFile,
	}

	err := run(*endpoint.Infobase(), DumpConfigToFilesOptions, endpoint.Options())

	if err != nil {
		return err
	}
	return nil

}

func (t tasker) FinishSyncProcess(endpoint V8Endpoint, dir string) {
	return
}

func (t tasker) ClearWorkDir(endpoint V8Endpoint, dir string, tempDir string) error {

	err := clearDir(dir, "VERSION", "AUTHORS", ".git") // TODO Сделать копирование файлов или избранную очистку

	if err != nil {
		return err
	}
	return nil
}

func (t tasker) MoveToWorkDir(endpoint V8Endpoint, dir string, tempDir string) error {

	err := CopyDir(tempDir, dir)

	return err
}

func (t tasker) WriteVersionFile(endpoint V8Endpoint, dir string, number int64, versionFile string) error {

	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<VERSION>%d</VERSION>`, number)

	filename := filepath.Join(dir, versionFile)
	err := ioutil.WriteFile(filename, []byte(data), 0644)

	return err
}

func (t tasker) CommitFiles(endpoint V8Endpoint, dir string, author RepositoryAuthor, date time.Time, comment string) error {

	err := CommitFiles(dir, author, date, comment)

	return err
}

func (t tasker) WithSubscribes(sm *subscription.SubscribeManager) Flow {

	return subscribeTasker{
		tasker: t,
		pm:     sm,
	}

}
