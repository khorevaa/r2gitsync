package main

import (
	"flag"
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/internal/env"
	"github.com/khorevaa/r2gitsync/internal/opts"
	"os"
	"path"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

type Application struct {
	*cli.Cli
	config *configApp
}

type configApp struct {
	debug     bool
	v8version string
	Infobase  struct {
		User             string
		Password         string
		ConnectionString string
	}

	v8path  string
	tempDir string

	pluginsLoader  *PluginsLoader
	pluginsManager *PluginsManager
	pluginsDir     string
}

var config = &configApp{}
var app = &Application{
	config: config,
}

var pwd, _ = os.Getwd()

func main() {

	app.Cli = cli.App("r2gitsync", "Синхронизация 1С Хранилища с git")

	app.Version("Version v", buildVersion(version, commit, date, builtBy))

	opts.StringOpt(app, "v8version", "8.3", "маска версии платформы 1С (8.3, 8.3.5, 8.3.6.2299 и т.п.)").
		Env(env.V8Version).
		Ptr(&config.v8version)
	opts.StringOpt(app, "v8-path v8path", "", "путь к исполняемому файлу платформы 1С (Например, /opt/1C/v8.3/x86_64/1cv8)").
		Env(env.V8Path).
		Ptr(&config.v8path)
	opts.StringOpt(app, "U ib-Author ib-usr db-Author", "", "пользователь информационной базы").
		Env("GITSYNC_IB_USR GITSYNC_IB_USER GITSYNC_DB_USER").
		Ptr(&config.Infobase.User)
	opts.StringOpt(app, "P ib-pwd db-pwd", "", "пароль пользователя информационной базы").
		Env("GITSYNC_IB_PASSWORD GITSYNC_IB_PWD GITSYNC_DB_PSW").
		Ptr(&config.Infobase.Password)
	opts.StringOpt(app, "C ib-connection ibconnection", "", "путь подключения к информационной базе").
		Env("GITSYNC_IB_CONNECTION GITSYNC_IBCONNECTION").
		Ptr(&config.Infobase.ConnectionString)
	opts.StringOpt(app, "t tempdir", "", "путь к каталогу временных файлов").
		Env("GITSYNC_TEMP GITSYNC_TEMPDIR").
		Ptr(&config.tempDir)
	opts.BoolOpt(app, "debug", false, "Bывод отладочной информации").
		Env(env.Versobe).
		Ptr(&config.debug)

	app.Before = func() {

		if config.debug {
			fmt.Println("Включен режим отладки")
		}

	}

	if len(config.pluginsDir) == 0 {
		appDataDir := getAppDataDir("r2gitsync")
		config.pluginsDir = path.Join(appDataDir, "plugins")
	}

	config.pluginsLoader = NewPluginsLoader(config.pluginsDir)

	config.pluginsManager = NewPluginsManager(ManagerConfig{
		enable: []string{
			"limit",
		},
	})
	err := config.pluginsManager.LoadPlugins(config.pluginsLoader)
	failOnErr(err)

	app.ErrorHandling = flag.ExitOnError

	// Define our command structure for usage like this:
	app.Command("init i", "Инициализация структуры нового хранилища git", app.cmdInit)
	app.Command("sync s", "Выполняет синхронизацию хранилища 1С с git-репозиторием", app.cmdSync)
	app.Command("set-Version sv", "Устанавливает необходимую версию в файл VERSION", app.cmdSetVersion)
	app.Command("plugins p", "Управление плагинами", func(pluginsCmd *cli.Cmd) {
		pluginsCmd.Command("list ls", "Вывод списка плагинов", app.cmdPlugins)
		pluginsCmd.Command("enable e", "Активизация установленных плагинов", app.cmdPlugins)
		pluginsCmd.Command("disable d", "Деактивизация установленных плагинов", app.cmdPlugins)
		pluginsCmd.Command("install i", "Установка новых плагинов", app.cmdPlugins)
		pluginsCmd.Command("clear c", "Очистка установленных плагинов", app.cmdPlugins)
		pluginsCmd.Command("init", "Инициализация предустановленных плагинов", app.cmdPlugins)
	})

	app.Run(os.Args)
}

func buildVersion(version, commit, date, builtBy string) string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}

func (app *Application) cmdInit(cmd *cli.Cmd) {

	cmd.LongDesc = `Данный режим работает по HTTP (REST API) с базой данных.
		Возможности:
		* самостоятельно получает список информационных баз к обновления;
		* поддержание нескольких потоков обновления
		* переодический/разовый опрос необходимости обновления
		* отправка журнала обновления на url.`

	cmd.Action = func() {
		//fmt.Printf("display account info for %s\n", *account)
	}
}
