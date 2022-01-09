package cmd

import (
	"github.com/urfave/cli/v2"
)

var Commands = []Command{

	&syncCmd{},
	// &listCmd{},

}

type Command interface {
	Cmd() *cli.Command
}
