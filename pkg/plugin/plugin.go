package plugin

import (
	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/r2gitsync/pkg/plugin/metadata"
	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
	"github.com/urfave/cli/v2"
	"sync"
)

// plugins список загруженных плагинов
var plugins = map[string]Symbol{}

// muPlugins для избежания дедлоков
var muPlugins = &sync.Mutex{}

type Symbol struct {
	Name    string
	Version string
	Desc    string
	New     func(cfg *ucfg.Config) (Plugin, error)
	Modules []string
	Flags   []cli.Flag
}

type Plugin interface {
	metadata.Plugin
}

// Register registers a global plugins
func Register(symbols ...Symbol) {
	muPlugins.Lock()
	defer muPlugins.Unlock()
	for _, pl := range symbols {
		plugins[pl.Name] = pl
	}
}

func Subscription(handlers ...interface{}) Subscriber {

	return subscriber{
		handlers: handlers,
	}
}

func LoadPlugins(dir string) error {

	loader := NewPluginsLoader(dir)

	err := loader.LoadPlugins(false)

	if err != nil {
		return err
	}

	Register(loader.Plugins()...)

	return nil
}
