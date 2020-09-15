package plugin

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"sort"
	"sync"
)

const defaultModule = "r2gitsync"

type manager struct {
	sync.Mutex

	modules    map[string]*pluginsModule
	registered RegisteredPluginList
	plugins    map[string]Symbol
}

type pluginsModule struct {
	sync.Mutex
	manager *manager
	id      string
	plugins map[string]*modulePlugin
	sm      *subscription.SubscribeManager
}

type modulePlugin struct {
	id      string
	sym     Symbol
	enabled bool
}

func (m *pluginsModule) changeEnable(name string, enable bool) {

	m.Lock()
	defer m.Unlock()

	if pl, ok := m.plugins[name]; ok {
		pl.enabled = enable
	}

}

func (m *pluginsModule) Enable(name string) {

	m.changeEnable(name, true)

}

func (m *pluginsModule) Disable(name string) {

	m.changeEnable(name, false)

}

func (m *pluginsModule) Register(sym Symbol) error {

	m.Lock()
	defer m.Unlock()

	name := sym.String()

	if _, ok := m.plugins[name]; ok {
		return fmt.Errorf("Plugin with name %s already registered", name)
	}

	m.plugins[name] = &modulePlugin{
		id:      name,
		sym:     sym,
		enabled: true,
	}

	return nil
}

func (m *pluginsModule) Plugins() (pl []Symbol) {

	if len(m.plugins) > 0 {
		for _, p := range m.plugins {
			pl = append(pl, p.sym)
		}

	}

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Name() > pl[j].Name()
	})

	return

}

func (m *pluginsModule) PluginsList() (pl []struct {
	sym    Symbol
	enable bool
}) {

	if len(m.plugins) > 0 {

		for _, p := range m.plugins {
			pl = append(pl, struct {
				sym    Symbol
				enable bool
			}{sym: p.sym, enable: p.enabled})
		}

		sort.Slice(pl, func(i, j int) bool {
			return pl[i].sym.Name() > pl[j].sym.Name()
		})
	}

	return
}

func (m *pluginsModule) IsEnabled(name string) bool {

	if !m.IsRegistered(name) {
		return false
	}

	return m.plugins[name].enabled

}
func (m *pluginsModule) IsRegistered(name string) bool {

	m.Lock()
	defer m.Unlock()

	if _, ok := m.plugins[name]; ok {
		return true
	}

	return false

}

func (m *pluginsModule) Plugin(name string) (Symbol, bool) {

	moduleP, ok := m.plugins[name]

	if ok {
		return nil, false
	}

	return moduleP.sym, moduleP.enabled

}

func (m *pluginsModule) Subscribe(ctx context.Context) error {

	mErr := &multierror.Error{}

	for _, pl := range m.plugins {

		if !m.IsEnabled(pl.id) {
			continue
		}

		p := pl.sym.Init()

		err := m.sm.Subscribe(p, ctx)

		mErr = multierror.Append(mErr, err)

	}

	return mErr.ErrorOrNil()
}

var (
	// global plugin manager
	defaultManager = newManager()
)

func mewModule(id string) *pluginsModule {
	return &pluginsModule{
		id:      id,
		plugins: make(map[string]*modulePlugin),
		sm:      subscription.NewSubscribeManager(),
	}
}

func (m *manager) findOrCreateModule(id string) *pluginsModule {

	if mod, ok := m.modules[id]; ok {
		return mod
	}

	mod := mewModule(id)
	m.modules[id] = mod

	return mod
}

func newManager() *manager {
	return &manager{
		modules:    make(map[string]*pluginsModule),
		registered: make(map[string]map[string]bool),
	}
}

func (m *manager) Plugins() (pl map[string][]string) {

	m.Lock()
	defer m.Unlock()

	pl = make(map[string][]string)

	for key, mod := range m.registered {

		pl[key] = []string{}

		for p, _ := range mod {

			pl[key] = append(pl[key], p)

		}

	}

	return
}

func (m *manager) Plugin(name, module string) (Symbol, bool) {

	mod, ok := m.modules[module]

	if !ok {
		return nil, false
	}

	return mod.Plugin(name)

}

func (m *manager) EnabledPlugins(modName string) (p []Symbol) {

	m.Lock()
	defer m.Unlock()

	mod, ok := m.modules[modName]

	if !ok {
		return
	}

	pl := mod.PluginsList()

	for _, s := range pl {
		if s.enable {
			p = append(p, s.sym)
		}
	}

	return
}

func (m *manager) Register(sym Symbol) error {

	modules := sym.Modules()

	m.Lock()
	defer m.Unlock()

	name := sym.Name()

	for _, modId := range modules {

		if reg, ok := m.registered[modId]; ok && reg[name] {
			return fmt.Errorf("plugin with name %s already registered", name)
		}

		mod := m.findOrCreateModule(modId)
		err := mod.Register(sym)

		if err != nil {
			return err
		}

		if _, ok := m.registered[modId]; !ok {
			m.registered[modId] = map[string]bool{name: true}
		} else {
			m.registered[modId][name] = true
		}

	}

	return nil
}

func (m *manager) IsRegistered(sym Symbol, module string) bool {

	m.Lock()
	defer m.Unlock()

	if _, ok := m.registered[module]; !ok {
		return false
	}

	return m.registered[module][sym.Name()]

}

func (m *manager) Enable(name string) {

	m.Lock()
	defer m.Unlock()

	for _, mod := range m.modules {
		mod.Enable(name)
	}

}

func (m *manager) Disable(name string) {

	m.Lock()
	defer m.Unlock()

	for _, mod := range m.modules {
		mod.Disable(name)
	}

}

func (m *manager) RegisterFlags(module string, cmd command, ctx context.Context) {

	plugins := m.EnabledPlugins(module)

	for _, symbol := range plugins {

		registryFlags(symbol.Flags(), cmd, ctx)

	}

	return
}

func (m *manager) Subscribe(modName string, ctx context.Context) error {

	mod, ok := m.modules[modName]

	if !ok {
		return errors.New("no found pluginsModule")
	}

	err := mod.Subscribe(ctx)

	return err
}

func (m *manager) SubscribeManager(modName string) *subscription.SubscribeManager {
	mod, ok := m.modules[modName]

	if !ok {
		return nil
	}
	return mod.sm

}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func registryFlags(flag []flags.Flag, cmd command, ctx context.Context) {

	for _, f := range flag {

		f.Apply(cmd, ctx)

	}

}
