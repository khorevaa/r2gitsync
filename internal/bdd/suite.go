package bdd

import (
	"errors"
	"github.com/cucumber/godog"
	"github.com/khorevaa/r2gitsync/context"
)

type Suite struct {
	Paths                []string
	Format               string
	Strict               bool
	Concurrency          int
	TestSuiteInitializer []func(context *godog.TestSuiteContext)
	ScenarioInitializer  []func(context *godog.ScenarioContext)

	context.Context
}

func (s *Suite) InitializeTestSuite(context *godog.TestSuiteContext) {

	for _, fn := range s.TestSuiteInitializer {
		fn(context)
	}

}

func (s *Suite) InitializeScenario(context *godog.ScenarioContext) {

	for _, fn := range s.ScenarioInitializer {
		fn(context)
	}

}

func (s *Suite) Run() error {

	opts := godog.Options{
		Format: s.Format,
		Paths:  s.Paths,
		//ShowStepDefinitions: true,
		//StopOnFailure: true,
		Strict:      s.Strict,
		Concurrency: s.Concurrency,
	}

	// godog v0.10.0 (latest)
	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: s.InitializeTestSuite,
		ScenarioInitializer:  s.InitializeScenario,
		Options:              &opts,
	}.Run()

	if status > 0 {

		return errors.New("test suite is fail")
	}

	return nil
}

func newSuite() *Suite {

	return &Suite{
		Format:      "pretty",
		Strict:      true,
		Concurrency: 1,
	}

}
