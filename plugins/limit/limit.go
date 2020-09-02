package limit

import (
	"fmt"
	"github.com/khorevaa/r2gitsync/internal/opts"
	"github.com/khorevaa/r2gitsync/plugins"
	"strings"
)

// command module
type limitPlugin struct {
	limit      int
	minversion int
	maxversion int
}

func (t *limitPlugin) Init() error {
	fmt.Println("test module loaded OK")
	return nil
}

func (t *limitPlugin) RegistryOptions(name string, cmd main.command) error {

	if !strings.EqualFold(name, "sync") {
		return nil
	}

	opts.IntOpt(cmd, "l limit", 0, "выгрузить не более <Количества> версий от текущей выгруженной").
		Env("GITSYNC_LIMIT R2GITSYNC_LIMIT").
		Ptr(&t.limit)

	opts.IntOpt(cmd, "minversion", 0, "<номер> минимальной версии для выгрузки").
		Ptr(&t.minversion)

	opts.IntOpt(cmd, "maxversion", 0, "<номер> максимальной версии для выгрузки").
		Ptr(&t.maxversion)

	return nil

}

func (t *limitPlugin) RegistryHandlers() map[string]interface{} {

	return nil

}

func (t *limitPlugin) Symbol() interface{} {
	return t
}

func (t *limitPlugin) Version() string {
	return main.buildVersion()
}

func (t *limitPlugin) Desc() string {
	return "Плагин добавляет возможность инкрементальной выгрузки в конфигурации"
}

func (t *limitPlugin) Name() string {
	return "increment"
}

func NewPlugin() plugin.Plugin {
	return new(gzipper)
}
