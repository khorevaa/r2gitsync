package plugins

import (
	"os"
	"path/filepath"

	"github.com/khorevaa/r2gitsync/internal/utils"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
)

func LoadPlugins() error {

	appDataDir := utils.GetAppDataDir("r2gitsync")
	pluginsDir := filepath.Join(appDataDir, "plugins")

	if _, err := os.Stat(pluginsDir); err != nil {
		return nil
	}

	err := plugin.LoadPlugins(pluginsDir)

	return err
}
