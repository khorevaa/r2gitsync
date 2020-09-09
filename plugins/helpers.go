package main

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/plugins/limit"
)

type command interface {
	String(p cli.StringParam) *string
	StringPtr(into *string, p cli.StringParam)
	Bool(p cli.BoolParam) *bool
	BoolPtr(into *bool, p cli.BoolParam)
	Int(p cli.IntParam) *int
	IntPtr(into *int, p cli.IntParam)
	Float64(p cli.Float64Param) *float64
	Strings(p cli.StringsParam) *[]string
}

type PluginsInterface interface {
	Name() string
	Symbol() interface{}
	Version() string
	Desc() string
	RegistryOptions(name string, cmd command) error
	RegistryHandlers() map[string]interface{}
}

// command module
type pluginsModule struct{}
type newPluginFunc func() PluginsInterface

func (t *pluginsModule) Init() error {
	return nil
}

func (t *pluginsModule) limitPlugin() PluginsInterface {
	return &limit.limitPlugin{}
}
func (t *pluginsModule) incrementPlugin() PluginsInterface {
	return &IncrementPlugin{}
}

func (t *pluginsModule) Registry() map[string]newPluginFunc {
	return map[string]newPluginFunc{
		"increment": t.limitPlugin,
		"limit":     t.incrementPlugin,
	}
}

var Plugins
