package plugin

import (
	"github.com/duke-git/lancet/slice"
	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/r2gitsync/pkg/plugin/subscription"
	"github.com/urfave/cli/v2"
)

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

	for _, symbol := range m.plugins {
		if !slice.Contain(symbol.Modules, module) {
			continue
		}

		flags = append(flags, symbol.Flags...)
	}

	return flags
}

func (m *manager) Subscriber(module string) (*subscription.SubscribeManager, error) {

	sm := subscription.NewSubscribeManager()

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

	}

	return sm, nil
}
