package flow

import (
	"encoding/xml"
	"fmt"
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/manager/types"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"github.com/v8platform/designer"
	"github.com/v8platform/designer/repository"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var _ Flow = (*tasker)(nil)

const ConfigDumpInfoFileName = "ConfigDumpInfo.xml"

type tasker struct {
	log log.Logger
}

func (t tasker) ConfigureRepositoryVersions(v8end types.V8Endpoint, versions *[]types.RepositoryVersion, NBegin, NNext, NMax *int64) (err error) {

	ver := *versions

	if len(ver) > 0 {
		maxVersion := ver[len(ver)-1].Number()
		*NMax = maxVersion
	}

	versions = &ver

	return
}

func (t tasker) StartSyncVersion(v8end types.V8Endpoint, workdir string, tempdir string, number int64) {
	return
}

func (t tasker) StartSyncProcess(v8end types.V8Endpoint, dir string) {
	return
}

func (t tasker) GetRepositoryVersions(v8end types.V8Endpoint, dir string, nBegin int64) (versions []types.RepositoryVersion, err error) {

	t.log.Infow("Get repository versions",
		zap.String("path", dir),
		zap.Int64("nBegin", nBegin))

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

	err = Run(*v8end.Infobase(), RepositoryReportOptions, v8end.Options()...)

	if err != nil {
		return
	}

	t.log.Debugw("Parse repository report", zap.String("file", report))
	versions, err = parseRepositoryReport(report)

	if err != nil {
		return
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Number() < versions[j].Number()
	})

	return
}

func (t tasker) GetRepositoryAuthors(v8end types.V8Endpoint, dir string, filename string) (map[string]types.RepositoryAuthor, error) {

	authors := make(map[string]types.RepositoryAuthor)

	file := filepath.Join(dir, filename)
	if ok, _ := Exists(file); !ok {

		return authors, nil

	}

	bytesRead, _ := ioutil.ReadFile(file)
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if len(line) == 0 || strings.HasPrefix(line, "//") {
			continue
		}

		data := strings.Split(line, "=")

		authors[data[0]] = NewAuthor(decodeAuthor([]byte(data[1])))

	}

	return authors, nil

}

func (t tasker) UpdateCfg(v8end types.V8Endpoint, workDir string, number int64) (err error) {

	RepositoryUpdateCfgOptions := repository.RepositoryUpdateCfgOptions{
		Version:   number,
		Force:     true,
		Extension: v8end.Extention(),
	}.WithRepository(*v8end.Repository())

	err = Run(*v8end.Infobase(), RepositoryUpdateCfgOptions, v8end.Options()...)

	return
}

func (t tasker) FinishSyncVersion(endpoint types.V8Endpoint, workdir string, tempdir string, number int64, err *error) {
	return
}

func (t tasker) DumpConfigToFiles(endpoint types.V8Endpoint, dir string, tempdir string, number int64, update bool) error {

	var configDumpInfoFile = ""

	if update {
		configDumpInfoFile = filepath.Join(dir, ConfigDumpInfoFileName)

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

	err := Run(*endpoint.Infobase(), DumpConfigToFilesOptions, endpoint.Options())

	if err != nil {
		return err
	}
	return nil

}

func (t tasker) FinishSyncProcess(endpoint types.V8Endpoint, dir string, err *error) {
	return
}

func (t tasker) ClearWorkDir(endpoint types.V8Endpoint, dir string, tempDir string) error {

	err := clearDir(dir, "VERSION", "AUTHORS", ".git")
	if err != nil {
		return err
	}
	return nil
}

func (t tasker) MoveToWorkDir(endpoint types.V8Endpoint, dir string, tempDir string) error {

	err := CopyDir(tempDir, dir)

	return err
}

func (t tasker) WriteVersionFile(endpoint types.V8Endpoint, dir string, number int64, versionFile string) error {

	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<VERSION>%d</VERSION>`, number)

	filename := filepath.Join(dir, versionFile)
	err := ioutil.WriteFile(filename, []byte(data), 0644)

	return err
}

func (t tasker) CommitFiles(endpoint types.V8Endpoint, dir string, author types.RepositoryAuthor, date time.Time, comment string) error {

	err := CommitFiles(dir, author, date, comment)

	return err
}

func (t tasker) WithSubscribes(sm *subscription.SubscribeManager) Flow {

	return subscribeTasker{
		tasker: t,
		log:    t.log.Named("subscription"),
		pm:     sm,
	}

}

func (t tasker) ReadVersionFile(end types.V8Endpoint, dir string, filename string) (int64, error) {

	type versionReader struct {
		CurrentVersion int64 `xml:"VERSION"`
	}

	fileVesrion := filepath.Join(dir, filename)

	// Open our xmlFile
	xmlFile, err := os.Open(fileVesrion)
	// if we os.Open returns an error then handle it
	if err != nil {
		return 0, err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	var r = versionReader{}

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(byteValue, &r.CurrentVersion)

	if err != nil {
		return 0, err
	}

	return r.CurrentVersion, nil
}
