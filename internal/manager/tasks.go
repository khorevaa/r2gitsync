package manager

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/khorevaa/r2gitsync/internal/manager/types"
	"github.com/khorevaa/r2gitsync/pkg/plugin/subscription"
	v8 "github.com/v8platform/api"
	"github.com/v8platform/designer"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UpdateCfg struct {
	V8end   V8Endpoint
	WorkDir string
	Number  int
}

func (t UpdateCfg) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}
	return pm.UpdateCfg.Before(t.V8end, t.WorkDir, t.Number)
}

func (t UpdateCfg) After(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}
	return pm.UpdateCfg.After(t.V8end, t.WorkDir, t.Number)
}

func (t UpdateCfg) On(pm *subscription.SubscribeManager, useStdHandler *bool) error {
	if pm == nil {
		return nil
	}
	return pm.UpdateCfg.On(t.V8end, t.WorkDir, t.Number, useStdHandler)
}

func (t UpdateCfg) Action(useStdHandler bool) error {

	if !useStdHandler {
		return nil
	}

	updateCfg := t.V8end.Repository().UpdateCfg(int64(t.Number), true)
	return v8.Run(*t.V8end.Infobase(), updateCfg, t.V8end.Options()...)

}

type DumpConfigToFiles struct {
	V8end   V8Endpoint
	WorkDir string
	TempDir       string
	Number        int
	Increment     bool
	IsIncremented *bool
}

func (t DumpConfigToFiles) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}
	return pm.DumpConfigToFiles.Before(t.V8end, t.WorkDir, t.TempDir, t.Number, &t.Increment)
}

func (t DumpConfigToFiles) After(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}
	return pm.DumpConfigToFiles.After(t.V8end, t.WorkDir, t.TempDir, t.Number, *t.IsIncremented)
}

func (t DumpConfigToFiles) On(pm *subscription.SubscribeManager, useStdHandler *bool) error {
	if pm == nil {
		return nil
	}
	return pm.DumpConfigToFiles.On(t.V8end, t.WorkDir, t.TempDir, t.Number, &t.Increment, useStdHandler)
}

func (t DumpConfigToFiles) Action(useStdHandler bool) error {

	if !useStdHandler {
		return nil
	}

	var configDumpInfoFile = ""

	if t.IsIncremented == nil {
		t.IsIncremented = new(bool)
	}

	if t.Increment {
		configDumpInfoFile = filepath.Join(t.WorkDir, ConfigDumpInfoFileName)

		if ok, _ := Exists(configDumpInfoFile); !ok {
			t.Increment = false
			configDumpInfoFile = ""
		}
	}

	if t.Increment {

		changesFile := v8.NewTempFile("", ".txt")

		getChangesForConfigDumpOptions := designer.GetChangesForConfigDumpOptions{
			Dir:                      t.TempDir,
			Force:                    true,
			GetChanges:               changesFile,
			Extension:                t.V8end.Extention(),
			ConfigDumpInfoForChanges: configDumpInfoFile,
		}

		err := v8.Run(*t.V8end.Infobase(), getChangesForConfigDumpOptions, t.V8end.Options()...)

		if err != nil {
			return err
		}

		t.Increment, err = t.checkChangesFile(changesFile)
		if err != nil {
			return err
		}

		if !t.Increment {
			configDumpInfoFile = ""
		}
	}

	DumpConfigToFilesOptions := designer.DumpConfigToFilesOptions{
		Dir:                      t.TempDir,
		Force:                    true,
		Update:                   t.Increment,
		Extension:                t.V8end.Extention(),
		ConfigDumpInfoForChanges: configDumpInfoFile,
	}

	err := v8.Run(*t.V8end.Infobase(), DumpConfigToFilesOptions, t.V8end.Options()...)
	if err == nil && t.Increment {
		*t.IsIncremented = true
	}

	return err

}

func (t DumpConfigToFiles) checkChangesFile(filename string) (bool, error) {
	// Open our xmlFile
	file, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		return false, err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer file.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(file)

	fullDump := bytes.Contains(byteValue, []byte("FullDump"))

	return !fullDump, nil
}

type StartSyncVersion struct {
	V8end   V8Endpoint
	WorkDir string
	TempDir string
	Number  int
}

func (t StartSyncVersion) Before(pm *subscription.SubscribeManager) error {
	return nil
}

func (t StartSyncVersion) After(pm *subscription.SubscribeManager) error {
	return nil
}

func (t StartSyncVersion) On(pm *subscription.SubscribeManager, _ *bool) error {
	if pm == nil {
		return nil
	}

	pm.SyncVersion.Start(t.V8end, t.WorkDir, t.TempDir, t.Number)

	return nil
}

func (t StartSyncVersion) Action(useStdHandler bool) error {
	return nil
}

type FinishSyncVersion struct {
	V8end   V8Endpoint
	WorkDir string
	TempDir string
	Number  int
	Err     *error
}

func (t FinishSyncVersion) Before(_ *subscription.SubscribeManager) error {
	return nil
}

func (t FinishSyncVersion) After(_ *subscription.SubscribeManager) error {
	return nil
}

func (t FinishSyncVersion) On(pm *subscription.SubscribeManager, _ *bool) error {
	if pm == nil {
		return nil
	}

	pm.SyncVersion.Finish(t.V8end, t.WorkDir, t.TempDir, t.Number, t.Err)

	return nil
}

func (t FinishSyncVersion) Action(_ bool) error {
	return nil
}

type StartSyncProcess struct {
	V8end   V8Endpoint
	WorkDir string
}

func (t StartSyncProcess) Before(pm *subscription.SubscribeManager) error {
	return nil
}

func (t StartSyncProcess) After(pm *subscription.SubscribeManager) error {
	return nil
}

func (t StartSyncProcess) On(pm *subscription.SubscribeManager, _ *bool) error {
	if pm == nil {
		return nil
	}

	pm.SyncProcess.Start(t.V8end, t.WorkDir)

	return nil
}

func (t StartSyncProcess) Action(useStdHandler bool) error {
	return nil
}

type FinishSyncProcess struct {
	V8end   V8Endpoint
	WorkDir string
	Err     *error
}

func (t FinishSyncProcess) Before(_ *subscription.SubscribeManager) error {
	return nil
}

func (t FinishSyncProcess) After(_ *subscription.SubscribeManager) error {
	return nil
}

func (t FinishSyncProcess) On(pm *subscription.SubscribeManager, _ *bool) error {
	if pm == nil {
		return nil
	}

	pm.SyncProcess.Finish(t.V8end, t.WorkDir, t.Err)

	return nil
}

func (t FinishSyncProcess) Action(_ bool) error {
	return nil
}

type ConfigureRepositoryVersions struct {
	V8end               V8Endpoint
	Versions            *types.RepositoryVersionsList
	NBegin, NNext, NMax *int
}

func (t ConfigureRepositoryVersions) Before(pm *subscription.SubscribeManager) error {
	return nil
}

func (t ConfigureRepositoryVersions) After(pm *subscription.SubscribeManager) error {
	return nil
}

func (t ConfigureRepositoryVersions) On(pm *subscription.SubscribeManager, std *bool) error {
	if pm == nil {
		return nil
	}

	return pm.ConfigureRepositoryVersions.On(t.V8end, t.Versions, t.NBegin, t.NNext, t.NMax)
}

func (t ConfigureRepositoryVersions) Action(useStdHandler bool) error {

	if !useStdHandler {
		return nil
	}

	// TODO Добавить логику работы органичения
	// TODO Добавить логику получения максимальной версии

	ver := *t.Versions

	if len(ver) > 0 {
		maxVersion := ver[len(ver)-1].Number()
		*t.NMax = maxVersion
	}

	return nil
}

type GetRepositoryVersions struct {
	V8end        V8Endpoint
	Workdir      string
	NBegin, NEnd int
	Versions     *types.RepositoryVersionsList
}

func (t GetRepositoryVersions) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}

	err := pm.GetRepositoryHistory.Before(t.V8end, t.Workdir, t.NBegin)

	return err
}

func (t GetRepositoryVersions) After(pm *subscription.SubscribeManager) error {

	if pm == nil {
		return nil
	}

	err := pm.GetRepositoryHistory.After(t.V8end, t.Workdir, t.NBegin, t.Versions)

	return err
}

func (t GetRepositoryVersions) On(pm *subscription.SubscribeManager, std *bool) error {
	if pm == nil {
		return nil
	}

	v, err := pm.GetRepositoryHistory.On(t.V8end, t.Workdir, t.NBegin, std)

	if err == nil {
		*t.Versions = v
	}

	return err
}

func (t GetRepositoryVersions) Action(useStdHandler bool) (err error) {

	if !useStdHandler {
		return nil
	}
	//
	//t.log.Infow("Get repository versions",
	//	zap.String("path", t.Workdir),
	//	zap.Int("nBegin", t.NBegin))

	report := v8.NewTempFile("", "v8_rep_history")
	defer os.Remove(report)

	getReport := t.V8end.Repository().Report(report, int64(t.NBegin)).GroupByComment()

	err = v8.Run(t.V8end.Infobase(), getReport, t.V8end.Options()...)

	if err != nil {
		return
	}

	//t.log.Debugw("Parse repository report", zap.String("file", report))

	versions, err := t.parseRepositoryReport(report)

	if err != nil {
		return
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Number() < versions[j].Number()
	})

	*t.Versions = versions

	return nil
}

func (t GetRepositoryVersions) parseRepositoryReport(file string) (versions []types.RepositoryVersion, err error) {

	err, bytes := readFile(file, nil)
	if err != nil {
		return
	}

	report := string(*bytes)
	// Двойные кавычки в комментарии мешают, по этому мы заменяем из на одинарные
	report = strings.Replace(report, "\"\"", "'", -1)

	tmpArray := [][]string{}
	reg := regexp.MustCompile(`[{]"#","([^"]+)["][\}]`)
	matches := reg.FindAllStringSubmatch(report, -1)
	for _, s := range matches {
		if s[1] == "Версия:" {
			tmpArray = append(tmpArray, []string{})
		}

		if len(tmpArray) > 0 {
			tmpArray[len(tmpArray)-1] = append(tmpArray[len(tmpArray)-1], s[1])
		}
	}

	for _, array := range tmpArray {
		versionInfo := repositoryVersion{}
		for id, s := range array {
			switch s {
			case "Версия:":
				if ver, err := strconv.Atoi(array[id+1]); err == nil {
					versionInfo.number = ver
				}
			case "Версия конфигурации:":
				versionInfo.version = array[id+1]
			case "Пользователь:":
				versionInfo.author = array[id+1]
			case "Комментарий:":
				// Комментария может не быть, по этому вот такой костыльчик
				if array[id+1] != "Изменены:" {
					versionInfo.comment = strings.Replace(array[id+1], "\n", " ", -1)
					versionInfo.comment = strings.Replace(array[id+1], "\r", "", -1)
				}
			case "Дата создания:":
				if t, err := time.Parse("02.01.2006", array[id+1]); err == nil {
					versionInfo.date = t
				}
			case "Время создания:":
				if !versionInfo.date.IsZero() {
					str := versionInfo.date.Format("02.01.2006") + " " + array[id+1]
					if t, err := time.Parse("02.01.2006 15:04:05", str); err == nil {
						versionInfo.date = t
					}
				}
			}
		}
		versions = append(versions, versionInfo)
	}

	return
}

type GetRepositoryAuthors struct {
	V8end    V8Endpoint
	Workdir  string
	Filename string
	Authors  *types.RepositoryAuthorsList
}

func (t GetRepositoryAuthors) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}

	err := pm.GetRepositoryAuthors.Before(t.V8end, t.Workdir, t.Filename)

	return err
}

func (t GetRepositoryAuthors) After(pm *subscription.SubscribeManager) error {

	if pm == nil {
		return nil
	}

	err := pm.GetRepositoryAuthors.After(t.V8end, t.Workdir, t.Authors)

	return err
}

func (t GetRepositoryAuthors) On(pm *subscription.SubscribeManager, std *bool) error {
	if pm == nil {
		return nil
	}

	v, err := pm.GetRepositoryAuthors.On(t.V8end, t.Workdir, t.Filename, std)

	if err == nil {
		*t.Authors = v
	}

	return err
}

func (t GetRepositoryAuthors) Action(useStdHandler bool) (err error) {

	if !useStdHandler {
		return nil
	}
	authors := make(map[string]types.RepositoryAuthor)

	file := filepath.Join(t.Workdir, t.Filename)
	if ok, _ := Exists(file); !ok {

		return nil

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

	*t.Authors = authors

	return nil
}

type ClearWorkdir struct {
	V8end   V8Endpoint
	Workdir string
	TempDir   string
	SkipFiles []string
}

func (t ClearWorkdir) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}

	return pm.ClearWorkdir.Before(t.V8end, t.Workdir, t.TempDir)
}

func (t ClearWorkdir) After(pm *subscription.SubscribeManager) error {

	if pm == nil {
		return nil
	}

	return pm.ClearWorkdir.After(t.V8end, t.Workdir, t.TempDir)
}

func (t ClearWorkdir) On(pm *subscription.SubscribeManager, std *bool) error {
	if pm == nil {
		return nil
	}

	return pm.ClearWorkdir.On(t.V8end, t.Workdir, t.TempDir, std)
}

func (t ClearWorkdir) Action(useStdHandler bool) (err error) {

	if !useStdHandler {
		return nil
	}
	err = clearDir(t.Workdir, t.SkipFiles...)

	return
}

type MoveToWorkdir struct {
	V8end   V8Endpoint
	Workdir string
	TempDir string
}

func (t MoveToWorkdir) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}

	return pm.MoveToWorkdir.Before(t.V8end, t.Workdir, t.TempDir)
}

func (t MoveToWorkdir) After(pm *subscription.SubscribeManager) error {

	if pm == nil {
		return nil
	}

	return pm.MoveToWorkdir.After(t.V8end, t.Workdir, t.TempDir)
}

func (t MoveToWorkdir) On(pm *subscription.SubscribeManager, std *bool) error {
	if pm == nil {
		return nil
	}

	return pm.MoveToWorkdir.On(t.V8end, t.Workdir, t.TempDir, std)
}

func (t MoveToWorkdir) Action(useStdHandler bool) (err error) {

	if !useStdHandler {
		return nil
	}
	err = CopyDir(t.TempDir, t.Workdir)

	return
}

type WriteVersionFile struct {
	V8end   V8Endpoint
	Workdir string
	Filename string
	Version  int
}

func (t WriteVersionFile) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}

	return pm.WriteVersionFile.Before(t.V8end, t.Workdir, t.Version, t.Filename)
}

func (t WriteVersionFile) After(pm *subscription.SubscribeManager) error {

	if pm == nil {
		return nil
	}

	return pm.WriteVersionFile.After(t.V8end, t.Workdir, t.Version, t.Filename)
}

func (t WriteVersionFile) On(pm *subscription.SubscribeManager, std *bool) error {
	if pm == nil {
		return nil
	}

	return pm.WriteVersionFile.On(t.V8end, t.Workdir, t.Version, t.Filename, std)
}

func (t WriteVersionFile) Action(useStdHandler bool) (err error) {

	if !useStdHandler {
		return nil
	}

	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<VERSION>%d</VERSION>`, t.Version)

	filename := filepath.Join(t.Workdir, t.Filename)
	err = ioutil.WriteFile(filename, []byte(data), 0644)

	return
}

type CommitFiles struct {
	V8end   V8Endpoint
	Workdir string
	Author  types.RepositoryAuthor
	Date    time.Time
	Comment string
}

func (t CommitFiles) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}

	return pm.CommitFiles.Before(t.V8end, t.Workdir, t.Author, t.Date, t.Comment)
}

func (t CommitFiles) After(pm *subscription.SubscribeManager) error {

	if pm == nil {
		return nil
	}

	return pm.CommitFiles.After(t.V8end, t.Workdir, t.Author, t.Date, t.Comment)
}

func (t CommitFiles) On(pm *subscription.SubscribeManager, std *bool) error {
	if pm == nil {
		return nil
	}

	return pm.CommitFiles.On(t.V8end, t.Workdir, t.Author, &t.Date, &t.Comment, std)
}

func (t CommitFiles) Action(useStdHandler bool) (err error) {

	if !useStdHandler {
		return nil
	}

	return commitFiles(t.Workdir, t.Author, t.Date, t.Comment)
}

type ReadVersionFile struct {
	V8end   V8Endpoint
	Workdir string
	Filename string
	Version  *int
}

func (t ReadVersionFile) Before(pm *subscription.SubscribeManager) error {
	if pm == nil {
		return nil
	}

	return pm.ReadVersionFile.Before(t.V8end, t.Workdir, t.Filename)
}

func (t ReadVersionFile) After(pm *subscription.SubscribeManager) error {

	if pm == nil {
		return nil
	}

	return pm.ReadVersionFile.After(t.V8end, t.Workdir, t.Filename, t.Version)
}

func (t ReadVersionFile) On(pm *subscription.SubscribeManager, std *bool) error {
	if pm == nil {
		return nil
	}

	v, err := pm.ReadVersionFile.On(t.V8end, t.Workdir, t.Filename, std)

	if err == nil && t.Version != nil {
		*t.Version = v
	}

	return err
}

func (t ReadVersionFile) Action(useStdHandler bool) (err error) {

	if !useStdHandler {
		return nil
	}

	if t.Version == nil {
		t.Version = new(int)
	}

	type versionReader struct {
		CurrentVersion int `xml:"VERSION"`
	}

	fileVersion := filepath.Join(t.Workdir, t.Filename)

	// Open our xmlFile
	xmlFile, err := os.Open(fileVersion)
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	var r = versionReader{}

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(byteValue, &r.CurrentVersion)

	if err != nil {
		return err
	}

	*t.Version = r.CurrentVersion

	return nil
}
