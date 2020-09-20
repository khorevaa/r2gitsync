package cmd

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/manager"
	"github.com/v8platform/designer/repository"
)

// Sample use: vault creds reddit.com
func (app *Application) CmdInit2(cmd *cli.Cmd) {

	app.config.Repository = new(repository.Repository)
	cmd.Spec = "[ -U=<storage-user> ] [ -P=<storage-pwd> ] PATH WORKDIR"
	cmd.LongDesc = `Инициализация структуры нового хранилища git. Подготовка к синхронизации`

	flags.StringOpt("U storage-user",
		"Администратор",
		"пользователь хранилища конфигурации").
		Env("GITSYNC_STORAGE_USER").
		Ptr(&app.config.Repository.User).Apply(cmd, app.ctx)

	flags.StringOpt("P storage-pwd",
		"",
		"пароль пользователя хранилища конфигурации").
		Env("GITSYNC_STORAGE_PASSWORD $GITSYNC_STORAGE_PWD").
		Ptr(&app.config.Repository.Password).Apply(cmd, app.ctx)

	//flags.StringOpt(" E ext extension",
	//	"",
	//	" имя расширения для работы с хранилищем расширения").
	//	Env("GITSYNC_EXTENSION").
	//	Ptr(&app.config.Repository.Extension).Apply(cmd, app.ctx)

	flags.StringArg("PATH",
		"",
		"Путь к хранилищу конфигурации 1С.").
		Env("GITSYNC_STORAGE_USER").
		Ptr(&app.config.Repository.Path).Apply(cmd, app.ctx)

	flags.StringArg("WORKDIR",
		".",
		" Адрес локального репозитория GIT.\n"+
			"Каталог исходников внутри локальной копии git-репозитория. По умолчанию текущий каталог").
		Env("GITSYNC_WORKDIR").
		Ptr(&app.config.workdir).Apply(cmd, app.ctx)

	cmd.Action = func() {
		initProject(app)
	}
}

func  initProject(app *Application) {
	CreateFileVersion(app)
}

func CreateFileVersion(app *Application) {
	// Get storage user
	repo := manager.SyncRepository{
		Repository: *app.config.Repository,
		Workdir: app.config.workdir,
	}
	repo.GetRepositoryHistory()

	fmt.Print("INIT!!")
	// Vrite storage user to Autors
}

func (app *Application) createFileAutors() {
	// Create file Version
}
