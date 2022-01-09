package cmd

//
// import (
// 	cli "github.com/jawher/mow.cli"
// 	"github.com/khorevaa/r2gitsync/internal/app"
// 	flags2 "github.com/khorevaa/r2gitsync/internal/app/flags"
// 	manager2 "github.com/khorevaa/r2gitsync/internal/manager"
// 	"github.com/khorevaa/r2gitsync/pkg/plugin"
// 	"github.com/khorevaa/r2gitsync/pkg/plugin/subscription"
// )
//
// func (app *app.Application) cmdSetVersion(cmd *cli.Cmd) {
//
// 	var (
// 		doCommit   bool
// 		setVersion int
// 		workdir    string
// 	)
//
// 	flags2.BoolOpt("c commit", false, "закоммитить изменения в git").
// 		Ptr(&doCommit).
// 		Apply(cmd, app.ctx)
// 	flags2.IntArg("VERSION", 0, "Номер версии для записи в файл.").
// 		Ptr(&setVersion).
// 		Apply(cmd, app.ctx)
// 	app.WorkdirArg.Ptr(&workdir).Apply(cmd, app.ctx)
//
// 	cmd.Spec = "[OPTIONS] VERSION [WORKDIR]"
//
// 	var sm *subscription.SubscribeManager
//
// 	cmd.Before = func() {
// 		var err error
// 		sm, err = plugin.Subscribe("set-version", app.ctx)
//
// 		if err != nil {
// 			app.failOnErr(err)
// 		}
//
// 	}
//
// 	cmd.Action = func() {
//
// 		err := manager2.DoTask(manager2.WriteVersionFile{Workdir: workdir,
// 			Filename: manager2.VERSION_FILE, Version: setVersion}, sm)
//
// 		app.failOnErr(err)
//
// 		if doCommit {
// 			err = manager2.CommitVersionFile(workdir)
// 			app.failOnErr(err)
// 		}
//
// 	}
// }
