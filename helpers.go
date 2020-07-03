package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
)

func StringOpt(name, value, desc, env string) cli.StringOpt {
	return cli.StringOpt{
		Name:   name,
		Value:  value,
		Desc:   desc,
		EnvVar: env,
	}
}

func BoolOpt(name string, value bool, desc, env string) cli.BoolOpt {
	return cli.BoolOpt{
		Name:   name,
		Value:  value,
		Desc:   desc,
		EnvVar: env,
	}
}

func StringArg(name, value, desc, env string) cli.StringArg {
	return cli.StringArg{
		Name:   name,
		Value:  value,
		Desc:   desc,
		EnvVar: env,
	}
}

func WorkdirArg() cli.StringArg {
	return cli.StringArg{
		Name:      "WORKDIR",
		Value:     pwd,
		HideValue: true,
		Desc:      "Каталог исходников внутри локальной копии git-репозитория.",
		EnvVar:    "R2GITSYNC_WORKDIR GITSYNC_WORKDIR"}
}

func failOnErr(err error) {
	if err != nil {
		fmt.Printf("Ошибка выполненния программы: %v \n", err.Error())
		cli.Exit(1)
	}
}
