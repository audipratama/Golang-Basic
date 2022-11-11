package config

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"os"
	"sync"
)

const (
	dbModule   = "db"
	mainModule = "main"
)

var (
	config    *Config
	once      sync.Once
	confPaths = getConfPaths()
)

func getConfPaths() string {
	return "files/"
}

func NewModuleConfig() *Config {
	once.Do(func() {
		dbConfig := new(DBConfig)
		err := parseConfig(dbConfig, dbModule, confPaths)
		if err != nil {
			panic(fmt.Sprintf("unable to resolve db config files, err: %s", err.Error()))
		}

		mainConfig := new(MainConfig)
		err = parseConfig(mainConfig, mainModule, confPaths)
		if err != nil {
			panic(fmt.Sprintf("unable to resolve main config files, err: %s", err.Error()))
		}

		config = &Config{
			DB:   dbConfig,
			Main: mainConfig,
		}
	})

	return config
}

//parseConfig will parse module config
func parseConfig(cfg interface{}, module string, path string) error {
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "development"
	}

	if confPaths == "test"{
		return nil
	}

	return readModuleConfig(cfg, environment, module, path)
}

//readModuleConfig will read config from .ini file
func readModuleConfig(cfg interface{}, env, module string, path string) error {
	var err error
	fname := path + "/" + module + "." + env + ".ini"
	err = gcfg.ReadFileInto(cfg, fname)

	return err
}
