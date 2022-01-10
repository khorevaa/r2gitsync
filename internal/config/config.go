package config

import "github.com/elastic/go-ucfg"

const DefaultConfig = `
v8version: 8.3
plugins:
  limit:
`

type Config struct {
	WorkDir   string                  `config:"workdir"`
	TempDir   string                  `config:"tempdir"`
	V8version string                  `config:"v8version"`
	V8Path    string                  `config:"v8path"`
	Storage   StorageConfig           `config:"storage"`
	Infobase  *InfobaseConfig         `config:"infobase"`
	Plugins   map[string]*ucfg.Config `config:"plugins"`
}

type StorageConfig struct {
	Path     string `config:"path"`
	User     string `config:"user"`
	Password string `config:"password"`
}

type InfobaseConfig struct {
	ConnectionString string `config:"connection,required"`
	User             string `config:"user,required"`
	Password         string `config:"password"`
}
