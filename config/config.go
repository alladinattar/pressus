package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Api struct {
		Port    string `yaml:"port" envconfig:"API_PORT"`
		Version string `yaml:"version" envconfig:"API_VERSION"`
	} `yaml:"api"`
	Parser struct {
		DefaultRoute string `yaml:"default_route" envconfig:"DEFAULT_ROUTE"`
	} `yaml:"parser"`
	Queue struct {
		Ip       string `yaml:"ip" envconfig:"IP"`
		Username string `yaml:"username" `
		Password string `yaml:"password" envconfig:"PASSWORD"`
	} `yaml:"queue"`
}

func ReadConfFile(cfg *Config) error {
	f, err := os.Open("../config.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil

}

func ReadEnv(cfg *Config) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}
	return nil
}
