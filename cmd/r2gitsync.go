package main

import (
	"fmt"
	"github.com/khorevaa/r2gitsync/internal/app/cmd"
	"github.com/khorevaa/r2gitsync/internal/config"
	"github.com/urfave/cli/v2"
	"strings"

	"github.com/khorevaa/logos"
	"os"
)

const (
	VerboseEnv         = "GITSYNC_VERBOSE R2GITSYNC_VERBOSE"
	V8VersionEnv       = "GITSYNC_V8_VERSION GITSYNC_V8VERSION"
	V8PathEnv          = "GITSYNC_V8_PATH"
	WorkDirEnv         = "R2GITSYNC_WORKDIR GITSYNC_WORKDIR"
	PluginsDirEnv      = "GITSYNC_PLUGINS_PATH GITSYNC_PLUGINS_DIR GITSYNC_PL_DIR"
	DisabledPluginsEnv = "R2GITSYNC_DISABLED_PLUGINS"
	EnabledPluginsEnv  = "R2GITSYNC_ENABLED_PLUGINS"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

var log = logos.New("github.com/khorevaa/r2gitsync").Sugar()

func main() {

	globalConfig := &config.Config{}

	app := &cli.App{
		Name:    "r2gitsync",
		Usage:   "Приложение для синхронизация Хранилища 1С с git",
		Version: buildVersion(),
		Flags:   setFlags(),
		Before: func(c *cli.Context) error {

			debug := c.Bool("debug")

			if debug {
				logos.SetLevel("github.com/khorevaa/r2gitsync", logos.DebugLevel)
			}

			return nil
		},
	}

	for _, command := range cmd.Commands {
		app.Commands = append(app.Commands, command.Cmd())
	}

	err := app.Run(os.Args)
	defer log.Sync()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func buildVersion() string {
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

func setFlags() []cli.Flag {

	return []cli.Flag{
		&cli.StringFlag{
			Name:    "V8version",
			Value:   "8.3",
			EnvVars: strings.Fields(V8VersionEnv),
			Usage:   "маска версии платформы 1С (8.3, 8.3.5, 8.3.6.2299 и т.п.)",
		},
		&cli.StringFlag{
			Name:    "v8-path",
			Aliases: []string{"v8path"},
			EnvVars: strings.Fields(V8PathEnv),
			Usage:   "путь к исполняемому файлу платформы 1С (Например, /opt/1C/v8.3/x86_64/1cv8)",
		},
		&cli.StringFlag{
			Name:    "ib-user",
			Aliases: []string{"U ib-usr"},
			EnvVars: strings.Fields("GITSYNC_IB_USR GITSYNC_IB_USER GITSYNC_DB_USER"),
			Usage:   "пользователь информационной базы",
		},
		&cli.StringFlag{
			Name:    "ib-password",
			Aliases: []string{"P ib-pwd"},
			EnvVars: strings.Fields("GITSYNC_IB_PASSWORD GITSYNC_IB_PWD GITSYNC_DB_PSW"),
			Usage:   "пароль пользователя информационной базы",
		},
		&cli.StringFlag{
			Name:    "ib-connection",
			Aliases: []string{"C ib-connection ibconnection"},
			EnvVars: strings.Fields("GITSYNC_IB_CONNECTION GITSYNC_IBCONNECTION"),
			Usage:   "путь подключения к информационной базе",
		},
		&cli.StringFlag{
			Name:    "temp-dir",
			Aliases: []string{"t tempdir"},
			EnvVars: strings.Fields("GITSYNC_TEMP GITSYNC_TEMPDIR"),
			Usage:   "путь к каталогу временных файлов",
		},

		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{},
			EnvVars: strings.Fields("GITSYNC_CONFIG"),
			Usage:   "путь к файлу настройки приложения",
		},

		&cli.BoolFlag{
			Name:    "debug",
			EnvVars: strings.Fields(VerboseEnv),
			Usage:   "Режим отладки приложения",
		},
	}
}
