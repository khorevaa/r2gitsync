package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	pm "plugin"
	"strings"
	"sync"
)

const (
	disabledPluginsFileName = "disabled-plugins"
	enabledPluginsFileName  = "enabled-plugins"
	envSepKey               = ","
)

type PluginsConfig interface {
	Disable(name ...string)
	Enable(name ...string)

	State() map[string]bool
	List() (idxEnable []string, idxDisable []string)

	Load() error
	Save(saveEnable bool) error
}

type FilePluginsConfig struct {
	Dir      string
	Filename string

	IdxEnable  map[string]bool
	IdxDisable map[string]bool
}

func (s *FilePluginsConfig) State() (pl map[string]bool) {

	pl = make(map[string]bool)

	for name, _ := range s.IdxDisable {
		pl[name] = false
	}

	for name, _ := range s.IdxEnable {
		pl[name] = true
	}

	return

}

func (s *FilePluginsConfig) Disable(names ...string) {

	for _, name := range names {

		s.changeEnable(name, false)

	}

}

func (s *FilePluginsConfig) Enable(names ...string) {

	for _, name := range names {

		s.changeEnable(name, true)

	}

}

func (s *FilePluginsConfig) List() (idxEnable []string, idxDisable []string) {

	for name, _ := range s.IdxEnable {
		idxEnable = append(idxEnable, name)
	}

	for name, _ := range s.IdxDisable {
		idxDisable = append(idxDisable, name)
	}

	return
}

func (s *FilePluginsConfig) Load() error {

	if len(s.Dir) == 0 {
		return nil
	}

	s.loadPluginsState()

	return nil
}

func (s *FilePluginsConfig) Save(saveEnable bool) error {

	plEnable, plDisable := s.List()
	data := strings.Join(plDisable, "\n")

	file := filepath.Join(s.Dir, disabledPluginsFileName)
	err := ioutil.WriteFile(file, []byte(data), 0644)

	if saveEnable {

		data2 := strings.Join(plEnable, "\n")

		file2 := filepath.Join(s.Dir, enabledPluginsFileName)
		err = ioutil.WriteFile(file2, []byte(data2), 0644)

	}

	return err
}

func (s *FilePluginsConfig) changeEnable(name string, enable bool) {

	enableState := s.IdxEnable[name]
	disableState := s.IdxEnable[name]

	switch enable {

	case true:
		if enableState {
			delete(s.IdxEnable, name)
		}

		s.IdxDisable[name] = true

	case false:

		if disableState {
			delete(s.IdxDisable, name)
		}

		s.IdxEnable[name] = true
	}

}

func (s *FilePluginsConfig) loadPluginsState() {

	fileEnable := filepath.Join(s.Dir, enabledPluginsFileName)

	plEnable, err := getPluginsFromFile(fileEnable)

	if err != nil {
		return
	}

	for _, name := range plEnable {

		s.IdxEnable[name] = true
	}

	fileDisable := filepath.Join(s.Dir, disabledPluginsFileName)

	plDisable, err := getPluginsFromFile(fileDisable)

	if err != nil {
		return
	}

	for _, name := range plDisable {

		s.IdxDisable[name] = true
	}

}

type PluginStorage interface {
	Load() error
	Install(filename string) error
	Delete(name string) error
	List() []StoragePlugin
	Clear() error
	Registry(m *manager) error
}

type StoragePlugin struct {
	Symbol
	File string
	hash string
}

type FilePluginStorage struct {
	mu  sync.Mutex
	Dir string
	Idx map[string]StoragePlugin
}

func NewFilePluginStorage(dir string) *FilePluginStorage {

	return &FilePluginStorage{
		Dir: dir,
		Idx: make(map[string]StoragePlugin),
	}
}

func (s *FilePluginStorage) Install(filename string) error {

	return nil

}

func (s *FilePluginStorage) Delete(name string) error {

	return nil

}

func (s *FilePluginStorage) Registry(mng *manager) error {

	for _, storagePlugin := range s.Idx {
		_ = mng.Register(storagePlugin)
	}

	return nil
}

func (s *FilePluginStorage) List() (ls []StoragePlugin) {

	for _, storagePlugin := range s.Idx {
		ls = append(ls, storagePlugin)
	}

	return
}

func (s *FilePluginStorage) Load() error {

	if len(s.Dir) == 0 {
		return nil
	}
	if _, err := os.Stat(s.Dir); err != nil {
		return err
	}

	plugins, err := listFiles(s.Dir, `*.so`)
	if err != nil {
		return err
	}

	for _, cmdPlugin := range plugins {
		plFile := filepath.Join(s.Dir, cmdPlugin.Name())
		pluginFile, err := pm.Open(plFile)
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

		pl := StoragePlugin{
			Symbol: pluginSymbol.(Symbol),
			File:   plFile,
			hash:   "", // TODO Расчет хеша файла
		}

		s.Idx[pl.Name()] = pl
	}

	return nil
}

func (s *FilePluginStorage) Clear() error {

	return clearDir(s.Dir)
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

func getPluginsFromEnv(env string) (pl []string, err error) {

	value := getEnv(env)

	if value == "" {
		return
	}

	pl = strings.Split(value, envSepKey)
	return
}

func getPluginsFromFile(file string) (pl []string, err error) {

	if ok, _ := Exists(file); ok {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return
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

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func IsNoExist(name string) (bool, error) {

	ok, err := Exists(name)
	return !ok, err
}

func clearDir(dir string) error {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {

		err = os.Remove(f.Name())
		if err != nil {
			return err
		}
		//fmt.Println(f.Name())
	}

	return nil
}
