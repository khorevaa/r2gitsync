package main

import (
	"github.com/hashicorp/go-multierror"
	"github.com/khorevaa/r2gitsync/plugin/Subscription"
)

type ManagerConfig struct {
	enable  []string
	disable []string
	flags   map[string]interface{}
}

type PluginsManager struct {
	*Subscription.SubscribeManager
	plugins []*PluginSymbol
	config  ManagerConfig
}

func NewPluginsManager(cfg ManagerConfig) *PluginsManager {

	return &PluginsManager{
		SubscribeManager: &Subscription.SubscribeManager{},
		config:           cfg,
	}

}

func (m *PluginsManager) LoadPlugins(pl *PluginsLoader) (err error) {

	for _, pName := range m.config.enable {

		pSymbol := pl.NewPluginSymbol(pName)
		if pSymbol != nil {
			m.plugins = append(m.plugins, pSymbol)
		}
	}

	// TODO Сделать более информативное сообщение об ошибках
	// TODO Добавить имя плагина и имя функции
	for _, plugin := range m.plugins {

		subs := plugin.plugin.RegistryHandlers()
		for tupic, fn := range subs {
			errSub := m.sm.Subscribe(tupic, fn)

			if errSub != nil {
				err = multierror.Append(err, errSub)
			}
		}
	}

	return
}

func (m *PluginsManager) ConfigurePlugins(map[string]interface{}) (err error) {

	// TODO Сделать настройку плагина по прочитанным данным
	// TODO Добавить имя плагина и имя функции

	return
}

func (m *PluginsManager) RegistryOptions(name string, cmd command) {

	for _, plugin := range m.plugins {

		plugin.plugin.RegistryOptions(name, cmd)

	}

	return
}
