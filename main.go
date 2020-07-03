package main

import (
	"flag"
	"fmt"
	"github.com/jawher/mow.cli"
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

	app.Version("version v", buildVersion(version, commit, date, builtBy))

	var (
		debug        = app.Bool(BoolOpt("debug", false, "dывод отладочной информации", "GITSYNC_VERBOSE"))
		v8version    = app.String(StringOpt("v8version", "8.3", "маска версии платформы 1С (8.3, 8.3.5, 8.3.6.2299 и т.п.)", "GITSYNC_V8_VERSION GITSYNC_V8VERSION"))
		v8path       = app.String(StringOpt("v8-path v8path", "8.3", "путь к исполняемому файлу платформы 1С (Например, /opt/1C/v8.3/x86_64/1cv8)", "GITSYNC_V8_PATH"))
		ibUser       = app.String(StringOpt("U ib-user ib-usr db-user", "", "пользователь информационной базы", "GITSYNC_IB_USR GITSYNC_IB_USER GITSYNC_DB_USER"))
		ibPassword   = app.String(StringOpt("P ib-pwd db-pwd", "", "пароль пользователя информационной базы", "GITSYNC_IB_PASSWORD GITSYNC_IB_PWD GITSYNC_DB_PSW"))
		ibConnection = app.String(StringOpt("C ib-connection ibconnection", "", "путь подключения к информационной базе", "GITSYNC_IB_CONNECTION GITSYNC_IBCONNECTION"))
		tempDir      = app.String(StringOpt("t tempdir", "", "путь к каталогу временных файлов", "GITSYNC_TEMP GITSYNC_TEMPDIR"))
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
	app.Command("set-version sv", "Устанавливает необходимую версию в файл VERSION", cmdSetVersion)
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

// Sample use: vault list OR vault config list
func cmdList(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("list the contents of the safe here")
	}
}

// Sample use: vault list OR vault config list
func cmdSetVersion(cmd *cli.Cmd) {

	var (
		doCommit   = cmd.Bool(BoolOpt("c commit", false, "закоммитить изменения в git", ""))
		setVersion = cmd.String(StringArg("VERSION", "", "Номер версии для записи в файл.", ""))
		workdir    = cmd.String(WorkdirArg())
	)

	cmd.Spec = "[OPTIONS] VERSION [WORKDIR]"

	cmd.Action = func() {

		err := writeVersionFile(*workdir, *setVersion)

		failOnErr(err)

		if *doCommit {
			err = commitVersionFile(*workdir)
			failOnErr(err)
		}

	}
}

// Sample use: vault list OR vault config list
func cmdPlugins(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("list the contents of the safe here")
	}
}

// Sample use: vault creds reddit.com
func cmdInit(cmd *cli.Cmd) {

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

// Sample use: vault creds reddit.com
func cmdSync(cmd *cli.Cmd) {

	cmd.LongDesc = `Выполнение синхронизации Хранилища 1С с git репозиторием`

	repositoryUser := cmd.String(StringOpt(
		"storage-user u",
		"Администратор",
		"пользователь хранилища 1C конфигурации",
		"R2GITSYNC_STORAGE_USER GITSYNC_STORAGE_USER"))

	repositoryUserPassword := cmd.String(StringOpt(
		"storage-pwd p",
		"",
		"пользователь хранилища 1C конфигурации",
		"R2GITSYNC_STORAGE_PASSWORD GITSYNC_STORAGE_PWD GITSYNC_STORAGE_PASSWORD"))

	extention := cmd.String(StringOpt(
		"extension e ext",
		"",
		"имя расширения для работы с хранилищем расширения",
		"R2GITSYNC_EXTENSION GITSYNC_EXTENSION"))

	repositoryPath := cmd.String(StringArg(
		"PATH",
		"",
		"Путь к хранилищу конфигурации 1С.",
		"R2GITSYNC_STORAGE_PATH GITSYNC_STORAGE_PATH"))

	workdir := cmd.String(WorkdirArg())

	cmd.Spec = "[OPTIONS] PATH [WORKDIR]"

	cmd.Action = func() {
		fmt.Printf("User %s\n"+
			"Password %s\n"+
			"Ext %s\n"+
			"Repository %s\n"+
			"Workdir %s\n", *repositoryUser, *repositoryUserPassword, *extention,
			*repositoryPath, *workdir)
	}
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
