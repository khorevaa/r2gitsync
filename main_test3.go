package main

//
//import (
//	"fmt"
//	"github.com/go-git/go-billy/v5/osfs"
//	"github.com/go-git/go-git/v5"
//	"github.com/go-git/go-git/v5/plumbing/cache"
//	"github.com/go-git/go-git/v5/storage/filesystem"
//	"github.com/khorevaa/r2gitsync/manager/flow"
//	"github.com/stretchr/testify/require"
//	"github.com/stretchr/testify/suite"
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"testing"
//)
//
//type mainTestSuite struct {
//	suite.Suite
//
//	RepositoryPath string
//	WorkdirPath    string
//}
//
//func TestManagerTestSuite(t *testing.T) {
//	suite.Run(t, new(mainTestSuite))
//}
//
//func (s *mainTestSuite) r() *require.Assertions {
//	return s.Require()
//}
//
//func (t *mainTestSuite) AfterTest(suite, testName string) {
//
//}
//func (t *mainTestSuite) BeforeTest(suite, testName string) {
//
//}
//
//func (t *mainTestSuite) TearDownSuite() {
//	//os.RemoveAll(t.RepositoryPath)
//	//os.RemoveAll(t.WorkdirPath)
//}
//
//func (t *mainTestSuite) TearDownTest() {
//
//}
//func (t *mainTestSuite) SetupSuite() {
//
//	//t.prepare()
//}
//
//func (t *mainTestSuite) prepare() {
//
//	// Скопировать хранилище 1 файл
//	// Создать git репо
//	// Создать файл AUTHORS
//	// СОздать файл VERSION
//	//
//	//t.preRepository()
//	//t.preWorkdir()
//	//t.preVersion()
//
//}
//
//func (t *mainTestSuite) preVersion() {
//
//	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
//<VERSION>%d</VERSION>`, 0)
//
//	filename := filepath.Join(t.WorkdirPath, VERSION_FILE)
//	err := ioutil.WriteFile(filename, []byte(data), 0644)
//	t.r().NoError(err)
//
//}
//
//func (t *mainTestSuite) preRepository() {
//
//	t.RepositoryPath, _ = ioutil.TempDir("", "1c_repo")
//
//	srcRepositoryFile := filepath.Join(pwd, "..", "tests", "fixtures", "1cv8ddb.1CD")
//	dstRepoFile := filepath.Join(t.RepositoryPath, "1cv8ddb.1CD")
//
//	err := flow.CopyFile(srcRepositoryFile, dstRepoFile)
//	t.r().NoError(err)
//	fileBaseCreated, err := Exists(dstRepoFile)
//	t.r().True(fileBaseCreated, "Файл базы хранилища должен быть создан")
//
//}
//
//func (t *mainTestSuite) preWorkdir() {
//
//	t.WorkdirPath, _ = ioutil.TempDir("", "1c_git")
//
//	fs := osfs.New(t.WorkdirPath)
//	dot, _ := fs.Chroot(".git")
//	storage := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())
//
//	r, err := git.Init(storage, fs)
//	t.r().NoError(err)
//
//	_, err = fs.ReadDir(".git")
//	t.r().NoError(err)
//
//	cfg, err := r.Config()
//	t.r().NoError(err)
//	t.r().Equal(cfg.Core.Worktree, "")
//
//}
