package plugin

import (
	"github.com/duke-git/lancet/slice"
	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/logos"
	"github.com/khorevaa/r2gitsync/pkg/plugin/subscription"
	"github.com/urfave/cli/v2"
)

var log = logos.New("github.com/khorevaa/r2gitsync/pkg/plugin")

type manager struct {
	plugins       map[string]Symbol
	pluginsConfig map[string]*ucfg.Config
}

type Manager interface {
	Flags(module string) []cli.Flag
	Subscriber(module string) (*subscription.SubscribeManager, error)
}

func NewPluginManager(cfg map[string]*ucfg.Config) Manager {
	m := &manager{
		plugins:       map[string]Symbol{},
		pluginsConfig: map[string]*ucfg.Config{},
	}
	for name, config := range cfg {
		if symbol, ok := plugins[name]; ok {
			m.plugins[name] = symbol
			m.pluginsConfig[name] = config
		}
	}

	return m
}

func (m *manager) Flags(module string) []cli.Flag {
	var flags []cli.Flag
	log.Info("Получаю дополните флаги",
		logos.String("module", module))

	for _, symbol := range m.plugins {
		if !slice.Contain(symbol.Modules, module) {
			continue
		}

		flags = append(flags, symbol.Flags...)
	}
	log.Info("Получены дополнительные флаги",
		logos.String("module", module),
		logos.Int("количество", len(flags)))

	return flags
}

func (m *manager) Subscriber(module string) (*subscription.SubscribeManager, error) {

	log.Info("Получаю подписчиков",
		logos.String("module", module),
		logos.String("method", "Subscriber"))

	sm := subscription.NewSubscribeManager()

	var pList []string

	for _, symbol := range m.plugins {
		if !slice.Contain(symbol.Modules, module) {
			continue
		}
		cfg := m.pluginsConfig[symbol.Name]
		pl, err := symbol.New(cfg)
		if err != nil {
			return nil, err
		}

		sm.Subscribe(pl.Subscribe())
		pList = append(pList, symbol.Name)
	}

	log.Info("Получено подписчиков",
		logos.String("module", module),
		logos.Any("плагины", pList),
		logos.Int("количество", sm.Count()),
		logos.String("method", "Subscriber"))

	return sm, nil
}
