package limit

import (
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/plugin"
	. "github.com/khorevaa/r2gitsync/plugin/types"
)

var (
	version = "dev"
	commit  = ""
)

var NewPlugin = plugin.NewPlugin(
	"limit",
	plugin.BuildVersion(version, commit),
	"Плагин добавляет возможность органичений при выгрузке конфигурации",
	func() plugin.Plugin {
		return new(LimitPlugin)
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

type LimitPlugin struct {
	limit      int
	minversion int
	maxversion int
}

func (t *LimitPlugin) Subscriber() Subscriber {

	return plugin.Subscription(
		UpdateCfgSubscriber{
			Before: t.beforeUpdateCfg,
		})
	//return subscription.Subscriber{
	//	UpdateCfg: subscription.UpdateCfgSubscriber{
	//		Before: t.beforeUpdateCfg,
	//	},
	//}

}

func (t *LimitPlugin) beforeUpdateCfg(v8end V8Endpoint, workdir string, version int64) error {

	return nil

}

func (t *LimitPlugin) InitContext(ctx context.Context) {

	t.limit = ctx.Int("limit")
	t.maxversion = ctx.Int("maxversion")
	t.minversion = ctx.Int("minversion")

}
