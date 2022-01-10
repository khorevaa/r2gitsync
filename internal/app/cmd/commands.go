package cmd

import (
	"github.com/khorevaa/r2gitsync/pkg/plugin"
	"github.com/urfave/cli/v2"
)

var Commands = []Command{

	&syncCmd{},
	// &listCmd{},

}

type Command interface {
	Cmd(manager plugin.Manager) *cli.Command
}
