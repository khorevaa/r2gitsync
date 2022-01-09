package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	disabledPluginsFileName = "disabled-plugins"
	enabledPluginsFileName  = "enabled-plugins"
	envSepKey               = ","
)

type Config interface {
	Disable(name ...string)
	Enable(name ...string)

	State() map[string]bool
	List() (idxEnable []string, idxDisable []string)

	Load() error
	Save(saveEnable bool) error
}

type FilePluginsConfig struct {
	Dir        string
	Filename   string
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

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
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

func getPluginsFromFile(file string) (pl []string, err error) {

	if ok, _ := Exists(file); ok {
		content, err2 := ioutil.ReadFile(file)
		if err2 != nil {
			return pl, err2
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
