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
			"выгрузить не более <Количества> версий от текущей выгруженной",
			0,
			"GITSYNC_LIMIT"),
		flags.IntOpt(
			"minversion",
			"<номер> минимальной версии для выгрузки",
			0,
			"GITSYNC_MIN_VERSION"),
		flags.IntOpt(
			"maxversion",
			"<номер> максимальной версии для выгрузки",
			0,
			"GITSYNC_MAX_VERSION"),
	))

type limitPlugin struct {
	limit      int
	minversion int
	maxversion int
}

func (t *limitPlugin) Init(sm plugin.SubscribeManager) error {

	sm.Handle(subscription.GetRepositoryHistory, subscription.OnEvent, "")

}

func (t *limitPlugin) InitContext(ctx context.Context) {

	t.limit = ctx.Int("limit")
	t.maxversion = ctx.Int("maxversion")
	t.minversion = ctx.Int("minversion")

}
