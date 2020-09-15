package cmd

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/plugin"
	"os"
	"strings"
	"text/tabwriter"
)
import "github.com/mgutz/ansi"

func (app *Application) cmdPlugins(cmd *cli.Cmd) {
}

func (app *Application) cmdPluginsList(cmd *cli.Cmd) {

	var showAll bool

	flags.BoolOpt("a all", false, "показать все плагины").
		Ptr(&showAll).
		Apply(cmd, app.ctx)

	cmd.Action = func() {
		list := plugin.Plugins()

		if len(list) > 0 {

			//w := os.Stdout
			stdOut := os.Stderr
			//fmt.Fprintln(stdOut, "\t\nPlugins list:\t\n")
			w := tabwriter.NewWriter(stdOut, 10, 1, 3, ' ', 0)
			defer w.Flush()
			fmt.Fprintf(stdOut, "Список плагинов:\t\n")
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", "Состояние", "Версия", "Имя", "Команды", "Описание")
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", "---------", "------", "----", "-------", "--------")
			for _, arg := range list {

				if arg.Enable || showAll {
					var (
						enable  = ansi.Color("выкл\t", "red")
						name    = arg.Name
						desc    = arg.Desc
						version = arg.ShortVersion
						modules = strings.Join(arg.Modules, ",")
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
