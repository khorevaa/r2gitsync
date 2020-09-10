package plugin

import (
	cli2 "github.com/jawher/mow.cli"
	"github.com/micro/cli/v2"
	"net/http"
)

type Option func(o *plugin)

type InitFn func() Plugin

// WithFlag adds flags to a plugin
func WithFlag(flag ...cli2.VarParam) Option {
	return func(o *plugin) {
		o.Flags = append(o.Flags, flag...)
	}
}

// WithCommand adds commands to a plugin
func WithCommand(cmd ...*cli.Command) Option {
	return func(o *plugin) {
		o.Commands = append(o.Commands, cmd...)
	}
}

// WithInit sets the init function
func WithInit(fn InitFn) Option {
	return func(o *plugin) {
		o.init = fn
	}
}

type plugin struct {
	version  string
	desk     string
	name     string
	init     InitFn
	flags    []cli.Flag
	commands []*cli.Command
}

func (p *plugin) Flags() []cli.Flag {
	return p.opts.Flags
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
