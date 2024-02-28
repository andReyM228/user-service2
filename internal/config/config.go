package config

import (
	"github.com/andReyM228/lib/database"
	"github.com/andReyM228/one/chain_client"
	"gopkg.in/yaml.v3"

	"log"
	"os"
)

type (
	Config struct {
		Chain  chain_client.ClientConfig `yaml:"chain"`
		DB     database.DBConfig         `yaml:"db" validate:"required"`
		HTTP   HTTP                      `yaml:"http" validate:"required"`
		Rabbit Rabbit                    `yaml:"rabbit" validate:"required"`
		Extra  Extra                     `yaml:"extra" validate:"required"`
	}

	HTTP struct {
		Port int `yaml:"port" validate:"required"`
	}

	Rabbit struct {
		Url string `yaml:"url" validate:"required"`
	}

	Extra struct {
		CarSystemWallet string `yaml:"car-system-wallet" validate:"required"`
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
