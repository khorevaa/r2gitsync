package manager

import (
	"fmt"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/manager/flow"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/v8platform/designer/repository"
	"github.com/v8platform/designer/tests"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var pwd, _ = os.Getwd()

type managerTestSuite struct {
	suite.Suite

	RepositoryPath string
	WorkdirPath    string
}

func TestManagerTestSuite(t *testing.T) {
	suite.Run(t, new(managerTestSuite))
}

func (s *managerTestSuite) r() *require.Assertions {
	return s.Require()
}

func (t *managerTestSuite) AfterTest(suite, testName string) {

}
func (t *managerTestSuite) BeforeTest(suite, testName string) {

}

func (t *managerTestSuite) TearDownSuite() {
	os.RemoveAll(t.RepositoryPath)
	os.RemoveAll(t.WorkdirPath)
}

func (t *managerTestSuite) TearDownTest() {

}
func (t *managerTestSuite) SetupSuite() {

	t.prepare()
}

func (t *managerTestSuite) prepare() {

	// Скопировать хранилище 1 файл
	// Создать git репо
	// Создать файл AUTHORS
	// СОздать файл VERSION

	t.preRepository()
	t.preWorkdir()
	t.preVersion()

}

func (t *managerTestSuite) preVersion() {

	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<VERSION>%d</VERSION>`, 0)

	filename := filepath.Join(t.WorkdirPath, VERSION_FILE)
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	t.r().NoError(err)

}

func (t *managerTestSuite) preRepository() {

	t.RepositoryPath, _ = ioutil.TempDir("", "1c_repo")

	srcRepositoryFile := filepath.Join(pwd, "..", "tests", "fixtures", "1cv8ddb.1CD")
	dstRepoFile := filepath.Join(t.RepositoryPath, "1cv8ddb.1CD")

	err := flow.CopyFile(srcRepositoryFile, dstRepoFile)
	t.r().NoError(err)
	fileBaseCreated, err := Exists(dstRepoFile)
	t.r().True(fileBaseCreated, "Файл базы хранилища должен быть создан")

}

func (t *managerTestSuite) preWorkdir() {

	t.WorkdirPath, _ = ioutil.TempDir("", "1c_git")

	fs := osfs.New(t.WorkdirPath)
	dot, _ := fs.Chroot(".git")
	storage := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())

	r, err := git.Init(storage, fs)
	t.r().NoError(err)

	_, err = fs.ReadDir(".git")
	t.r().NoError(err)

	cfg, err := r.Config()
	t.r().NoError(err)
	t.r().Equal(cfg.Core.Worktree, "")

}

func NewTempIB() tests.TempInfobase {

	path, _ := ioutil.TempDir("", "1c_DB_")

	ib := tests.TempInfobase{
		//InfoBase: InfoBase{},
		File: path,
	}

	return ib
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func (t *managerTestSuite) TestSimpleSync() {

	repo := SyncRepository{
		Name: "TestRepo",
		Repository: repository.Repository{
			Path: t.RepositoryPath,
		},
		Workdir: t.WorkdirPath,
	}

	logger := log.NewLogger()
	logger.SetDebug()

	err := repo.Sync(
		WithLogger(logger), // TODO FIX
	)
	t.r().NoError(err)

}

func (t *managerTestSuite) TestSyncExtension() {

	repo := SyncRepository{
		Name: "TestExtension",
		Repository: repository.Repository{
			Path: t.RepositoryPath,
		},
		Workdir:   t.WorkdirPath,
		Extension: "MyExtension",
	}

	logger := log.NewLogger()
	logger.SetDebug()

	err := repo.Sync(
		WithLogger(logger), // TODO FIX
	)
	t.r().NoError(err)

}
