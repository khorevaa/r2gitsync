package steps

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/manager"
	"github.com/khorevaa/r2gitsync/manager/flow"
	"io/ioutil"
	"os"
	"path/filepath"
)

var pwd, _ = os.Getwd()

type Context struct {
	context.Context
	Temp string

	Args []string
}

type BaseSuite struct {
}

func (s *BaseSuite) InitializeCmdSyncTestSuite(ctx *godog.TestSuiteContext) {

}

func (s *BaseSuite) InitializeCmdSyncScenario(ctx *godog.ScenarioContext) {
}

func (s *Context) createTempDirAndSaveToContext(name string) error {

	dir, _ := ioutil.TempDir(s.Temp, "1c_temp")

	s.Set(name, dir)

	return nil

}

func (s *Context) initGitRepoFromContext(name string) error {

	dir := s.String(name)

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

	return err

}

func (s *Context) CopyDirToDirFromContext(dir, name string) error {

	dst := s.String(name)
	src := filepath.Join(pwd, dir)
	err := flow.CopyDir(src, dst)
	return err

}

func (s *Context) createTestFile(file, name string, data *messages.PickleStepArgument_PickleDocString) error {

	dir := s.String(name)

	//data := `Администратор=Админ <Администратор@localhost>`

	filename := filepath.Join(dir, file)
	err := ioutil.WriteFile(filename, []byte(data.String()), 0644)
	return err

}

func (s *Context) createTestVersion(ver int, name string) error {

	data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<VERSION>%d</VERSION>`, ver)
	dir := s.String(name)

	filename := filepath.Join(dir, manager.VERSION_FILE)
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	return err

}

func (s *Context) fileContain(file, name string, data *messages.PickleStepArgument_PickleDocString) error {

	dir := s.String(name)

	filename := filepath.Join(dir, file)

	// Open our xmlFile
	xmlFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	contains := bytes.Contains(byteValue, []byte(data.String()))

	if !contains {

		return errors.New(fmt.Sprintf("файл не содержить нужной строки\n Содержание:\n %s", string(byteValue)))
	}

	return nil

}

func (s *Context) fileContainText(file, name string, data string) error {

	dir := s.String(name)

	filename := filepath.Join(dir, file)

	// Open our xmlFile
	xmlFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	contains := bytes.Contains(byteValue, []byte(data))

	if !contains {

		return errors.New(fmt.Sprintf("файл не содержить нужной строки\n Содержание:\n %s", string(byteValue)))
	}

	return nil

}

func (s *Context) setEnvValueFromContext(env, name string) error {

	value := s.String(name)
	err := os.Setenv(env, value)

	return err
}

func (s *Context) clearEnvs(envs *messages.PickleStepArgument_PickleTable) error {
	rows := envs.GetRows()

	for _, row := range rows {
		envName := row.Cells[0].Value

		err := os.Unsetenv(envName)
		if err != nil {
			return err
		}

	}

	return nil
}

func SharedContext(ctx *godog.ScenarioContext) *Context {

	step := &Context{
		Context: context.NewContext(),
	}
	ctx.BeforeScenario(func(sc *godog.Scenario) {
		step.Temp, _ = ioutil.TempDir("", "dbb_step")
	})
	ctx.Step(`^Я скопировал каталог "([^"]*)" в каталог из переменной "([^"]*)"$`, step.CopyDirToDirFromContext)
	ctx.Step(`^Я создаю временный каталог и сохраняю его в переменной "([^"]*)"$`, step.createTempDirAndSaveToContext)
	ctx.Step(`^Я инициализирую репозиторий в каталоге из переменной "([^"]*)"$`, step.initGitRepoFromContext)
	ctx.Step(`^Я создаю тестовой файл "([^"]*)" в каталоге из переменной "([^"]*)" с текстом:$`, step.createTestFile)
	//ctx.Step(`^Я записываю "([^"]*)" в файл "([^"]*) в каталоге из переменной "([^"]*)"$`, step.createTestVersion)
	ctx.Step(`^Файл "([^"]*)" в каталоге из переменной "([^"]*)" содержит "([^"]*)"$`, step.fileContainText)
	ctx.Step(`^Файл "([^"]*)" в каталоге из переменной "([^"]*)" содержит$`, step.fileContain)
	ctx.Step(`^Я устанавливаю переменную окружения "([^"]*)" из переменной "([^"]*)"$`, step.setEnvValueFromContext)
	ctx.Step(`^Я очищаю значение переменных окружения$`, step.clearEnvs)
	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		_ = os.RemoveAll(step.Temp)
	})
	return step
}
