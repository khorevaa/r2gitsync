package cmd

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/plugin"
	"os"
	"strings"
	"text/tabwriter"
)
import "github.com/mgutz/ansi"

func (app *Application) cmdPluginsList(cmd *cli.Cmd) {

	var showAll bool

	flags.BoolOpt("a all", false, "показать все плагины").
		Ptr(&showAll).
		Apply(cmd, app.ctx)

	cmd.Action = func() {
		list := plugin.Plugins()

		log.Debugw("plugins list", "list", list)

		if len(list) > 0 {

			//w := os.Stdout
			stdOut := os.Stderr
			//fmt.Fprintln(stdOut, "\t\nPlugins list:\t\n")
			w := tabwriter.NewWriter(stdOut, 10, 1, 3, ' ', 0)
			defer w.Flush()
			fmt.Fprintf(stdOut, "Список плагинов:\t\n")
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t  %s\t\n", "Состояние", "Версия", "Имя", "Модули", "Описание")
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t  %s\t\n", "---------", "------", "----", "------", "--------")
			for _, arg := range list {

				if arg.Enable || showAll {
					var (
						enable  = ansi.Color("выкл\t", "red")
						name    = arg.Name()
						desc    = arg.Desc()
						version = arg.ShortVersion()
						modules = strings.Join(arg.Modules(), ",")
					)

					if arg.Enable {
						enable = ansi.Color("вкл\t", "green")
					}

					fmt.Fprintf(w, "%s\t %s\t %s\t %s\t%s\n", enable, version, name, modules, desc)
				}
			}

		}

	}
}

func (app *Application) cmdPluginsEnable(cmd *cli.Cmd) {

	var (
		all     bool
		global  bool
		plNames []string
	)
	flags.BoolOpt("a all", false, "включить все плагины").
		Ptr(&all).
		Apply(cmd, app.ctx)

	flags.BoolOpt("g global", false, "использовать глобальное хранилище плагинов").
		Ptr(&global).
		Apply(cmd, app.ctx)

	flags.StringsArg("PLUGINS", plNames, "имена плагинов").
		Ptr(&plNames).
		Apply(cmd, app.ctx)

	cmd.Spec = "[-g | --global] PLUGINS... | --all"

	cmd.Action = func() {

		fmt.Printf("Включаю плагины")
		fmt.Printf("Список плагинов: %s \n", plNames)

	}
}

func (app *Application) cmdPluginsDisable(cmd *cli.Cmd) {

	var (
		all     bool
		global  bool
		plNames []string
		//local  bool
	)

	flags.BoolOpt("a all", false, "выключить все плагины").
		Ptr(&all).
		Apply(cmd, app.ctx)

	flags.BoolOpt("g global", false, "использовать глобальное хранилище плагинов").
		Ptr(&global).
		Apply(cmd, app.ctx)

	flags.StringsArg("PLUGINS", plNames, "имена плагинов").
		Ptr(&plNames).
		Apply(cmd, app.ctx)

	cmd.Spec = "[-g | --global] PLUGINS... | --all"

	//flags.BoolOpt("l local", true, "использовать локальное хранилище плагинов").
	//	Ptr(&local).
	//	Apply(cmd, app.ctx)

	cmd.Action = func() {

		fmt.Printf("Отключаю плагины")
		fmt.Printf("Список плагинов: %s \n", plNames)

	}
}

func (app *Application) cmdPluginsClear(cmd *cli.Cmd) {

	var (
		all     bool
		global  bool
		plNames []string
		//local  bool
	)

	flags.BoolOpt("a all", false, "удалить все плагины").
		Ptr(&all).
		Apply(cmd, app.ctx)

	flags.BoolOpt("g global", false, "использовать глобальное хранилище плагинов").
		Ptr(&global).
		Apply(cmd, app.ctx)

	flags.StringsArg("PLUGINS", plNames, "имена плагинов").
		Ptr(&plNames).
		Apply(cmd, app.ctx)

	cmd.Spec = "[-g | --global] PLUGINS... | --all"

	cmd.Action = func() {

		fmt.Printf("Удаляю плагины")
		fmt.Printf("Список плагинов: %s \n", plNames)

	}
}

func (app *Application) cmdPluginsInstall(cmd *cli.Cmd) {

	var (
		files  []string
		global bool
	)

	flags.BoolOpt("g global", false, "использовать глобальное хранилище плагинов").
		Ptr(&global).
		Apply(cmd, app.ctx)

	flags.StringsArg("FILE", files, "Путь к файлу плагина или zip ахриву с плагинами").
		Ptr(&files).
		Apply(cmd, app.ctx)

	cmd.Spec = "[OPTIONS] FILE..."

	cmd.Action = func() {

		fmt.Printf("Устанавливаю новые плагины")
		//fmt.Printf( "Список плагинов: %s \n", files )

	}
}
