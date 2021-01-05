package cmd

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/manager"
	"github.com/khorevaa/r2gitsync/plugin"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
)

func (app *Application) cmdSetVersion(cmd *cli.Cmd) {

	var (
		doCommit   bool
		setVersion int
		workdir    string
	)

	flags.BoolOpt("c commit", false, "закоммитить изменения в git").
		Ptr(&doCommit).
		Apply(cmd, app.ctx)
	flags.IntArg("VERSION", 0, "Номер версии для записи в файл.").
		Ptr(&setVersion).
		Apply(cmd, app.ctx)
	WorkdirArg.Ptr(&workdir).Apply(cmd, app.ctx)

	cmd.Spec = "[OPTIONS] VERSION [WORKDIR]"

	var sm *subscription.SubscribeManager

	cmd.Before = func() {
		var err error
		sm, err = plugin.Subscribe("set-version", app.ctx)

		if err != nil {
			app.failOnErr(err)
		}

	}

	cmd.Action = func() {

		err := manager.DoTask(manager.WriteVersionFile{Workdir: workdir,
			Filename: manager.VERSION_FILE, Version: setVersion}, sm)

		app.failOnErr(err)

		if doCommit {
			err = manager.CommitVersionFile(workdir)
			app.failOnErr(err)
		}

	}
}
