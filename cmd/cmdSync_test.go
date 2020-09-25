package cmd

import (
	"github.com/cucumber/godog"
	"github.com/khorevaa/r2gitsync/cmd/features/steps"
	"github.com/khorevaa/r2gitsync/internal/bdd"
	"github.com/khorevaa/r2gitsync/manager/flow"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/v8platform/designer"
	"github.com/v8platform/find"
	v8 "github.com/v8platform/v8"
	"testing"
)

type cmdSyncTestSuite struct {
	suite.Suite
	steps.BaseSuite

	RepositoryPath string
	WorkdirPath    string
	v8Version      string
}

func TestCmdSyncTestSuite(t *testing.T) {
	suite.Run(t, new(cmdSyncTestSuite))
}

func (s *cmdSyncTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *cmdSyncTestSuite) TestFeatures() {

	err := bdd.Run([]string{
		"features/sync.feature",
	},
		s.InitializeCmdSyncTestSuite,
		s.InitializeCmdSyncScenario)

	s.r().NoError(err)

}

type cmdSyncPluginsTestSuite struct {
	suite.Suite
	steps.BaseSuite

	RepositoryPath string
	WorkdirPath    string
	v8Version      string
}

func TestSyncPluginsTestSuite(t *testing.T) {
	suite.Run(t, new(cmdSyncPluginsTestSuite))
}

func (s *cmdSyncPluginsTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *cmdSyncPluginsTestSuite) TestFeatures() {

	err := bdd.Run([]string{
		"features/plugins.feature",
	},
		s.InitializeCmdSyncTestSuite,
		s.InitializeCmdSyncScenario)

	s.r().NoError(err)

}

type cmdSyncContext struct {
	*steps.Context
	app *AppTestSuite
}

func (s *cmdSyncContext) createTempIB(name string) error {
	dir := s.Context.String(name)
	return createTempIb(dir)
}

func (s *cmdSyncContext) findAndSavev8Path(name string) error {

	path, err := find.PlatformByVersion("8.3")

	s.Context.Set(name, path)

	return err
}

func (s *cmdSyncTestSuite) InitializeCmdSyncScenario(ctx *godog.ScenarioContext) {

	sharedCtx := steps.SharedContext(ctx)

	step := &cmdSyncContext{
		Context: sharedCtx,
		app:     SharedAppSteps(sharedCtx, ctx),
	}
	ctx.Step(`^Я создаю пустую базы в каталоге из переменной "([^"]*)"$`, step.createTempIB)
	ctx.Step(`^Я ищю платформу и сохраняю в переменной "([^"]*)"$`, step.findAndSavev8Path) //ctx.Step(`^Я инициализирую новое приложение$`, step.InitNewApp)
	//ctx.Step(`^я скопировал каталог "([^"]*)" в каталог из переменной "([^"]*)"$`, step.CopyDirToDirFromContext )

}

func (s *cmdSyncPluginsTestSuite) InitializeCmdSyncScenario(ctx *godog.ScenarioContext) {

	sharedCtx := steps.SharedContext(ctx)

	step := &cmdSyncContext{
		Context: sharedCtx,
		app:     SharedAppSteps(sharedCtx, ctx),
	}

	step.app.Application = NewApp("dev", true)

}

func createTempIb(dir string) error {

	crOpts := designer.CreateFileInfoBaseOptions{
		File: dir,
	}

	ib := v8.NewFileIB(dir)

	var opts []interface{}

	err := flow.Run(ib, crOpts, opts)

	return err

}
