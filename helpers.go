package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	cmd2 "github.com/khorevaa/r2gitsync/cmd"
	"github.com/khorevaa/r2gitsync/internal/args"
)

func WorkdirArg(cmd *cli.Cmd) *string {
	return args.StringArg(cmd, "WORKDIR", pwd, "Каталог исходников внутри локальной копии git-репозитория.").
		HideValue(true).
		Env(cmd2.WorkDir).
		Arg()
}

func WorkdirArgPtr(into *string, cmd *cli.Cmd) {
	args.StringArg(cmd, "WORKDIR", pwd, "Каталог исходников внутри локальной копии git-репозитория.").
		HideValue(true).
		Env(cmd2.WorkDir).
		Ptr(into)
}

func failOnErr(err error) {
	if err != nil {
		fmt.Printf("Ошибка выполненния программы: %v \n", err.Error())
		cli.Exit(1)
	}
}
