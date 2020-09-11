package limit

import (
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/plugin"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
)

var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

var NewPlugin = plugin.NewPlugin(
	"limit",
	plugin.BuildVersion(version, commit, date, builtBy),
	"Плагин добавляет возможность инкрементальной выгрузки в конфигурации",
	func() plugin.Plugin {
		return new(limitPlugin)
	},
	plugin.WithCommand("sync"),
	plugin.WithFlag(
		flags.IntOpt(
			"l limit",
			0,
			"выгрузить не более <Количества> версий от текущей выгруженной").
			Env("GITSYNC_LIMIT"),
		flags.IntOpt(
			"minversion",
			0,
			"<номер> минимальной версии для выгрузки").
			Env("GITSYNC_MIN_VERSION"),
		flags.IntOpt(
			"maxversion",
			0,
			"<номер> максимальной версии для выгрузки").
			Env("GITSYNC_MAX_VERSION"),
	))

type limitPlugin struct {
	limit      int
	minversion int
	maxversion int
}

func (t *limitPlugin) Init(sm plugin.SubscribeManager) error {

	sm.Handle(subscription.GetRepositoryHistory, subscription.OnEvent, "")

	return nil

}

func (t *limitPlugin) InitContext(ctx context.Context) {

	t.limit = ctx.Int("limit")
	t.maxversion = ctx.Int("maxversion")
	t.minversion = ctx.Int("minversion")

}
