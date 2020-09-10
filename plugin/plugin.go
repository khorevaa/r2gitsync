package plugin

import (
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
)

// Plugin is the interface for plugins to micro. It differs from go-micro in that it's for
// the micro API, Web, Sidecar, CLI. It's a method of building middleware for the HTTP side.
type PluginSymbol interface {

	// Global Flags
	Flags() []flags.Flag
	// Sub-commands
	Commands() []string

	// Name of the plugin
	String() string
	Desc() string
	Version() string
	Name() string
	Init() Plugin
}

type Plugin interface {
	Init(sm SubscribeManager) error
	InitContext(tx context.Context)
}

type SubscribeManager interface {
	Handle(endpoint subscription.EndPointType, event subscription.EventType, handler interface{})
}

// Manager is the plugin manager which stores plugins and allows them to be retrieved.
// This is used by all the components of micro.
type Manager interface {
	Plugins(...PluginOption) []Plugin
	Register(Plugin, ...PluginOption) error
}

type PluginOptions struct {
	Module string
}

type PluginOption func(o *PluginOptions)

// Module will scope the plugin to a specific module, e.g. the "api"
func Module(m string) PluginOption {
	return func(o *PluginOptions) {
		o.Module = m
	}
}

// Plugins lists the global plugins
func Plugins(opts ...PluginOption) []Plugin {
	return defaultManager.Plugins(opts...)
}

// Register registers a global plugins
func Register(plugin Plugin, opts ...PluginOption) error {
	return defaultManager.Register(plugin, opts...)
}

// IsRegistered check plugin whether registered global.
// Notice plugin is not check whether is nil
func IsRegistered(plugin Plugin, opts ...PluginOption) bool {
	return defaultManager.isRegistered(plugin, opts...)
}

// NewManager creates a new plugin manager
func NewManager() Manager {
	return newManager()
}
