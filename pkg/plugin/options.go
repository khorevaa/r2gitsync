package plugin

import (
	"github.com/khorevaa/r2gitsync/internal/app/flags"
	metadata2 "github.com/khorevaa/r2gitsync/pkg/plugin/metadata"
	"strings"
)

type Option func(o *plugin)

type InitFn func() Plugin

// WithFlag adds flags to a plugin
func WithFlag(flag ...flags.Flag) Option {
	return func(o *plugin) {
		o.flags = append(o.flags, flag...)
	}
}

// WithModule adds modules to a plugin
func WithModule(module ...string) Option {
	return func(o *plugin) {
		o.modules = append(o.modules, module...)
	}
}

// WithInit sets the Init function
func WithInit(fn InitFn) Option {
	return func(o *plugin) {
		o.init = fn
	}
}

type plugin struct {
	version string
	desk    string
	name  string
	init  InitFn
	flags []flags.Flag
	modules []string
}

func (p *plugin) Flags() []flags.Flag {
	return p.flags
}

func (p *plugin) Modules() []string {
	return p.modules
}

func (p *plugin) Init() metadata2.Plugin {
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

func (p *plugin) ShortVersion() string {

	v := strings.Split(p.version, "+")
	return v[0]
}

func newPlugin(name, version, desc string, init InitFn, opts ...Option) Symbol {

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
func NewPlugin(name, version, desc string, init InitFn, opts ...Option) Symbol {
	return newPlugin(name, version, desc, init, opts...)
}

type pluginsPackage struct {
	name    string
	version string
	plugins []metadata2.PluginSymbol
}

func (p *pluginsPackage) Name() string {
	return p.name
}

func (p *pluginsPackage) Version() string {
	return p.version
}

func (p *pluginsPackage) Plugins() []metadata2.PluginSymbol {
	return p.plugins
}

type PkgOption func(pkg *pluginsPackage)

func WithPlugin(sym Symbol) PkgOption {

	return func(pkg *pluginsPackage) {
		if sym != nil {
			pkg.plugins = append(pkg.plugins, sym)
		}
	}

}

func WithNewPlugin(name, version, desc string, init InitFn, opts ...Option) PkgOption {
	sym := newPlugin(name, version, desc, init, opts...)

	return func(pkg *pluginsPackage) {
		pkg.plugins = append(pkg.plugins, sym)
	}

}

// NewPlugin makes it easy to create a new plugin
func NewPkg(name, version string, opts ...PkgOption) metadata2.PkgSymbol {
	return newPluginsPackage(name, version, opts...)
}

func newPluginsPackage(name, version string, opts ...PkgOption) metadata2.PkgSymbol {

	p := pluginsPackage{
		name:    name,
		version: version,
		plugins: []metadata2.PluginSymbol{},
	}

	for _, o := range opts {
		o(&p)
	}

	return &p
}
