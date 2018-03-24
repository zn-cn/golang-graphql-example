package conf

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// config
type config struct {
	Env      string `toml:"env"`
	LogLevel string `toml:"log_level"`
	APP      app    `toml:"app"`
	DB       db     `toml:"db"`
	JWT      jwt    `toml:"jwt"`
}

type jwt struct {
	SignMethod string `toml:"sign_method"`
	Secret     string `toml:"secret"`
	Exp        int    `toml:"exp"`
}

// ENVMAP map[string]string
var ENVMAP = map[string]string{"dev": "conf/devConfig.toml", "pro": "conf/proConfig.toml"}

// Config config
var Config config

// InitConfig initializes the app configuration by the env("dev", "pro")
func InitConfig(env string) error {
	configFile, ok := ENVMAP[env]
	if !ok {
		return errors.New("config file err: don't have the config env")
	}
	// Set default config.
	Config = config{
		Env:      env,
		LogLevel: "DEBUG",
	}
	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	}
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.New("config load err:" + err.Error())
	}
	_, err = toml.Decode(string(configBytes), &Config)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}

	return nil
}
