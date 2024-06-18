package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port   string `yaml:"port"`
	Env    string `yaml:"env" env-default:"local"`
	DBUrl  string `yaml:"db_url"`
	SsoUrl string `yaml:"sso_url"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()

	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg

}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
