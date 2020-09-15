package cmd

import (
	"flag"
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/plugin"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	disabledPluginsFileName = "disabled-plugins"
	appDirName              = ".r2gitsync"
)

var (
	appDirPwd     = path.Join(pwd, appDirName)
	pluginsDirPwd = path.Join(appDirPwd, "plugins")
)

type Application struct {
	*cli.Cli
	config *configApp
	ctx    context.Context
}

type configApp struct {
	debug     bool
	v8version string
	Infobase  struct {
		User             string
		Password         string
		ConnectionString string
	}

	v8path           string
	tempDir          string
	workspace        string
	disableIncrement bool
}

func NewApp(version string) *Application {

	config := &configApp{}

	app := &Application{
		config: config,
	}

	loadPlugins()
	disablePlugins()

	app.Cli = cli.App("r2gitsync", "Синхронизация 1С Хранилища с git")
	app.ctx = context.NewContext()

	app.Version("version v", version)

	flags.StringOpt("v8version", "8.3", "маска версии платформы 1С (8.3, 8.3.5, 8.3.6.2299 и т.п.)").
		Env(V8VersionEnv).
		Ptr(&config.v8version).Apply(app, app.ctx)
	flags.StringOpt("ws workspace", "", "рабочая область приложения").
		//Env(V8VersionEnv).
		Ptr(&config.workspace).Apply(app, app.ctx)
	flags.StringOpt("v8-path v8path", "", "путь к исполняемому файлу платформы 1С (Например, /opt/1C/v8.3/x86_64/1cv8)").
		Env(V8PathEnv).
		Ptr(&config.v8path).Apply(app, app.ctx)
	flags.StringOpt("U ib-author ib-usr db-author", "", "пользователь информационной базы").
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
		Ptr(&config.tempDir).Apply(app, app.ctx)
	flags.BoolOpt("debug", false, "Bывод отладочной информации").
		Env(VersobeEnv).
		Ptr(&config.debug).Apply(app, app.ctx)

	app.Before = func() {

		if config.debug {
			fmt.Println("Включен режим отладки")
		}

		//err := plugin.Subscribe("", app.ctx)

		//failOnErr(err)

	}

	app.After = func() {}

	app.ErrorHandling = flag.ExitOnError

	// Define our command structure for usage like this:
	app.Command("init i", "Инициализация структуры нового хранилища git", app.cmdInit)
	app.Command("sync s", "Выполняет синхронизацию хранилища 1С с git-репозиторием", app.cmdSync)
	app.Command("set-version sv", "Устанавливает необходимую версию в файл VERSION", app.cmdSetVersion)
	app.Command("plugins p", "Управление плагинами", func(pluginsCmd *cli.Cmd) {
		pluginsCmd.Command("list ls", "Вывод списка плагинов", app.cmdPluginsList)
		pluginsCmd.Command("enable e", "Активизация установленных плагинов", app.cmdPlugins)
		pluginsCmd.Command("disable d", "Деактивизация установленных плагинов", app.cmdPlugins)
		pluginsCmd.Command("install i", "Установка новых плагинов", app.cmdPlugins)
		pluginsCmd.Command("clear c", "Очистка установленных плагинов", app.cmdPlugins)
		pluginsCmd.Command("init", "Инициализация предустановленных плагинов", app.cmdPlugins)
	})

	return app
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

func loadPlugins() {

	pluginsDir := getEnv(PluginsDirEnv)

	if len(pluginsDir) == 0 {
		appDataDir := getAppDataDir("r2gitsync")
		pluginsDir = path.Join(appDataDir, "plugins")
	}

	if ok, _ := IsNoExist(pluginsDir); ok {
		return
	}

	err := plugin.LoadPlugins(pluginsDir)
	failOnErr(err)

}

func getEnv(envs ...string) string {

	for _, env := range envs {

		keys := strings.Fields(env)

		for _, key := range keys {
			value := strings.TrimSpace(os.Getenv(key))

			if len(value) > 0 {
				return value
			}
		}

	}

	return ""

}

func disablePlugins() {

	pl := getEnv("R2GITSYNC_DISABLE_PLUGINS")
	plugin.Disable(strings.Split(pl, ",")...)

	filename := path.Join(appDirPwd, disabledPluginsFileName)
	if ok, _ := Exists(filename); ok {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			failOnErr(err)
		}
		lines := strings.Split(string(content), "\n")
		plugin.Disable(lines...)

	}

}
