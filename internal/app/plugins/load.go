package plugins

import (
	"github.com/khorevaa/r2gitsync/internal/plugins"
	"github.com/khorevaa/r2gitsync/internal/utils"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func initPluginsDirs(config *configApp) {

	appDataDir := utils.GetAppDataDir("r2gitsync")
	config.Plugins.GlobalDir = filepath.Join(appDataDir, "plugins")

	localDir := getEnv(PluginsDirEnv)

	if len(localDir) == 0 {
		localDir = pluginsDirPwd
	}

	config.Plugins.LocalDir = localDir

}

func LoadPlugins(config *configApp) error {

	if err := loadInternalPlugins(); err != nil {
		return err
	}
	if err := loadGlobalPlugins(config.Plugins.GlobalDir); err != nil {
		return err
	}
	if err := loadLocalPlugins(config.Plugins.LocalDir); err != nil {
		return err
	}

}

func LoadDisabledPlugins(config *configApp) {

	loadGlobalDisabledPlugins(config.Plugins.GlobalDir)
	loadLocalDisabledPlugins(config.Plugins.LocalDir)
	loadLocalEnabledPlugins(config.Plugins.LocalDir)
	loadDisabledPluginsEnv()

}

const (
	disabledPluginsFileName = "disabled-plugins"
	enabledPluginsFileName  = "enabled-plugins"
)

func loadGlobalPlugins(dir string) error {

	if ok, _ := IsNoExist(dir); ok {
		return nil
	}
	err := plugin.LoadPlugins(dir)

	return err

}

func loadLocalPlugins(dir string) error {

	if ok, _ := IsNoExist(dir); ok {
		return nil
	}

	err := plugin.LoadPlugins(dir)
	return err

}

func loadInternalPlugins() error {

	err := plugin.Register(plugins.Plugins...)
	return err

}

func getEnv(envs ...string) string {

	for _, env := range envs {

		keys := strings.Fields(env)

		for _, key := range keys {
			value := strings.TrimSpace(os.Getenv(key))

			if len(value) > 0 {
				return value
			}
		}

	}

	return ""

}

func loadLocalEnabledPlugins(dir string) {

	filename := filepath.Join(filepath.Dir(dir), disabledPluginsFileName)
	pl := getPluginsFromFile(filename)

	plugin.Enable(pl...)

}

func loadLocalDisabledPlugins(dir string) {

	filename := filepath.Join(filepath.Dir(dir), disabledPluginsFileName)
	pl := getPluginsFromFile(filename)

	plugin.Disable(pl...)

}

func getPluginsFromFile(file string) (pl []string) {

	if ok, _ := Exists(file); ok {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			FailOnErr(err)
		}
		lines := strings.Split(string(content), "\n")

		for _, line := range lines {
			if strings.HasPrefix(strings.TrimSpace(line), "//") {
				continue
			}
			pl = append(pl, line)
		}

	}

	return
}

func loadGlobalDisabledPlugins(dir string) {

	filename := filepath.Join(filepath.Dir(dir), disabledPluginsFileName)
	pl := getPluginsFromFile(filename)

	plugin.Disable(pl...)

}

func loadDisabledPluginsEnv() {

	pl := getEnv("R2GITSYNC_DISABLE_PLUGINS")
	plugin.Disable(strings.Split(pl, ",")...)

}

func saveGlobalDisabledPlugins(dir string, pl []string) error {

	return nil
}

func saveLocalDisabledPlugins(dir string, pl []string) error {

	return nil
}

func saveLocalEnabledPlugins(dir string, pl []string) error {

	return nil
}
