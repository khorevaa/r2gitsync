package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/internal/args"
	"github.com/khorevaa/r2gitsync/internal/opts"
)

// Sample use: vault creds reddit.com
func cmdSync(cmd *cli.Cmd) {

	cmd.LongDesc = `Выполнение синхронизации Хранилища 1С с git репозиторием`

	repositoryUser := opts.StringOpt("storage-Author u", "Администратор", "пользователь хранилища 1C конфигурации").
		Env("R2GITSYNC_STORAGE_USER GITSYNC_STORAGE_USER").
		Opt(cmd)

	repositoryUserPassword := opts.StringOpt("storage-pwd p", "", "пользователь хранилища 1C конфигурации").
		Env("R2GITSYNC_STORAGE_PASSWORD GITSYNC_STORAGE_PWD GITSYNC_STORAGE_PASSWORD").
		Opt(cmd)

	extention := opts.StringOpt("extension e ext", "", "имя расширения для работы с хранилищем расширения").
		Env("R2GITSYNC_EXTENSION GITSYNC_EXTENSION").
		Opt(cmd)

	repositoryPath := args.StringArg("PATH", "", "Путь к хранилищу конфигурации 1С.").
		Env("R2GITSYNC_STORAGE_PATH GITSYNC_STORAGE_PATH").
		Arg(cmd)

	workdir := WorkdirArg(cmd)

	cmd.Spec = "[OPTIONS] PATH [WORKDIR]"

	cmd.Action = func() {
		fmt.Printf("Author %s\n"+
			"Password %s\n"+
			"Ext %s\n"+
			"Repository %s\n"+
			"Workdir %s\n", *repositoryUser, *repositoryUserPassword, *extention,
			*repositoryPath, *workdir)
	}
}
