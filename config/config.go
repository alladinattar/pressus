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
		UserAgent    string `yaml:"user_agent" envconfig:"USER_AGENT"`
	} `yaml:"parser"`
	Queue struct {
		Ip       string `yaml:"ip" envconfig:"QUEUE_IP"`
		Username string `yaml:"username" envconfig:"QUEUE_USERNAME"`
		Password string `yaml:"password" envconfig:"QUEUE_PASSWORD"`
	} `yaml:"queue"`

	SearchEngine struct {
		Ip   string `yaml:"ip" envconfig:"SEARCH_ENGINE_IP"`
		Port string `yaml:"port" envconfig:"SEARCH_ENGINE_PORT"`
	} `yaml:"search-engine"`

	NERAPI struct {
		IP   string `yaml:"ip" envconfig:"NERAPI_IP"`
		Port string `yaml:"port" envconfig:"NERAPI_PORT"`
	} `yaml:"ner-api"`
}

func ReadConfFile(cfg *Config) error {
	f, err := os.Open("./config.yml")
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
