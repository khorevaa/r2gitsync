package app

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/khorevaa/r2gitsync/pkg/context"
	"strings"
)

func (c *AppTestSuite) RunApp() (err error) {

	defer func() {
		if r := recover(); r != nil {

			switch e := r.(type) {

			case string:
				err = errors.New(e)
			case error:
				err = e
				if err != nil && strings.Contains(err.Error(), "incorrect usage") {

					err = errors.New(fmt.Sprintf("%s: %s", err.Error(), strings.Join(c.Args, " ")))
				}
			}

		}
	}()

	err = c.Run(c.Args)

	return err
}

func (c *AppTestSuite) RunAppWithErr(errStr *messages.PickleStepArgument_PickleDocString) (err error) {

	err = c.RunApp()
	data := errStr.GetContent()
	if err != nil && strings.Contains(err.Error(), data) {
		return nil
	}

	return errors.New("не обнаружена искомая ошибка выполнения")
}

type AppTestSuite struct {
	*Application
	context.Context

	Args []string
}

func SharedAppSteps(sharedContext context.Context, ctx *godog.ScenarioContext) *AppTestSuite {

	step := &AppTestSuite{
		Application: NewApp("dev", false),
		Context:     sharedContext,
	}

	step.ErrorHandling = flag.PanicOnError
	step.PanicOnErr = true

	ctx.Step(`^Я выполняю приложение$`, step.RunApp)
	ctx.Step(`^Я выполняю приложение c ошибкой:$`, step.RunAppWithErr)
	ctx.Step(`^Я добавляю параметр "([^"]*)" из переменной "([^"]*)"$`, step.AddArgNamedFromContext)
	ctx.Step(`^Я добавляю параметр "([^"]*)"$`, step.AddArg)
	ctx.Step(`^Я добавляю параметр из переменной "([^"]*)"$`, step.AddArgFromContext)
	ctx.Step(`^Я добавляю параметр без равно "([^"]*)" из переменной "([^"]*)"$`, step.AddArgNamedFromContextWithOutSep)

	return step

}

func (c *AppTestSuite) AddArgNamedFromContext(arg, name string) error {

	value := c.String(name)

	c.Args = append(c.Args, fmt.Sprintf("%s=%s", arg, value))

	return nil
}

func (c *AppTestSuite) AddArgNamedFromContextWithOutSep(arg, name string) error {

	value := c.String(name)

	c.Args = append(c.Args, fmt.Sprintf("%s%s", arg, value))

	return nil
}

func (c *AppTestSuite) AddArg(arg1 string) error {

	c.Args = append(c.Args, arg1)
	return nil

}

func (c *AppTestSuite) AddArgFromContext(name string) error {

	value := c.String(name)

	c.Args = append(c.Args, value)

	return nil
}
