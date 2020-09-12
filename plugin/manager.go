package plugin

import (
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
	plugins    map[string]Symbol
	registered map[string]bool
	enabled    map[string]bool

	sm *subscription.SubscribeManager
}

var (
	// global plugin manager
	defaultManager = newManager()
)

func newManager() *manager {
	return &manager{
		plugins:    make(map[string]Symbol),
		registered: make(map[string]bool),
		enabled:    make(map[string]bool),
		sm:         subscription.NewSubscribeManager(),
	}
}

func (m *manager) Plugins() (pl []Symbol) {

	m.Lock()
	defer m.Unlock()

	if len(m.plugins) > 0 {
		for _, p := range m.plugins {
			pl = append(pl, p)
		}

	}

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Name() > pl[j].Name()
	})

	return
}

func (m *manager) EnabledPlugins() []Symbol {

	m.Lock()
	defer m.Unlock()

	if len(m.plugins) > 0 {

		var pl []Symbol
		for _, p := range m.plugins {

			if !m.IsEnabled(p.String()) {
				continue
			}

			pl = append(pl, p)
		}

		return pl
	}

	return []Symbol{}
}

func (m *manager) Register(sym Symbol) error {

	m.Lock()
	defer m.Unlock()

	name := sym.String()

	if m.registered[name] {
		return fmt.Errorf("Plugin with name %s already registered", name)
	}

	m.plugins[name] = sym
	m.registered[name] = true
	m.enabled[name] = true

	return nil
}

func (m *manager) IsRegistered(name Symbol) bool {

	m.Lock()
	defer m.Unlock()
	return m.registered[name.String()]

}

func (m *manager) Enable(name string) error {

	m.Lock()
	defer m.Unlock()

	if !m.registered[name] {
		return fmt.Errorf("Plugin with name %s is not registered", name)
	}

	m.enabled[name] = true

	return nil

}

func (m *manager) Disable(name string) error {

	m.Lock()
	defer m.Unlock()

	if !m.registered[name] {
		return fmt.Errorf("Plugin with name %s is not registered", name)
	}

	m.enabled[name] = false

	return nil

}

func (m *manager) IsEnabled(name string) bool {

	m.Lock()
	defer m.Unlock()

	return m.registered[name] && m.enabled[name]

}

func (m *manager) RegisterFlags(name string, cmd command, ctx context.Context) {

	for _, symbol := range m.plugins {

		if !m.IsEnabled(symbol.String()) {
			continue
		}

		if contains(symbol.Commands(), name) {

			registryFlags(symbol.Flags(), cmd, ctx)

		}

	}

	return
}

func (m *manager) Subscribe() error {

	mErr := &multierror.Error{}

	for _, sym := range m.plugins {

		if !m.IsEnabled(sym.String()) {
			continue
		}

		p := sym.Init()
		err := m.sm.Subscribe(p)

		mErr = multierror.Append(mErr, err)

	}

	return mErr.ErrorOrNil()
}

func (m *manager) SendContext(ctx context.Context) {

	m.sm.SendContext(ctx)

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
