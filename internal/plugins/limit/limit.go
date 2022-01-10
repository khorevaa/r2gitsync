package limit

import (
	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/r2gitsync/internal/manager/types"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
	"github.com/urfave/cli/v2"
	"math"
)

var (
	version = "dev"
	commit  = ""
)

type Config struct {
	Limit      *uint
	MinVersion *uint
	MaxVersion *uint
}

var defaultConfig = Config{}
var flagsConfig = Config{}

func New(cfg *ucfg.Config) (plugin.Plugin, error) {

	config := defaultConfig

	err := cfg.Unpack(&config)
	if err != nil {
		return nil, err
	}

	// TODO Сделать слияние конфигов

	return &Plugin{
		limit:      config.Limit,
		minVersion: config.MinVersion,
		maxVersion: config.MaxVersion,
	}, nil
}

//goland:noinspection ALL
var Symbol = plugin.Symbol{
	"limit",
	plugin.BuildVersion(version, commit),
	"Плагин добавляет возможность органичений при выгрузке конфигурации",
	New,
	[]string{"sync"},
	[]cli.Flag{
		&cli.UintFlag{
			Name:        "limit",
			Usage:       "выгрузить не более <Количества> версий от текущей выгруженной",
			EnvVars:     []string{"GITSYNC_LIMIT"},
			Destination: flagsConfig.Limit,
		},
		&cli.UintFlag{
			Name:        "min-version",
			Usage:       "в<номер> минимальной версии для выгрузки",
			EnvVars:     []string{"GITSYNC_MIN_VERSION"},
			Destination: flagsConfig.MinVersion,
		},
		&cli.UintFlag{
			Name:        "maxVersion",
			Usage:       "<номер> максимальной версии для выгрузки",
			EnvVars:     []string{"GITSYNC_MAX_VERSION"},
			Destination: flagsConfig.MaxVersion,
		},
	}}

type Plugin struct {
	limit      *uint
	minVersion *uint
	maxVersion *uint
}

func (t *Plugin) Subscribe() Subscriber {

	return plugin.Subscription(
		ConfigureRepositoryVersionsSubscriber{
			On: t.ConfigureRepositoryVersions,
		})

}

func (t *Plugin) ConfigureRepositoryVersions(end V8Endpoint, versions *types.RepositoryVersionsList, Current, Next, Max *int) error {

	ls := *versions

	if len(ls) == 0 {
		return nil
	}

	if *t.minVersion > 0 {
		*Next = int(*t.minVersion)
	}

	if *t.limit > 0 {

		startVersion := math.Max(float64(*Next), float64(*Current))

		limitVersion := startVersion + float64(*t.limit) - 1 // -1, для того чтобы учесть следующую версию она всегда на 1 больше текущей
		*Max = int(limitVersion)

	}

	if *t.maxVersion > 0 {
		if *t.limit > 0 {
			*Max = int(math.Min(float64(*Max), float64(*t.maxVersion)))
		} else {
			*Max = int(*t.maxVersion)
		}

	}

	return nil
}
