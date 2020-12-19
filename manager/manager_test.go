package manager

import (
	"encoding/xml"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/khorevaa/r2gitsync/internal/bdd"
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/manager/flow"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	v8 "github.com/v8platform/api"
	"github.com/v8platform/designer"
	"github.com/v8platform/designer/tests"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var pwd, _ = os.Getwd()

type managerTestSuite struct {
	suite.Suite

	SycnRepo       *SyncRepository
	RepositoryPath string
	WorkdirPath    string
	ctx            map[string]string
	v8Version      string
}

func TestManagerTestSuite(t *testing.T) {
	suite.Run(t, new(managerTestSuite))
}

func (s *managerTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *managerTestSuite) AfterTest(suite, testName string) {

}
func (s *managerTestSuite) BeforeTest(suite, testName string) {

}

func (s *managerTestSuite) TearDownSuite() {
	os.RemoveAll(s.RepositoryPath)
	os.RemoveAll(s.WorkdirPath)
}

func (s *managerTestSuite) TearDownTest() {

}
func (s *managerTestSuite) SetupSuite() {

	s.prepare()
}

func (s *managerTestSuite) prepare() {

	// Скопировать хранилище 1 файл
	// Создать git репо
	// Создать файл AUTHORS
	// СОздать файл VERSION
	srcRepositoryFile := filepath.Join(pwd, "..", "tests", "fixtures", "storage")

	s.preRepository(srcRepositoryFile)
	s.preWorkdir()
	s.preVersion()

}

func (s *managerTestSuite) preVersion() {

	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<VERSION>%d</VERSION>`, 0)

	filename := filepath.Join(s.WorkdirPath, VERSION_FILE)
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	s.r().NoError(err)

}

func (s *managerTestSuite) preRepository(srcRepositoryFile string) {

	s.RepositoryPath, _ = ioutil.TempDir("", "1c_repo")

	dstRepoFile := filepath.Join(s.RepositoryPath, "1cv8ddb.1CD")

	err := flow.CopyDir(srcRepositoryFile, s.RepositoryPath)
	s.r().NoError(err)
	fileBaseCreated, err := Exists(dstRepoFile)
	s.r().True(fileBaseCreated, "Файл базы хранилища должен быть создан")

}

func (s *managerTestSuite) preWorkdir() {

	s.WorkdirPath, _ = ioutil.TempDir("", "1c_git")

	fs := osfs.New(s.WorkdirPath)
	dot, _ := fs.Chroot(".git")
	storage := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())

	r, err := git.Init(storage, fs)
	s.r().NoError(err)

	_, err = fs.ReadDir(".git")
	s.r().NoError(err)

	cfg, err := r.Config()
	s.r().NoError(err)
	s.r().Equal(cfg.Core.Worktree, "")

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

func (s *managerTestSuite) TestSimpleSync() {

	repo := SyncRepository{
		Name: "TestRepo",
		Repository: designer.Repository{
			Path: s.RepositoryPath,
		},
		Workdir: s.WorkdirPath,
	}

	logger := log.NewLogger()
	logger.SetDebug()

	err := repo.Sync(Options{
		MinVersion: 0,
		TempDir:    v8.NewTempDir("", "tests_"),
		Logger:     logger,
	})
	s.r().NoError(err)

}

func (s *managerTestSuite) TestSyncExtension() {

	srcRepositoryFile := filepath.Join(pwd, "..", "tests", "fixtures", "extension_storage")

	s.preRepository(srcRepositoryFile)

	repo := SyncRepository{
		Name: "TestExtension",
		Repository: designer.Repository{
			Path:      s.RepositoryPath,
			Extension: "MyExtension",
		},
		Workdir: s.WorkdirPath,
	}

	logger := log.NewLogger()
	logger.SetDebug()

	err := repo.Sync(Options{
		TempDir: v8.NewTempDir("", "tests_"),
		Logger:  logger,
	})
	s.r().NoError(err)

}

func (s *managerTestSuite) TestFeatures() {

	err := bdd.Run([]string{"features/sync.feature"},
		InitializeTestSuite,
		InitializeScenario)

	s.r().NoError(err)

}

func InitializeTestSuite(context *godog.TestSuiteContext) {

}

func (s *managerTestSuite) ToContext(name string, value string) {

	s.ctx[name] = value

}

func (s *managerTestSuite) FromContext(name string) string {

	return s.ctx[name]

}

func (s *managerTestSuite) createTempDirAndSaveToContext(name string) error {

	dir, _ := ioutil.TempDir("", "1c_temp")

	s.ToContext(name, dir)

	return nil

}

func (s *managerTestSuite) initGitRepoFromContext(name string) error {

	dir := s.FromContext(name)

	fs := osfs.New(dir)
	dot, _ := fs.Chroot(".git")
	storage := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())

	r, err := git.Init(storage, fs)
	if err != nil {
		return err
	}
	_, err = fs.ReadDir(".git")
	if err != nil {
		return err
	}
	_, err = r.Config()

	if err != nil {
		return err
	}
	s.WorkdirPath = dir

	return err

}

func (s *managerTestSuite) createTestAuthors() error {

	data := `Администратор=Админ <Администратор@localhost>`

	filename := filepath.Join(s.WorkdirPath, AUTHORS_FILE)
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	return err

}

func (s *managerTestSuite) createTestVersion(ver int) error {

	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<VERSION>%d</VERSION>`, ver)

	filename := filepath.Join(s.WorkdirPath, VERSION_FILE)
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	return err

}

func (s *managerTestSuite) copyTestRepoFromContext(name string) error {

	s.RepositoryPath = s.FromContext(name)

	srcRepository := filepath.Join(pwd, "..", "tests", "fixtures", "1cv8ddb.1CD")
	//dstRepoFile := filepath.Join(s.RepositoryPath, "1cv8ddb.1CD")

	err := flow.CopyDir(srcRepository, s.RepositoryPath)
	return err

}

func (s *managerTestSuite) CopyDirToDirFromContext(dir, name string) error {

	s.RepositoryPath = s.FromContext(name)

	srcRepositoryDir := filepath.Join(pwd, "..", dir)
	dstRepoDir := filepath.Join(s.RepositoryPath)

	err := flow.CopyDir(srcRepositoryDir, dstRepoDir)
	return err

}

func (s *managerTestSuite) setAuth(name string, pass string) error {

	if len(name) == 0 {
		return nil
	}

	s.SycnRepo.User = name
	s.SycnRepo.Password = pass
	return nil

}

func (s *managerTestSuite) setPlatformVersion(ver string) error {

	s.v8Version = ver
	return nil

}

func (s *managerTestSuite) versionFileContain(ver int) error {

	type versionReader struct {
		CurrentVersion int `xml:"VERSION"`
	}
	fileVersion := filepath.Join(s.WorkdirPath, VERSION_FILE)

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

	if r.CurrentVersion != ver {
		return fmt.Errorf("файл не содержит нужной версии: %s", string(byteValue))
	}

	return nil

}

func (s *managerTestSuite) doSync() error {

	repo := SyncRepository{
		Name: "TestRepo",
		Repository: designer.Repository{
			Path:     s.RepositoryPath,
			User:     s.SycnRepo.User,
			Password: s.SycnRepo.Password,
		},
		Workdir: s.WorkdirPath,
	}

	logger := log.NewLogger()
	logger.SetDebug()

	err := repo.Sync(
		Options{
			TempDir: v8.NewTempDir("", "tests_"),
			Logger:  logger,
		})

	return err

}

func (s *managerTestSuite) doSyncExt(extension string) error {

	repo := SyncRepository{
		Name: "TestRepo",
		Repository: designer.Repository{
			Path:      s.RepositoryPath,
			User:      s.SycnRepo.User,
			Password:  s.SycnRepo.Password,
			Extension: extension,
		},

		Workdir: s.WorkdirPath,
	}
	logger := log.NewLogger()
	logger.SetDebug()

	err := repo.Sync(
		Options{
			TempDir: v8.NewTempDir("", "tests_"),
			//Logger:           logger,
		},
	)

	return err

}

func InitializeScenario(ctx *godog.ScenarioContext) {
	feature := &managerTestSuite{
		ctx:      make(map[string]string),
		SycnRepo: &SyncRepository{},
	}

	ctx.Step(`^Я устанавливаю версию платформы "([^"]*)"$`, feature.setPlatformVersion)
	ctx.Step(`^Я выполняю выполняют синхронизацию$`, feature.doSync)
	ctx.Step(`^Я выполняю выполняют синхронизацию для расширения "([^"]*)"$`, feature.doSyncExt)
	ctx.Step(`^Я создаю временный каталог и сохраняю его в переменной "([^"]*)"$`, feature.createTempDirAndSaveToContext)
	ctx.Step(`^я скопировал каталог тестового хранилища конфигурации в каталог из переменной "([^"]*)"$`, feature.copyTestRepoFromContext)
	ctx.Step(`^Я инициализирую репозиторий в каталоге из переменной "([^"]*)"$`, feature.initGitRepoFromContext)
	//ctx.Step(`^Я включаю отладку лога с именем "([^"]*)"$`, StepDefinitioninition7)
	ctx.Step(`^Я устанавливаю авторизацию в хранилище пользователя "([^"]*)" с паролем "([^"]*)"$`, feature.setAuth)
	ctx.Step(`^Я создаю тестовой файл AUTHORS$`, feature.createTestAuthors)
	ctx.Step(`^Я записываю "([^"]*)" в файл VERSION$`, feature.createTestVersion)
	ctx.Step(`^Файл VERSION содержит (\d+)$`, feature.versionFileContain)
	ctx.Step(`^я скопировал каталог "([^"]*)" в каталог из переменной "([^"]*)"$`, feature.CopyDirToDirFromContext)
	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		//os.RemoveAll(feature.RepositoryPath)
		//os.RemoveAll(feature.WorkdirPath)
	})
}
