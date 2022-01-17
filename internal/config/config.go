package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"
)

func DefaultConfig() *Config {
	rawConfig, err := yaml.NewConfig([]byte(defaultConfig))
	if err != nil {
		panic(err)
	}

	cfg := &Config{}

	if err := rawConfig.Unpack(cfg); err != nil {
		panic(err)
	}

	return cfg
}

const defaultConfig = `
v8version: 8.3
`

type Config struct {
	WorkDir   string                  `config:"workdir"`
	TempDir   string                  `config:"tempdir"`
	V8version string                  `config:"v8version"`
	V8Path    string                  `config:"v8path"`
	Storage   StorageConfig           `config:"storage"`
	Infobase  *InfobaseConfig         `config:"infobase"`
	Plugins   map[string]*ucfg.Config `config:"plugins"`

	Debug            bool `config:"debug"`
	TraceSQLCommands bool
	SQLSlowThreshold time.Duration
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

func LoadConfig(config *Config, configFile string) (*Config, error) {

	if configFile == "" {
		cf, err := resolveConfigFileFromEnv()
		if err == nil {
			configFile = cf
		}
	}

	if configFile == "" {
		cf, err := resolveConfigFileFromWorkDir()
		if err == nil {
			configFile = cf
		}
	}

	if configFile != "" {
		var (
			err       error
			rawConfig *ucfg.Config
		)

		if configFile, err = filepath.Abs(configFile); err != nil {
			return nil, err
		}

		if rawConfig, err = yaml.NewConfigWithFile(configFile); err != nil {
			return nil, err
		}

		cfg, err := ucfg.NewFrom(config)
		if err != nil {
			return nil, err
		}

		if err := cfg.Merge(rawConfig); err != nil {
			return nil, err
		}

		if err := cfg.Unpack(config); err != nil {
			return nil, err
		}

		return config, nil
	}

	return config, nil
}

func resolveConfigFileFromEnv() (string, error) {
	f := os.Getenv("R2GITSYNC_CONFIG_FILE")
	if f == "" {
		return "", errors.New("environment variable 'R2GITSYNC_CONFIG_FILE' is not set")
	}
	return f, nil
}

func resolveConfigFileFromWorkDir() (string, error) {
	matches1, _ := filepath.Glob("r2gitsync.yaml")
	matches2, _ := filepath.Glob("r2gitsync.yml")
	matches := append(matches1, matches2...)
	switch len(matches) {
	case 0:
		return "", errors.New("no config file found in work dir")
	case 1:
		return matches[0], nil
	default:
		panic(fmt.Errorf("multiple config files found %v", matches))
	}
}
