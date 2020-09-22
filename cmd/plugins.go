package cmd

import (
	"github.com/khorevaa/r2gitsync/plugin"
	"github.com/khorevaa/r2gitsync/plugins"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	disabledPluginsFileName = "disabled-plugins"
	enabledPluginsFileName  = "enabled-plugins"
)

func loadGlobalPlugins(dir string) {

	if ok, _ := IsNoExist(dir); ok {
		return
	}
	err := plugin.LoadPlugins(dir)
	failOnErr(err)

}

func loadLocalPlugins(dir string) {

	if ok, _ := IsNoExist(dir); ok {
		return
	}

	err := plugin.LoadPlugins(dir)
	failOnErr(err)

}

func loadInternalPlugins() {

	err := plugin.Register(plugins.Plugins...)
	failOnErr(err)

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
			failOnErr(err)
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
