package cmd

import (
	"fmt"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
	"github.com/khorevaa/r2gitsync/pkg/plugin/subscription"
	"github.com/urfave/cli/v2"
	"log"
	"strings"
)

type syncCmd struct {
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

func (c *syncCmd) Cmd(manager plugin.Manager) *cli.Command {
	cmd := &cli.Command{
		Name:      "sync",
		Aliases:   []string{"s"},
		Usage:     "Выполнение синхронизации Хранилища 1С с git репозиторием",
		ArgsUsage: "WORKDIR PATH",
		Flags:     append(c.Flags(), manager.Flags("sync")...),
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
			if c.sm, err = manager.Subscriber("sync"); err != nil {
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
			Usage:       "отключает инкрементальную выгрузку",
			EnvVars:     strings.Fields("GITSYNC_DISABLE_INCREMENT"),
		},
		&cli.StringFlag{
			Destination: &c.Extension,
			Name:        "extension",
			Aliases:     strings.Fields("e ext"),
			Usage:       "имя расширения для работы с хранилищем расширения",
			EnvVars:     strings.Fields("R2GITSYNC_EXTENSION GITSYNC_EXTENSION"),
		},
	}
}

func (c *syncCmd) run(ctx *cli.Context) error {

	log.Println(c.sm.ConfigureRepositoryVersions)
	return nil
}
