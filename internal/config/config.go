package config

import (
	"github.com/andReyM228/lib/database"
	"github.com/andReyM228/one/chain_client"
	"gopkg.in/yaml.v3"

	"log"
	"os"
)

// TODO: добавить валидацию конфига

type (
	Config struct {
		Chain  chain_client.ClientConfig `yaml:"chain"`
		DB     database.DBConfig         `yaml:"db"`
		HTTP   HTTP                      `yaml:"http"`
		Rabbit Rabbit                    `yaml:"rabbit"`
		Extra  Extra                     `yaml:"extra"`
	}

	HTTP struct {
		Port int `yaml:"port"`
	}

	Rabbit struct {
		Url string `yaml:"url"`
	}

	Extra struct {
		CarSystemWallet string `yaml:"car-system-wallet"`
	}
)

func ParseConfig() (Config, error) {
	file, err := os.ReadFile("./cmd/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}
