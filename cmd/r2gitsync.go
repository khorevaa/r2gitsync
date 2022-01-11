package main

import (
	"fmt"
	"strings"

	"github.com/khorevaa/r2gitsync/internal/commands"
	"github.com/khorevaa/r2gitsync/internal/config"
	_ "github.com/khorevaa/r2gitsync/internal/plugins"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
	"github.com/urfave/cli/v2"

	"os"

	"github.com/khorevaa/logos"
)

const (
	VerboseEnv    = "GITSYNC_VERBOSE R2GITSYNC_VERBOSE"
	V8VersionEnv  = "GITSYNC_V8_VERSION GITSYNC_V8VERSION"
	V8PathEnv     = "GITSYNC_V8_PATH"
	WorkDirEnv    = "R2GITSYNC_WORKDIR GITSYNC_WORKDIR"
	PluginsDirEnv = "GITSYNC_PLUGINS_PATH GITSYNC_PLUGINS_DIR GITSYNC_PL_DIR"
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

	appConfig := config.DefaultConfig()

	manager := plugin.NewPluginManager(appConfig.Plugins)

	app := &cli.App{
		Name:    "r2gitsync",
		Usage:   "Приложение для синхронизация Хранилища 1С с git",
		Version: buildVersion(),
		Flags:   setFlags(appConfig),
		Before: func(c *cli.Context) error {
			var err error

			if appConfig, err = config.LoadConfig(appConfig, c.String("config")); err != nil {
				return err
			}

			if appConfig.Debug {
				// TODO Не работает
				logos.SetLevel("github.com/khorevaa/r2gitsync", logos.DebugLevel)
			}

			manager = plugin.NewPluginManager(appConfig.Plugins)
			c.App.Commands = []*cli.Command{}
			for _, command := range commands.Commands {
				c.App.Commands = append(c.App.Commands, command.Cmd(manager))
			}

			return nil
		},
	}

	for _, command := range commands.Commands {
		app.Commands = append(app.Commands, command.Cmd(manager))
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

func setFlags(cfg *config.Config) []cli.Flag {

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "V8version",
			Value:       "8.3",
			EnvVars:     strings.Fields(V8VersionEnv),
			Usage:       "маска версии платформы 1С (8.3, 8.3.5, 8.3.6.2299 и т.п.)",
			Destination: &cfg.V8version,
		},
		&cli.StringFlag{
			Name:        "v8-path",
			Aliases:     []string{"v8path"},
			EnvVars:     strings.Fields(V8PathEnv),
			Usage:       "путь к исполняемому файлу платформы 1С (Например, /opt/1C/v8.3/x86_64/1cv8)",
			Destination: &cfg.V8Path,
		},
		&cli.StringFlag{
			Name:        "temp-dir",
			Aliases:     strings.Fields("t tempdir"),
			EnvVars:     strings.Fields("GITSYNC_TEMP GITSYNC_TEMPDIR"),
			Usage:       "путь к каталогу временных файлов",
			Destination: &cfg.TempDir,
		},

		&cli.StringFlag{
			Name:    "config",
			EnvVars: strings.Fields("GITSYNC_CONFIG"),
			Usage:   "путь к файлу настройки приложения",
		},

		&cli.BoolFlag{
			Name:        "debug",
			EnvVars:     strings.Fields(VerboseEnv),
			Usage:       "Режим отладки приложения",
			Destination: &cfg.Debug,
		},
	}
}
