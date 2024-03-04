package config

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"os"
)

const configPath = "./config/config.json"

type ConfigPg struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	SSLMode  string
}

type Config struct {
	Server struct {
		Host                        string `validate:"required"`
		IPHeader                    string `validate:"required"`
		ShowUnknownErrorsInResponse bool
		CookieDomain                string `validate:"required"`
	}
	GRPCServer struct {
		Host string `validate:"required"`
	}
	Postgres ConfigPg
}

func LoadConfig() (c *Config, err error) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(jsonFile).Decode(&c)
	if err != nil {
		return nil, err
	}

	err = validator.New().Struct(c)
	if err != nil {
		return nil, err
	}
	return
}
