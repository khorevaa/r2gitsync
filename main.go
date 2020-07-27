package main

import (
	"flag"
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/internal/env"
	"github.com/khorevaa/r2gitsync/internal/opts"
	"os"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

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
}

var config *configApp
var pwd, _ = os.Getwd()

func main() {

	config = &configApp{}

	app := cli.App("r2gitsync", "Синхронизация 1С Хранилища с git")

	app.Version("Version v", buildVersion(version, commit, date, builtBy))

	var (
		debug = opts.BoolOpt("debug", false, "Bывод отладочной информации").
			Env(env.Versobe).
			Opt(app)
		v8version = opts.StringOpt("v8version", "8.3", "маска версии платформы 1С (8.3, 8.3.5, 8.3.6.2299 и т.п.)").
				Env(env.V8Version).
				Opt(app)
		v8path = opts.StringOpt("v8-path v8path", "", "путь к исполняемому файлу платформы 1С (Например, /opt/1C/v8.3/x86_64/1cv8)").
			Env(env.V8Path).
			Opt(app)
		ibUser = opts.StringOpt("U ib-Author ib-usr db-Author", "", "пользователь информационной базы").
			Env("GITSYNC_IB_USR GITSYNC_IB_USER GITSYNC_DB_USER").
			Opt(app)
		ibPassword = opts.StringOpt("P ib-pwd db-pwd", "", "пароль пользователя информационной базы").
				Env("GITSYNC_IB_PASSWORD GITSYNC_IB_PWD GITSYNC_DB_PSW").
				Opt(app)
		ibConnection = opts.StringOpt("C ib-connection ibconnection", "", "путь подключения к информационной базе").
				Env("GITSYNC_IB_CONNECTION GITSYNC_IBCONNECTION").
				Opt(app)
		tempDir = opts.StringOpt("t tempdir", "", "путь к каталогу временных файлов").
			Env("GITSYNC_TEMP GITSYNC_TEMPDIR").
			Opt(app)
	)

	app.Before = func() {

		config.debug = *debug
		config.v8version = *v8version
		config.v8path = *v8path
		config.tempDir = *tempDir

		if len(*ibConnection) > 0 {

			config.Infobase.ConnectionString = *ibConnection
			config.Infobase.User = *ibUser
			config.Infobase.Password = *ibPassword
		}

		if config.debug {
			fmt.Println("Включен режим отладки")
		}

	}

	app.ErrorHandling = flag.ExitOnError

	// Define our command structure for usage like this:
	app.Command("init i", "Инициализация структуры нового хранилища git", cmdInit)
	app.Command("sync s", "Выполняет синхронизацию хранилища 1С с git-репозиторием", cmdSync)
	app.Command("set-Version sv", "Устанавливает необходимую версию в файл VERSION", cmdSetVersion)
	app.Command("plugins p", "Управление плагинами", func(plugins *cli.Cmd) {
		plugins.Command("list ls", "Вывод списка плагинов", cmdPlugins)
		plugins.Command("enable e", "Активизация установленных плагинов", cmdPlugins)
		plugins.Command("disable d", "Деактивизация установленных плагинов", cmdPlugins)
		plugins.Command("install i", "Установка новых плагинов", cmdPlugins)
		plugins.Command("clear c", "Очистка установленных плагинов", cmdPlugins)
		plugins.Command("init", "Инициализация предустановленных плагинов", cmdPlugins)
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
