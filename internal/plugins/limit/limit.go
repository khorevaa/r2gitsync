package limit

//
//import (
//	"github.com/khorevaa/r2gitsync/cmd/flags"
//	"github.com/khorevaa/r2gitsync/context"
//	"github.com/khorevaa/r2gitsync/manager/types"
//	"github.com/khorevaa/r2gitsync/plugin"
//	. "github.com/khorevaa/r2gitsync/plugin/types"
//	"math"
//)
//
//var (
//	version = "dev"
//	commit  = ""
//)
//
//var NewPlugin = plugin.NewPlugin(
//	"limit",
//	plugin.BuildVersion(version, commit),
//	"Плагин добавляет возможность органичений при выгрузке конфигурации",
//	func() plugin.Plugin {
//		return new(LimitPlugin)
//	},
//	plugin.WithModule("sync"),
//	plugin.WithFlag(
//		flags.IntOpt(
//			"l limit",
//			0,
//			"выгрузить не более <Количества> версий от текущей выгруженной").
//			Env("GITSYNC_LIMIT"),
//		flags.IntOpt(
//			"minversion",
//			0,
//			"<номер> минимальной версии для выгрузки").
//			Env("GITSYNC_MIN_VERSION"),
//		flags.IntOpt(
//			"maxversion",
//			0,
//			"<номер> максимальной версии для выгрузки").
//			Env("GITSYNC_MAX_VERSION"),
//	))
//
//type LimitPlugin struct {
//	plugin.BasePlugin
//	limit      int
//	minversion int
//	maxversion int
//}
//
//func (t *LimitPlugin) Subscribe(ctx context.Context) Subscriber {
//
//	t.Context = ctx
//
//	t.limit = ctx.Int("limit")
//	t.maxversion = ctx.Int("maxversion")
//	t.minversion = ctx.Int("minversion")
//
//	return plugin.Subscription(
//		ConfigureRepositoryVersionsSubscriber{
//			On: t.ConfigureRepositoryVersions,
//		})
//
//}
//
//func (t *LimitPlugin) ConfigureRepositoryVersions(end V8Endpoint, versions *[]types.RepositoryVersion, Current *int64, Next *int64, Max *int64) error {
//
//	ls := *versions
//
//	if len(ls) == 0 {
//		return nil
//	}
//
//	if t.minversion > 0 {
//		*Next = int64(t.minversion)
//	}
//
//	if t.limit > 0 {
//
//		startVersion := math.Max(float64(*Next), float64(*Current))
//
//		limitVersion := startVersion + float64(t.limit) - 1 // -1 для того чтобы учеть следущую версию она всегда на 1 больше текущей
//		*Max = int64(limitVersion)
//
//	}
//
//	if t.maxversion > 0 {
//		if t.limit > 0 {
//			*Max = int64(math.Min(float64(*Max), float64(t.maxversion)))
//		} else {
//			*Max = int64(t.maxversion)
//		}
//
//	}
//
//	return nil
//}
