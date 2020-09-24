package main

import (
	"errors"
	"github.com/cucumber/godog"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//format := "progress"
	//for _, arg := range os.Args[1:] {
	//	if arg == "-test.v=true" { // go test transforms -v option
	//		format = "pretty"
	//		break
	//	}
	//}

	//opts := godog.Options{
	//	Format: "pretty",
	//	Paths:     []string{"features"},
	//}
	//
	//
	//// godog v0.10.0 (latest)
	//status := godog.TestSuite{
	//	Name: "godogs",
	//	TestSuiteInitializer: InitializeTestSuite,
	//	ScenarioInitializer:  InitializeScenario,
	//	Options: &opts,
	//}.Run()
	//
	//if st := m.Run(); st > status {
	//	status = st
	//}

	//os.Exit(status)

	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestA(t *testing.T) {
	log.Println("TestA running")

	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"features/simple.feature"},
		//ShowStepDefinitions: true,
		//StopOnFailure: true,
		Strict: true,
	}

	// godog v0.10.0 (latest)
	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if status > 0 {
		t.Fail()
	}

}

func InitializeTestSuite(context *godog.TestSuiteContext) {

}

type simple struct {
	one, two int
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	simpleFeature := &simple{}

	ctx.Step(`^I add (\d+) and (\d+)$`, simpleFeature.iAddAnd)
	ctx.Step(`^I the result should equal (\d+)$`, simpleFeature.iTheResultShouldEqual)
}

func (s *simple) iAddAnd(arg1, arg2 int) error {

	s.one += arg1
	s.two += arg2

	return nil

}

func (s *simple) iTheResultShouldEqual(received int) error {

	sum := s.one + s.two

	if sum != received {
		return errors.New("the math does not work for you")
	}

	return nil
}
