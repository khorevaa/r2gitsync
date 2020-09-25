package cmd

import (
	"flag"
	"github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/log"
	"path/filepath"
)

const (
	appDirName = ".r2gitsync"
)

var (
	appDirPwd     = filepath.Join(pwd, appDirName)
	pluginsDirPwd = filepath.Join(appDirPwd, "plugins")
)

type Application struct {
	*cli.Cli
	config *configApp
	ctx    context.Context

	PanicOnErr  bool
	initPlugins bool
}

type configApp struct {
	Debug     bool
	V8version string
	Infobase  struct {
		User             string
		Password         string
		ConnectionString string
	}

	v8path           string
	TempDir          string
	Workspace        string
	disableIncrement bool
	DomainEmail      string
	Plugins          struct {
		GlobalDir string
		LocalDir  string
	}
}

func NewApp(version string, initPlugins bool) *Application {

	config := &configApp{}

	if initPlugins {

		initPluginsDirs(config)
		loadPlugins(config)
		loadDisabledPlugins(config)

	}
	app := &Application{
		config: config,
	}

	app.Cli = cli.App("r2gitsync", "Синхронизация 1С Хранилища с git")
	app.ctx = context.NewContext()

	app.Version("version v", version)

	flags.StringOpt("V8version", "8.3", "маска версии платформы 1С (8.3, 8.3.5, 8.3.6.2299 и т.п.)").
		Env(V8VersionEnv).
		Ptr(&config.V8version).Apply(app, app.ctx)
	flags.StringOpt("ws Workspace", "", "рабочая область приложения").
		//Env(V8VersionEnv).
		Ptr(&config.Workspace).Apply(app, app.ctx)
	flags.StringOpt("v8-path v8path", "", "путь к исполняемому файлу платформы 1С (Например, /opt/1C/v8.3/x86_64/1cv8)").
		Env(V8PathEnv).
		Ptr(&config.v8path).Apply(app, app.ctx)
	flags.StringOpt("U ib-user ib-usr", "", "пользователь информационной базы").
		Env("GITSYNC_IB_USR GITSYNC_IB_USER GITSYNC_DB_USER").
		Ptr(&config.Infobase.User).Apply(app, app.ctx)
	flags.StringOpt("P ib-pwd db-pwd", "", "пароль пользователя информационной базы").
		Env("GITSYNC_IB_PASSWORD GITSYNC_IB_PWD GITSYNC_DB_PSW").
		Ptr(&config.Infobase.Password).Apply(app, app.ctx)
	flags.StringOpt("C ib-connection ibconnection", "", "путь подключения к информационной базе").
		Env("GITSYNC_IB_CONNECTION GITSYNC_IBCONNECTION").
		Ptr(&config.Infobase.ConnectionString).Apply(app, app.ctx)
	flags.StringOpt("t tempdir", "", "путь к каталогу временных файлов").
		Env("GITSYNC_TEMP GITSYNC_TEMPDIR").
		Ptr(&config.TempDir).Apply(app, app.ctx)
	flags.BoolOpt("debug", false, "вывод отладочной информации").
		Env(VersobeEnv).
		Ptr(&config.Debug).Apply(app, app.ctx)

	app.Before = func() {

		if config.Debug {
			log.SetDebug()
			log.Debugw("Включен режим отладки", "config", config)
		}

		//err := plugin.Subscribe("", app.ctx)

		//failOnErr(err)

	}

	app.After = func() {
		log.Debug("after app")

	}

	app.ErrorHandling = flag.ExitOnError

	// Define our command structure for usage like this:
	app.Command("init i", "Инициализация структуры нового хранилища git", app.cmdInit)
	app.Command("sync s", "Выполняет синхронизацию хранилища 1С с git-репозиторием", app.cmdSync)
	app.Command("set-version sv", "Устанавливает необходимую версию в файл VERSION", app.cmdSetVersion)
	app.Command("plugins p", "Управление плагинами", func(pluginsCmd *cli.Cmd) {
		pluginsCmd.Command("list ls", "Вывод списка плагинов", app.cmdPluginsList)
		pluginsCmd.Command("enable e", "Активизация установленных плагинов", app.cmdPluginsEnable)
		pluginsCmd.Command("disable d", "Деактивизация установленных плагинов", app.cmdPluginsDisable)
		pluginsCmd.Command("install i", "Установка новых плагинов", app.cmdPluginsInstall)
		pluginsCmd.Command("clear c", "Очистка установленных плагинов", app.cmdPluginsClear)
	})

	return app
}

func initPluginsDirs(config *configApp) {

	appDataDir := getAppDataDir("r2gitsync")
	config.Plugins.GlobalDir = filepath.Join(appDataDir, "plugins")

	localDir := getEnv(PluginsDirEnv)

	if len(localDir) == 0 {
		localDir = pluginsDirPwd
	}

	config.Plugins.LocalDir = localDir

}

func loadPlugins(config *configApp) {

	loadInternalPlugins()
	loadGlobalPlugins(config.Plugins.GlobalDir)
	loadLocalPlugins(config.Plugins.LocalDir)

}

func loadDisabledPlugins(config *configApp) {

	loadGlobalDisabledPlugins(config.Plugins.GlobalDir)
	loadLocalDisabledPlugins(config.Plugins.LocalDir)
	loadLocalEnabledPlugins(config.Plugins.LocalDir)
	loadDisabledPluginsEnv()

}
