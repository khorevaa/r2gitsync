package plugin

import (
	"github.com/hashicorp/go-multierror"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"sync"
)

type manager struct {
	sync.Mutex

	modules    map[string]pluginsModule
	registered RegisteredPluginList // тут х

}

type pluginsModule struct {
	sync.Mutex
	manager *manager
	id      string
	sm      *subscription.SubscribeManager
}

func (m *manager) Subscribe(module string, ctx context.Context) (*subscription.SubscribeManager, error) {

	mErr := &multierror.Error{}

	sm := subscription.NewSubscribeManager()

	for _, pl := range m.registered.ByModule(module) {

		if !pl.Enable {
			continue
		}

		p := pl.Init()

		err := sm.Subscribe(p, ctx)

		mErr = multierror.Append(mErr, err)

	}

	if err := mErr.ErrorOrNil(); err != nil {
		return nil, err
	}

	return sm, nil
}

var (
	// global plugin manager
	defaultManager = newManager()
)

func newManager() *manager {
	return &manager{
		registered: make(RegisteredPluginList),
	}
}

func (m *manager) Plugins() []RegisteredPlugin {

	m.Lock()
	defer m.Unlock()

	return m.registered.Items()
}

func (m *manager) Register(sym Symbol) error {

	m.Lock()
	defer m.Unlock()

	//if _, ok := m.registered[sym.Name()]; ok {
	//	return fmt.Errorf("plugin with name %s already registered", sym.Name())
	//}

	m.registered[sym.Name()] = RegisteredPlugin{
		sym,
		true,
	}

	return nil
}

func (m *manager) IsRegistered(name string) bool {

	m.Lock()
	defer m.Unlock()

	if _, ok := m.registered[name]; ok {
		return true
	}

	return false

}

func (m *manager) Enable(name string) {

	m.Lock()
	defer m.Unlock()

	m.registered.Enable(name)

}

func (m *manager) Disable(name string) {

	m.Lock()
	defer m.Unlock()

	m.registered.Disable(name)

}

func (m *manager) RegisterFlags(module string, cmd command, ctx context.Context) {

	plugins := m.registered.ByModule(module)

	for _, pl := range plugins {

		registryFlags(pl.Flags(), cmd, ctx)

	}

	return
}

func registryFlags(flag []flags.Flag, cmd command, ctx context.Context) {

	for _, f := range flag {

		f.Apply(cmd, ctx)

	}

}
