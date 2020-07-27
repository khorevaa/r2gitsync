package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hashicorp/go-multierror"
	"github.com/v8platform/designer"
	"github.com/v8platform/designer/repository"
	"github.com/v8platform/v8"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func syncInfobase(connString, user, password string) v8.Infobase {

	if len(connString) == 0 {
		return v8.NewTempIB()
	}
	// TODO Сделать получение базы для выполнения синхронизации
	return v8.NewTempIB()

}

type repositoryVersion struct {
	Version string
	Author  string
	Date    time.Time
	Comment string
	Number  int64
}

type author struct {
	Name  string
	Email string
}

type AuthorsList map[string]author

func (a author) Desc() string {

	return fmt.Sprintf("%s <%s> ", a.Name, a.Email)
}

type SyncRepository struct {
	repository.Repository
	name               string
	workDir            string
	currentVersion     int64 `xml:"VERSION"`
	maxVersion         int64
	repositoryVersions []repositoryVersion
	authors            *AuthorsList
	extention          string
}

func NewSyncRepository(path string) *SyncRepository {

	return &SyncRepository{
		Repository: repository.Repository{
			Path: path,
		},
	}
}

func (r *SyncRepository) Auth(user, passowrd string) {

	r.User = user
	r.Password = passowrd

}

func (r *SyncRepository) readCurrentVersion() error {

	fileVesrion := path.Join(r.workDir, VERSION_FILE)

	// Open our xmlFile
	xmlFile, err := os.Open(fileVesrion)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, r)

	return nil

}

func (r *SyncRepository) sync(opts *SyncOptions) error {

	plugins := opts.plugins

	plugins.BeforeStartSyncProcess(r.Repository, r.workDir)

	err := r.prepare(opts)

	if err != nil {
		return err
	}

	if len(r.repositoryVersions) == 0 {
		fmt.Printf("No versions to sync")
		return nil
	}

	nextVersion := r.repositoryVersions[0].Number
	maxVersion := r.maxVersion

	plugins.BeforeStartSyncVersions(&r.repositoryVersions, r.currentVersion, nextVersion, &maxVersion)

	for _, rVersion := range r.repositoryVersions {

		if r.maxVersion != 0 && rVersion.Number > r.maxVersion {
			break
		}

		cancelSync := false

		err := plugins.BeforeSyncVersion(rVersion.Number, rVersion.Author, rVersion.Comment, &cancelSync, opts)

		if err != nil {
			return err
		}

		if cancelSync {
			break
		}

		err = r.syncVersionFiles(rVersion, opts)

		if err != nil {
			return err
		}

		r.currentVersion = rVersion.Number
		err = plugins.AfterSyncVersion(rVersion.Number, rVersion.Author, rVersion.Comment, opts)

		if err != nil {
			return err
		}

	}

	return nil
}

func (r *SyncRepository) WriteVersionFile(version int64) error {

	data := fmt.Sprintf(`<?xml Version=""1.0"" encoding=""UTF-8""?>
	"<VERSION>"%s"</VERSION>"`, strconv.FormatInt(version, 10))

	filename := filepath.Join(r.workDir, VERSION_FILE)
	err := ioutil.WriteFile(filename, []byte(data), 0644)

	return err

}

func (r *SyncRepository) CommitVersionFile(author string, when time.Time, comment string) error {

	g, err := git.PlainOpen(r.workDir)

	if err != nil {
		return err
	}

	filename := filepath.Join(r.workDir, VERSION_FILE)

	w, err := g.Worktree()

	if err != nil {
		return err
	}

	w.Add(filename)

	c, err := w.Commit(comment, &git.CommitOptions{
		Author: &object.Signature{
			Name:  author,
			Email: author,
			When:  when,
		},
	})

	if err != nil {
		return err
	}

	_, err = g.CommitObject(c)

	return err
}

func (r *SyncRepository) commitFiles(author string, when time.Time, comment string) error {

	g, err := git.PlainOpen(r.workDir)

	if err != nil {
		return err
	}

	w, err := g.Worktree()

	if err != nil {
		return err
	}

	w.AddGlob(r.workDir)

	c, err := w.Commit(comment, &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  author,
			Email: author,
			When:  when,
		},
	})

	if err != nil {
		return err
	}

	_, err = g.CommitObject(c)

	return err

}

func (r *SyncRepository) prepare(opts *SyncOptions) error {

	r.currentVersion = 0

	err := r.readCurrentVersion()

	if err != nil {
		return err
	}

	err = r.GetRepositoryAuthors(opts)

	if err != nil {
		return err
	}

	r.maxVersion = 0

	err = r.GetRepositoryHistory(opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *SyncRepository) Sync(workDir string, opts ...SyncOption) error {

	options := &SyncOptions{}

	r.workDir = workDir

	for _, opt := range opts {
		opt(options)
	}

	return r.sync(options)

}

func (r *SyncRepository) syncVersionFiles(rVersion repositoryVersion, opts *SyncOptions) (err error) {

	plugins := opts.plugins

	tempDir, err := ioutil.TempDir(opts.tempDir, fmt.Sprintf("v%d", rVersion.Number))

	if err != nil {
		return err
	}

	err = plugins.BeforeStartSyncVersionHandler(r.workDir, tempDir, r.Repository, rVersion.Number, r.extention)

	if err != nil {
		return err
	}

	defer func() {

		plugins.FinishSyncVersionHandler(r.workDir, tempDir, r.Repository, rVersion.Number, r.extention, err)

		_ = os.RemoveAll(tempDir)

	}()

	err = plugins.BeforeUpdateCfgHandler(r.workDir, opts.infobase, r.Repository, rVersion.Number, r.extention)

	if err != nil {
		return err
	}

	err = r.RepositoryUpdateCfg(rVersion.Number, opts)

	if err != nil {
		return err
	}

	err = plugins.BeforeDumpConfigToFiles(r.workDir, tempDir, opts.infobase, r.Repository, rVersion.Number, r.extention)

	if err != nil {
		return err
	}

	err = r.DumpConfigToFiles(tempDir, opts)

	if err != nil {
		return err
	}

	err = plugins.BeforeClearWorkDir(r.workDir, rVersion.Number)

	if err != nil {
		return err
	}

	err = r.ClearWorkDir(opts)

	if err != nil {
		return err
	}

	err = plugins.BeforeMoveToWorkDir(r.workDir, tempDir, rVersion.Number)

	if err != nil {
		return err
	}

	err = r.MoveToWorkDir(tempDir, opts)

	if err != nil {
		return err
	}

	err = r.WriteVersionFile(rVersion.Number)

	if err != nil {
		return err
	}

	err = r.commitFiles(rVersion.Author, rVersion.Date, rVersion.Comment)

	if err != nil {

		errV := r.WriteVersionFile(r.currentVersion)
		if errV != nil {
			return multierror.Append(err, errV)
		}
		return err
	}

	r.currentVersion = rVersion.Number

	return

}

func (r *SyncRepository) RepositoryUpdateCfg(version int64, options *SyncOptions) (err error) {

	standartHandler := true

	err = options.plugins.WithUpdateCfgHandler(options.infobase, r.Repository, version, r.extention, &standartHandler)

	if err != nil {
		return
	}

	if standartHandler {
		err = r.repositoryUpdateCfgHandler(version, options)
	}

	err = options.plugins.AfterUpdateCfgHandler(options.infobase, r.Repository, version, r.extention)

	return

}

func (r *SyncRepository) repositoryUpdateCfgHandler(version int64, opts *SyncOptions) error {

	RepositoryUpdateCfgOptions := repository.RepositoryUpdateCfgOptions{
		Version:   version,
		Force:     true,
		Extension: r.extention,
	}.
		WithRepository(r.Repository)

	err := Run(opts.infobase, RepositoryUpdateCfgOptions, opts)

	if err != nil {
		return err
	}
	return nil
}

func (r *SyncRepository) DumpConfigToFiles(dumpDir string, opts *SyncOptions) (err error) {

	standartHandler := true

	err = opts.plugins.WithDumpCfgToFilesHandler(opts.infobase, r.Repository, r.extention, &dumpDir, &standartHandler)

	if err != nil {
		return
	}

	if standartHandler {
		err = r.dumpConfigToFilesHandler(dumpDir, opts)
	}

	err = opts.plugins.AfterDumpCfgToFilesHandler(opts.infobase, r.Repository, r.extention, &dumpDir)

	return

}

func (r *SyncRepository) dumpConfigToFilesHandler(dumpDir string, opts *SyncOptions) error {

	DumpConfigToFilesOptions := designer.DumpConfigToFilesOptions{
		Dir:       dumpDir,
		Force:     true,
		Extension: r.extention,
	}

	err := Run(opts.infobase, DumpConfigToFilesOptions, opts)

	if err != nil {
		return err
	}
	return nil
}

func (r *SyncRepository) ClearWorkDir(opts *SyncOptions) (err error) {

	standartHandler := true

	err = opts.plugins.WithClearWorkdirHandler(r.workDir, &standartHandler)

	if err != nil {
		return
	}

	if standartHandler {
		err = r.clearWorkDirHandler(opts)
	}

	err = opts.plugins.AfterClearWorkdirHandler(r.workDir)

	return

}

func (r *SyncRepository) clearWorkDirHandler(opts *SyncOptions) error {

	err := os.RemoveAll(r.workDir) // TODO Сделать копирование файлов или избранную очистку

	if err != nil {
		return err
	}
	return nil
}

func (r *SyncRepository) MoveToWorkDir(dumpDir string, opts *SyncOptions) (err error) {

	standartHandler := true

	err = opts.plugins.WithMoveToWorkDirHandler(r.workDir, dumpDir, &standartHandler)

	if err != nil {
		return
	}

	if standartHandler {
		err = r.moveToWorkDirHandler(dumpDir, opts)
	}

	err = opts.plugins.AfterMoveToWorkDirHandler(r.workDir, dumpDir)

	return

}

func (r *SyncRepository) moveToWorkDirHandler(dumpDir string, opts *SyncOptions) error {

	err := os.RemoveAll(r.workDir) // TODO Сделать копирование файлов или избранную очистку

	if err != nil {
		return err
	}
	return nil
}

func (r *SyncRepository) GetRepositoryHistory(opts *SyncOptions) (err error) {

	standartHandler := true

	err = opts.plugins.WithGetRepositoryHistoryHandler(r.workDir, r.Repository, &r.repositoryVersions, opts, &standartHandler)

	if err != nil {
		return
	}

	if standartHandler {
		err = r.getRepositoryHistoryHandler(opts)
	}

	err = opts.plugins.AfterGetRepositoryHistoryHandler(r.workDir, r.Repository, &r.repositoryVersions)

	return

}

func (r *SyncRepository) getRepositoryHistoryHandler(opts *SyncOptions) error {

	reportFile, err := ioutil.TempFile(opts.tempDir, "v8_rep_history")
	if err != nil {
		return err
	}
	reportFile.Close()
	report := reportFile.Name()

	defer os.Remove(report)

	RepositoryReportOptions := repository.RepositoryReportOptions{
		Repository: r.Repository,
		File:       report,
		Extension:  r.extention,
		NBegin:     r.currentVersion,
	}.GroupByComment()

	err = Run(opts.infobase, RepositoryReportOptions, opts)

	if err != nil {
		return err
	}

	r.repositoryVersions, err = parseRepositoryReport(report)

	if err != nil {
		return err
	}

	sort.Slice(r.repositoryVersions, func(i, j int) bool {
		return r.repositoryVersions[i].Number < r.repositoryVersions[j].Number
	})

	if len(r.repositoryVersions) > 0 {
		r.maxVersion = r.repositoryVersions[len(r.repositoryVersions)-1].Number
	}

	return nil
}

func (r *SyncRepository) GetRepositoryAuthors(opts *SyncOptions) (err error) {

	standartHandler := true
	authors := new(AuthorsList)

	err = opts.plugins.WithGetRepositoryAuthorsHandler(r.workDir, r.Repository, authors, opts, &standartHandler)

	if err != nil {
		r.authors = authors
		return
	}

	if standartHandler {
		err = r.getRepositoryAuthorsHandler(authors, opts)
	}

	err = opts.plugins.AfterGetGetRepositoryAuthorsHandler(r.workDir, r.Repository, authors)

	return

}

func (r *SyncRepository) getRepositoryAuthorsHandler(authors *AuthorsList, opts *SyncOptions) error {

	file := path.Join(r.workDir, AUTHORS_FILE)
	_, err := os.Lstat(file)

	r.authors = new(AuthorsList)
	authorsTable := *r.authors
	if err != nil {
		authors = &authorsTable
		return nil
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

		authorsTable[data[0]] = NewAuthor(decodeAuthor([]byte(data[1])))

	}

	r.authors = &authorsTable

	return nil

}

func NewAuthor(name, email string) author {

	return author{
		Name:  name,
		Email: email,
	}

}

// Decode decodes a byte slice into a signature
func decodeAuthor(b []byte) (string, string) {
	open := bytes.LastIndexByte(b, '<')
	closeSym := bytes.LastIndexByte(b, '>')
	if open == -1 || closeSym == -1 {
		return "", ""
	}

	if closeSym < open {
		return "", ""
	}

	Name := string(bytes.Trim(b[:open], " "))
	Email := string(b[open+1 : closeSym])

	return Name, Email

}
