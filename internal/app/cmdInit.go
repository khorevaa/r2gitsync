package app

import (
	cli "github.com/jawher/mow.cli"
	flags2 "github.com/khorevaa/r2gitsync/internal/app/flags"
	"github.com/khorevaa/r2gitsync/internal/log"
	manager2 "github.com/khorevaa/r2gitsync/internal/manager"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
)

func (app *Application) cmdInit(cmd *cli.Cmd) {

	var force bool

	cmd.LongDesc = `Инициализация структуры нового хранилища git. Подготовка к синхронизации`

	repo := manager2.SyncRepository{}

	flags2.StringOpt("storage-author u", "Администратор", "пользователь хранилища 1C конфигурации").
		Env("R2GITSYNC_STORAGE_USER GITSYNC_STORAGE_USER").
		Ptr(&repo.Repository.User).
		Apply(cmd, app.ctx)

	flags2.StringOpt("storage-pwd p", "", "пользователь хранилища 1C конфигурации").
		Env("R2GITSYNC_STORAGE_PASSWORD GITSYNC_STORAGE_PWD GITSYNC_STORAGE_PASSWORD").
		Ptr(&repo.Repository.Password).
		Apply(cmd, app.ctx)

	flags2.StringOpt("extension e ext", "", "имя расширения для работы с хранилищем расширения").
		Env("R2GITSYNC_EXTENSION GITSYNC_EXTENSION").
		Ptr(&repo.Extension).
		Apply(cmd, app.ctx)

	flags2.BoolOpt("f force", false, "принудительноя инициализация, игнорирование ошибок").
		Ptr(&force).
		Apply(cmd, app.ctx)

	flags2.StringArg("PATH", "", "Путь к хранилищу конфигурации 1С.").
		Env("R2GITSYNC_STORAGE_PATH GITSYNC_STORAGE_PATH").
		Ptr(&repo.Repository.Path).
		Apply(cmd, app.ctx)

	WorkdirArg.Ptr(&repo.Workdir).Apply(cmd, app.ctx)

	cmd.Spec = "[OPTIONS] PATH [WORKDIR]"

	var syncOptions *manager2.Options
	cmd.Before = func() {

		sm, err := plugin.Subscribe("init", app.ctx)

		if err != nil {
			app.failOnErr(err)
		}

		logger := log.Named("cmd")
		newOptions := *app.config.Options
		syncOptions = &newOptions
		syncOptions.Logger = logger
		syncOptions.Plugins = sm
		syncOptions.LicTryCount = 5

	}
	cmd.Action = func() {

		err := manager2.Init(repo, *syncOptions)

		app.failOnErr(err)

	}

	plugin.RegistryFlags("init", cmd, app.ctx)

}
