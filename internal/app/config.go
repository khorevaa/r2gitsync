package app

import (
	"errors"
	"fmt"
	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"
	"github.com/khorevaa/r2gitsync/internal/config"
	"os"
	"path/filepath"
)

func LoadConfig(configFile string) (*config.Config, error) {

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

	var err error
	var rawConfig *ucfg.Config

	if configFile != "" {
		// load ConfigFile
		configFile, err = filepath.Abs(configFile)
		if err != nil {
			panic(err)
		}

		rawConfig, err = ucfg.NewFrom(configFile)
	} else {
		rawConfig, err = yaml.NewConfig([]byte(config.DefaultConfig))
	}
	if err != nil {
		return nil, err
	}

	cfg := &config.Config{}

	if err := rawConfig.Unpack(cfg); err != nil {
		return nil, err
	}

	return cfg, err
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
