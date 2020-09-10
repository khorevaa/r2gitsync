package plugin

import (
	"github.com/khorevaa/r2gitsync/cmd"
	"github.com/micro/cli/v2"
	"net/http"
)

// Plugin is the interface for plugins to micro. It differs from go-micro in that it's for
// the micro API, Web, Sidecar, CLI. It's a method of building middleware for the HTTP side.
type PluginSymbol interface {

	// Global Flags
	Flags() []cli.Flag
	// Sub-commands
	Commands() []*cli.Command

	// Name of the plugin
	String() string
	Desc() string
	Version() string
	Name() string
	New() Plugin
}

type Plugin interface {

	// Init called when command line args are parsed.
	// The initialised cli.Context is passed in.
	Init(sm SubscribeManager) error
	Handler() Handler
}

type SubscribeManager interface {
	Handle(endpoint interface{}, handler interface{})
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

// Handler is the plugin middleware handler which wraps an existing http.Handler passed in.
// Its the responsibility of the Handler to call the next http.Handler in the chain.
type Handler func(http.Handler) http.Handler

type plugin struct {
	opts    Options
	new     func() Plugin
	handler Handler
}

func (p *plugin) Flags() []cli.Flag {
	return cmd.StringOpt{}
}

func (p *plugin) Commands() []*cli.Command {
	return p.opts.Commands
}

func (p *plugin) Handler() Handler {
	return p.handler
}

func (p *plugin) New() Plugin {
	return p.new()
}

func (p *plugin) String() string {
	return p.opts.Name
}

func (p *plugin) Name() string {
	return p.opts.Name
}

func newPlugin(opts ...Option) PluginSymbol {
	options := Options{
		Name: "default",
		Init: func(ctx *cli.Context) error { return nil },
	}

	for _, o := range opts {
		o(&options)
	}

	handler := func(hdlr http.Handler) http.Handler {
		for _, h := range options.Handlers {
			hdlr = h(hdlr)
		}
		return hdlr
	}

	return &plugin{
		opts:    options,
		handler: handler,
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

// NewPlugin makes it easy to create a new plugin
func NewPlugin(opts ...Option) Plugin {
	return newPlugin(opts...)
}
