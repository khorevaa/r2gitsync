package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/internal/args"
	"github.com/khorevaa/r2gitsync/internal/env"
)

func WorkdirArg(cmd *cli.Cmd) *string {
	return args.StringArg(cmd, "WORKDIR", pwd, "Каталог исходников внутри локальной копии git-репозитория.").
		HideValue(true).
		Env(env.WorkDir).
		Arg()
}

func WorkdirArgPtr(into *string, cmd *cli.Cmd) {
	args.StringArg(cmd, "WORKDIR", pwd, "Каталог исходников внутри локальной копии git-репозитория.").
		HideValue(true).
		Env(env.WorkDir).
		Ptr(into)
}

func failOnErr(err error) {
	if err != nil {
		fmt.Printf("Ошибка выполненния программы: %v \n", err.Error())
		cli.Exit(1)
	}
}
