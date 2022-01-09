package cmd

//
// import (
// 	"github.com/cucumber/godog"
// 	"github.com/khorevaa/r2gitsync/tests/bdd"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// 	"testing"
// )
//
// type cmdSetVersionTestSuite struct {
// 	suite.Suite
// 	BaseSuite
// }
//
// func TestCmdSetVersionTestSuite(t *testing.T) {
// 	suite.Run(t, new(cmdSetVersionTestSuite))
// }
//
// func (s *cmdSetVersionTestSuite) r() *require.Assertions {
// 	return s.Require()
// }
//
// func (s *cmdSetVersionTestSuite) TestFeatures() {
//
// 	err := bdd.Run([]string{
// 		"features/setVersion.feature",
// 	},
// 		s.InitializeCmdSyncTestSuite,
// 		s.InitializeCmdSyncScenario)
//
// 	s.r().NoError(err)
//
// }
//
// type cmdSetVersionContext struct {
// 	*Context
// 	app *AppTestSuite
// }
//
// func (s *cmdSetVersionTestSuite) InitializeCmdSyncScenario(ctx *godog.ScenarioContext) {
//
// 	sharedCtx := SharedContext(ctx)
//
// 	_ = &cmdSetVersionContext{
// 		Context: sharedCtx,
// 		app:     SharedAppSteps(sharedCtx, ctx),
// 	}
// 	// ctx.Step(`^Я инициализирую новое приложение$`, step.InitNewApp)
// 	// ctx.Step(`^я скопировал каталог "([^"]*)" в каталог из переменной "([^"]*)"$`, step.CopyDirToDirFromContext )
//
// }
