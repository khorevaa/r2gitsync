package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/v8platform/designer/repository"
	"io/ioutil"
	"os"
	"path"
	"plugin"
	"regexp"
)

const pluginSymbolName = "Plugins"

type command interface {
	String(p cli.StringParam) *string
	StringPtr(into *string, p cli.StringParam)
	Bool(p cli.BoolParam) *bool
	BoolPtr(into *bool, p cli.BoolParam)
	Int(p cli.IntParam) *int
	Float64(p cli.Float64Param) *float64
	Strings(p cli.StringsParam) *[]string
}

type PluginsInterface interface {
	Init() error
	Version() string
	Desc() string
	Help() string
	Symbol() interface{}
	RegistryOptions(name string, cmd command)
	RegistryHandlers() map[string]interface{}
}

// Module a plugin that can be initialized
type Module interface {
	Init() error
}

// Commands a plugin that contains one or more command
type ModulePlugins interface {
	Module
	Registry() map[string]newPluginFunc
}

type PluginSymbol struct {
	Name   string
	plugin PluginsInterface
}

func (pl *PluginsLoader) NewPluginSymbol(name string) *PluginSymbol {

	if newFn, ok := pl.plugins[name]; ok {
		p := newFn()
		return &PluginSymbol{
			Name:   name,
			plugin: p,
		}
	}

	return nil
}

type BeforeStartSyncProcess interface {
	BeforeStartSyncProcess(r repository.Repository, dir string)
}

func (p *PluginSymbol) BeforeStartSyncProcess(r repository.Repository, dir string) {

	pluginSys, ok := p.plugin.Symbol().(BeforeStartSyncProcess)
	if ok {
		pluginSys.BeforeStartSyncProcess(r, dir)
	}

}

type newPluginFunc func() PluginsInterface

type PluginsLoader struct {
	plugins    map[string]newPluginFunc
	pluginsDir string

	pluginsLoaded bool
}

func NewPluginsLoader(dir string) *PluginsLoader {

	pl := &PluginsLoader{
		pluginsDir: dir,
	}

	return pl
}

func (pl *PluginsLoader) LoadPlugins(force bool) error {

	if pl.pluginsLoaded && !force {
		return nil
	}

	if _, err := os.Stat(pl.pluginsDir); err != nil {
		return err
	}

	plugins, err := listFiles(pl.pluginsDir, `*.so`)
	if err != nil {
		return err
	}

	for _, cmdPlugin := range plugins {
		pluginFile, err := plugin.Open(path.Join(pl.pluginsDir, cmdPlugin.Name()))
		if err != nil {
			fmt.Printf("failed to open pluginFile %s: %v\n", cmdPlugin.Name(), err)
			continue
		}
		pluginSymbol, err := pluginFile.Lookup(pluginSymbolName)
		if err != nil {
			fmt.Printf("pluginFile %s does not export symbol \"%s\"\n",
				cmdPlugin.Name(), pluginSymbolName)
			continue
		}
		plugin, ok := pluginSymbol.(ModulePlugins)
		if !ok {
			fmt.Printf("plugin %s (from %s) does not implement Pluing module interface\n",
				pluginSymbolName, cmdPlugin.Name())
			continue
		}
		if err := plugin.Init(); err != nil {
			fmt.Printf("%s initialization failed: %v\n", cmdPlugin.Name(), err)
			continue
		}

		for name, newPluginFn := range plugin.Registry() {
			pl.plugins[name] = newPluginFn
		}

	}

	pl.pluginsLoaded = true

	return nil
}

func listFiles(dir, pattern string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filteredFiles := []os.FileInfo{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		matched, err := regexp.MatchString(pattern, file.Name())
		if err != nil {
			return nil, err
		}
		if matched {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles, nil
}
