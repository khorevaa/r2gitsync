package bdd

import "github.com/cucumber/godog"

type Option func(suite *Suite)

func Run(features []string,
	initSuite func(context *godog.TestSuiteContext),
	InitScenario func(context *godog.ScenarioContext),
	opts ...Option) error {

	suite := newSuite()
	suite.Paths = features
	suite.TestSuiteInitializer = append(suite.TestSuiteInitializer, initSuite)
	suite.ScenarioInitializer = append(suite.ScenarioInitializer, InitScenario)

	for _, opt := range opts {
		opt(suite)
	}

	return suite.Run()

}
