package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/internal/app"
	"github.com/khorevaa/r2gitsync/internal/log"
	"os"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

func main() {

	App := app.NewApp(buildVersion(), true)
	err := App.Run(os.Args)
	failOnErr(err)

}

func failOnErr(err error) {
	if err != nil {
		log.Errorf("Ошибка выполненния программы: %v \n", err.Error())
		cli.Exit(1)
	}
}

func buildVersion() string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
