package cmd

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"os"
)

var pwd, _ = os.Getwd()

var WorkdirArg = flags.StringFlag{
	FlagType: flags.ArgType,
	Name:     "WORKDIR",
	Desc:     "Каталог исходников внутри локальной копии git-репозитория.",
	EnvVar:   WorkDirEnv,
	Value:    pwd,
}

func failOnErr(err error) {
	if err != nil {
		fmt.Printf("Ошибка выполненния программы: %v \n", err.Error())
		cli.Exit(1)
	}
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
func IsNoExist(name string) (bool, error) {

	ok, err := Exists(name)
	return !ok, err
}
