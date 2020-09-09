package main

import (
	"context"
	"fmt"
	"github.com/v8platform/designer/repository"
	"io"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

var _ PluginsInterface = (*IncrementPlugin)(nil)

// command module
type IncrementPlugin struct {
}

func (t *IncrementPlugin) Init() error {
	fmt.Fprintln(out, "test module loaded OK")
	return nil
}

func (t *IncrementPlugin) RegistryOptions(name string, cmd command) error {

	return nil

}

func (t *IncrementPlugin) RegistryHandlers() map[string]interface{} {

	return map[string]interface{}{
		"BeforeUpdateCfgHandler": t.BeforeUpdateCfgHandler,
	}

}

func (t *IncrementPlugin) Version() string {
	return buildVersion()
}

func (t *IncrementPlugin) Symbol() interface{} {
	return t
}

func (t *IncrementPlugin) Desc() string {
	return "Плагин добавляет возможность инкрементальной выгрузки в конфигурации"
}

func (t *IncrementPlugin) Name() string {
	return "increment"
}

func (p *IncrementPlugin) BeforeUpdateCfgHandler(r repository.Repository, dir string) {
	fmt.Println("StartSyncProcess func run")

}

func (p *IncrementPlugin) BeforeStartSyncProcess(r repository.Repository, dir string) {
	fmt.Println("StartSyncProcess func run")

}

func buildVersion() string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
