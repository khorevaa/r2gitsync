package cmd

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
)

// Sample use: vault creds reddit.com
func (app *Application) CmdInit2(cmd *cli.Cmd) {

	cmd.Spec = "[ -U=<storage-user> ] [ -P=<storage-pwd> ] [ -E=<extension> ] PATH WORKDIR"
	cmd.LongDesc = `Инициализация структуры нового хранилища git. Подготовка к синхронизации`

	flags.StringOpt("U storage-user",
		"Администратор",
		"пользователь хранилища конфигурации").
		Env("GITSYNC_STORAGE_USER").
		Ptr(&app.config.Storage.User).Apply(cmd, app.ctx)

	flags.StringOpt("P storage-pwd",
		"",
		"пароль пользователя хранилища конфигурации").
		Env("GITSYNC_STORAGE_PASSWORD $GITSYNC_STORAGE_PWD").
		Ptr(&app.config.Storage.Password).Apply(cmd, app.ctx)

	flags.StringOpt(" E ext extension",
		"",
		" имя расширения для работы с хранилищем расширения").
		Env("GITSYNC_EXTENSION").
		Ptr(&app.config.Storage.Extension).Apply(cmd, app.ctx)

	flags.StringArg("PATH",
		"",
		"Путь к хранилищу конфигурации 1С.").
		Env("GITSYNC_STORAGE_USER").
		Ptr(&app.config.Storage.Path).Apply(cmd, app.ctx)

	flags.StringArg("WORKDIR",
		".",
		" Адрес локального репозитория GIT.\n"+
			"Каталог исходников внутри локальной копии git-репозитория. По умолчанию текущий каталог").
		Env("GITSYNC_WORKDIR").
		Ptr(&app.config.workdir).Apply(cmd, app.ctx)

	cmd.Action = func() {
		initProject()
	}
}

func initProject() {
	createFileAutors()
	CreateFileVersion()
}

func CreateFileVersion() {
	// Get storage user

	// Vrite storage user to Autors
}

func createFileAutors() {
	// Create file Version
}
