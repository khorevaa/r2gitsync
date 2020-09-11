package cmd

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
)

func (app *Application) cmdPlugins(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("list the contents of the safe here")
	}
}
