package app

import (
	"errors"
	cli "github.com/jawher/mow.cli"
	flags2 "github.com/khorevaa/r2gitsync/internal/app/flags"
	"github.com/khorevaa/r2gitsync/internal/log"
	manager2 "github.com/khorevaa/r2gitsync/internal/manager"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
)

func (app *Application) cmdSync(cmd *cli.Cmd) {

	cmd.LongDesc = `Выполнение синхронизации Хранилища 1С с git репозиторием`

	repo := manager2.SyncRepository{}

	flags2.StringOpt("storage-author u", "Администратор", "пользователь хранилища 1C конфигурации").
		Env("R2GITSYNC_STORAGE_USER GITSYNC_STORAGE_USER").
		Ptr(&repo.Repository.User).
		Apply(cmd, app.ctx)

	flags2.StringOpt("storage-pwd p", "", "пользователь хранилища 1C конфигурации").
		Env("R2GITSYNC_STORAGE_PASSWORD GITSYNC_STORAGE_PWD GITSYNC_STORAGE_PASSWORD").
		Ptr(&repo.Repository.Password).
		Apply(cmd, app.ctx)

	flags2.BoolOpt("disable-increment", false, "отключает инкрементальную выгрузку").
		Env("GITSYNC_DISABLE_INCREMENT").
		Ptr(&app.config.Options.DisableIncrement).
		Apply(cmd, app.ctx)

	flags2.StringOpt("extension e ext", "", "имя расширения для работы с хранилищем расширения").
		Env("R2GITSYNC_EXTENSION GITSYNC_EXTENSION").
		Ptr(&repo.Extension).
		Apply(cmd, app.ctx)

	flags2.StringArg("PATH", "", "Путь к хранилищу конфигурации 1С.").
		Env("R2GITSYNC_STORAGE_PATH GITSYNC_STORAGE_PATH").
		Ptr(&repo.Repository.Path).
		Apply(cmd, app.ctx)

	flags2.IntOpt(
		"l limit",
		0,
		"выгрузить не более <Количества> версий от текущей выгруженной").
		Env("GITSYNC_LIMIT").
		Ptr(&app.config.Options.LimitVersions).
		Apply(cmd, app.ctx)
	flags2.IntOpt(
		"min-version",
		0,
		"<номер> минимальной версии для выгрузки").
		Env("GITSYNC_MIN_VERSION").
		Ptr(&app.config.Options.MinVersion).
		Apply(cmd, app.ctx)
	flags2.IntOpt(
		"max-version",
		0,
		"<номер> максимальной версии для выгрузки").
		Env("GITSYNC_MAX_VERSION").
		Ptr(&app.config.Options.MaxVersion).
		Apply(cmd, app.ctx)

	//flags.StringArg("PATH", "", "Путь к хранилищу конфигурации 1С.").
	//	Env("R2GITSYNC_STORAGE_PATH", "GITSYNC_STORAGE_PATH").
	//	Ptr(&repo.Repository.Path).
	//	Apply(cmd, app.ctx)

	WorkdirArg.Ptr(&repo.Workdir).Apply(cmd, app.ctx)

	cmd.Spec = "[OPTIONS] [PATH] [WORKDIR]"

	var syncOptions *manager2.Options
	cmd.Before = func() {

		sm, err := plugin.Subscribe("sync", app.ctx)

		if err != nil {
			app.failOnErr(err)
		}

		logger := log.Named("cmd")
		logger.Debug("New logger inited")
		newOptions := *app.config.Options
		syncOptions = &newOptions
		syncOptions.Logger = logger
		syncOptions.Plugins = sm
		syncOptions.LicTryCount = 5

	}

	cmd.Action = func() {

		if len(repo.Repository.Path) == 0 {

			app.failOnErr(errors.New("путь к репозиторию должен быть установлен"))

		}

		err := manager2.Sync(repo, *syncOptions)

		app.failOnErr(err)

	}

	plugin.RegistryFlags("sync", cmd, app.ctx)

}
