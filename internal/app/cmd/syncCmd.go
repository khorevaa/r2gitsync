package cmd

import (
	"fmt"
	manager2 "github.com/khorevaa/r2gitsync/internal/manager"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
	"github.com/khorevaa/r2gitsync/pkg/plugin/subscription"
	"github.com/urfave/cli/v2"
	"strings"
)

type syncCmd struct {
	MinVersion       int
	MaxVersion       int
	LimitVersions    int
	InfobaseConnect  string
	InfobaseUser     string
	InfobasePassword string
	StorageUser      string
	StoragePassword  string
	Extension        string

	DisableIncrement bool

	DomainEmail string
	WordDir     string
	StoragePath string

	sm *subscription.SubscribeManager
}

func (c *syncCmd) Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:      "sync",
		Aliases:   []string{"s"},
		Usage:     "Выполнение синхронизации Хранилища 1С с git репозиторием",
		ArgsUsage: "WORKDIR PATH",
		Flags:     append(c.Flags(), plugin.RegistryFlags("sync")...),
		Action:    c.run,
		Before: func(ctx *cli.Context) error {
			if !ctx.Args().Present() {
				err := cli.ShowSubcommandHelp(ctx)
				if err != nil {
					return err
				}
				return fmt.Errorf("WRONG USAGE: Requires a PATH argument")
			}

			switch ctx.Args().Len() {
			case 1:
				c.StoragePath = ctx.Args().Get(0)
			case 2:
				c.WordDir = ctx.Args().Get(0)
				c.StoragePath = ctx.Args().Get(1)
			default:
				return fmt.Errorf("WRONG USAGE: Requires a PATH argument")
			}
			var err error

			if c.sm, err = plugin.Subscribe("sync", ctx); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}

func (c *syncCmd) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Destination: &c.StorageUser,
			Name:        "storage-user",
			Aliases:     []string{"u"},
			Value:       "Администратор",
			DefaultText: "Администратор",
			Usage:       "пользователь хранилища 1C конфигурации",
			EnvVars:     strings.Fields("R2GITSYNC_STORAGE_USER GITSYNC_STORAGE_USER"),
		},
		&cli.StringFlag{
			Destination: &c.StoragePassword,
			Name:        "storage-pwd",
			Aliases:     []string{"p"},
			Usage:       "пароль пользователя хранилища 1C конфигурации",
			EnvVars:     strings.Fields("R2GITSYNC_STORAGE_PASSWORD GITSYNC_STORAGE_PWD GITSYNC_STORAGE_PASSWORD"),
		},
		&cli.BoolFlag{
			Destination: &c.DisableIncrement,
			Name:        "disable-increment",
			Aliases:     []string{"p"},
			Usage:       "отключает инкрементальную выгрузку",
			EnvVars:     strings.Fields("GITSYNC_DISABLE_INCREMENT"),
		},
		&cli.StringFlag{
			Destination: &c.Extension,
			Name:        "extension",
			Aliases:     []string{"e ext"},
			Usage:       "имя расширения для работы с хранилищем расширения",
			EnvVars:     strings.Fields("R2GITSYNC_EXTENSION GITSYNC_EXTENSION"),
		},
		&cli.IntFlag{
			Destination: &c.LimitVersions,
			Name:        "limit",
			Aliases:     []string{"l"},
			Usage:       "выгрузить не более <Количества> версий от текущей выгруженной",
			EnvVars:     strings.Fields("GITSYNC_LIMIT"),
		},
		&cli.IntFlag{
			Destination: &c.MinVersion,
			Name:        "min-version",
			Usage:       "<номер> минимальной версии для выгрузки",
			EnvVars:     strings.Fields("GITSYNC_MIN_VERSION"),
		},
		&cli.IntFlag{
			Destination: &c.MaxVersion,
			Name:        "max-version",
			Usage:       "<номер> максимальной версии для выгрузки",
			EnvVars:     strings.Fields("GITSYNC_MAX_VERSION"),
		},
	}
}

func (c *syncCmd) run(ctx *cli.Context) error {

	newOptions := *app.config.Options
	syncOptions = &newOptions
	syncOptions.Plugins = c.sm
	syncOptions.LicTryCount = 5

	err := manager2.Sync(repo, *syncOptions)

	return nil
}
