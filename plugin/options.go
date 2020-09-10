package plugin

import (
	"github.com/khorevaa/r2gitsync/cmd/flags"
)

type Option func(o *plugin)

type InitFn func() Plugin

// WithFlag adds flags to a plugin
func WithFlag(flag ...flags.Flag) Option {
	return func(o *plugin) {
		o.flags = append(o.flags, flag...)
	}
}

// WithCommand adds commands to a plugin
func WithCommand(cmd ...string) Option {
	return func(o *plugin) {
		o.commands = append(o.commands, cmd...)
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
	flags    []flags.Flag
	commands []string
}

func (p *plugin) Flags() []flags.Flag {
	return p.flags
}

func (p *plugin) Commands() []string {
	return p.commands
}

func (p *plugin) Init() Plugin {
	return p.init()
}

func (p *plugin) String() string {
	return p.Name()
}

func (p *plugin) Name() string {
	return p.name
}

func (p *plugin) Desc() string {
	return p.desk
}

func (p *plugin) Version() string {
	return p.version
}

func newPlugin(name, version, desc string, init InitFn, opts ...Option) PluginSymbol {

	p := plugin{
		name:    name,
		version: version,
		desk:    desc,
		init:    init,
	}

	for _, o := range opts {
		o(&p)
	}

	return &p
}

// NewPlugin makes it easy to create a new plugin
func NewPlugin(name, version, desc string, init InitFn, opts ...Option) PluginSymbol {
	return newPlugin(name, version, desc, init, opts...)
}
