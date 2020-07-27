package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
)

func cmdPlugins(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("list the contents of the safe here")
	}
}
