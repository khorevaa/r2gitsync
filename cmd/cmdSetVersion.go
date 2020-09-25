package cmd

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/manager"
)

// Sample use: vault list OR vault config list
func (app *Application) cmdSetVersion(cmd *cli.Cmd) {

	var (
		doCommit            bool
		setVersion, workdir string
	)

	flags.BoolOpt("c commit", false, "закоммитить изменения в git").
		Ptr(&doCommit).
		Apply(cmd, app.ctx)
	flags.StringArg("VERSION", "", "Номер версии для записи в файл.").
		Ptr(&setVersion).
		Apply(cmd, app.ctx)
	WorkdirArg.Ptr(&workdir).Apply(cmd, app.ctx)

	cmd.Spec = "[OPTIONS] VERSION [WORKDIR]"

	cmd.Action = func() {

		err := manager.WriteVersionFile(workdir, setVersion)

		app.failOnErr(err)

		if doCommit {
			err = manager.CommitVersionFile(workdir)
			app.failOnErr(err)
		}

	}
}
