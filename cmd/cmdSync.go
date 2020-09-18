package cmd

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/manager"
	"github.com/khorevaa/r2gitsync/plugin"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
)

// Sample use: vault creds reddit.com
func (app *Application) cmdSync(cmd *cli.Cmd) {

	cmd.LongDesc = `Выполнение синхронизации Хранилища 1С с git репозиторием`

	repo := manager.SyncRepository{}

	flags.StringOpt("storage-author u", "Администратор", "пользователь хранилища 1C конфигурации").
		Env("R2GITSYNC_STORAGE_USER GITSYNC_STORAGE_USER").
		Ptr(&repo.Repository.User).
		Apply(cmd, app.ctx)

	flags.StringOpt("storage-pwd p", "", "пользователь хранилища 1C конфигурации").
		Env("R2GITSYNC_STORAGE_PASSWORD GITSYNC_STORAGE_PWD GITSYNC_STORAGE_PASSWORD").
		Ptr(&repo.Repository.Password).
		Apply(cmd, app.ctx)

	flags.BoolOpt("disable-increment", false, "отключает инкрементальную выгрузку").
		Env("GITSYNC_DISABLE_INCREMENT").
		Ptr(&app.config.disableIncrement).
		Apply(cmd, app.ctx)

	flags.StringOpt("extension e ext", "", "имя расширения для работы с хранилищем расширения").
		Env("R2GITSYNC_EXTENSION GITSYNC_EXTENSION").
		Ptr(&repo.Extention).
		Apply(cmd, app.ctx)

	flags.StringArg("PATH", "", "Путь к хранилищу конфигурации 1С.").
		Env("R2GITSYNC_STORAGE_PATH GITSYNC_STORAGE_PATH").
		Ptr(&repo.Repository.Path).
		Apply(cmd, app.ctx)

	//flags.StringArg("PATH", "", "Путь к хранилищу конфигурации 1С.").
	//	Env("R2GITSYNC_STORAGE_PATH", "GITSYNC_STORAGE_PATH").
	//	Ptr(&repo.Repository.Path).
	//	Apply(cmd, app.ctx)

	WorkdirArg.Ptr(&repo.Workdir).Apply(cmd, app.ctx)

	cmd.Spec = "[OPTIONS] PATH [WORKDIR]"

	var sm *subscription.SubscribeManager
	cmd.Before = func() {

		sm, _ = plugin.Subscribe("sync", app.ctx)
	}

	cmd.Action = func() {

		//err := manager.Sync(repo,
		//	manager.WithInfobaseConfig(app.config.Infobase),
		//	manager.WithTempDir(app.config.TempDir),
		//	manager.WithV8Path(app.config.v8path),
		//	manager.WithV8version(app.config.V8version),
		//	manager.WithLicTryCount(5),
		//	manager.WithPlugins(sm),
		//	manager.WithDisableIncrement(app.config.disableIncrement),
		//	//WithDomainEmail(config.),
		//)

		//failOnErr(err)

	}

	plugin.RegistryFlags("sync", cmd, app.ctx)

}
