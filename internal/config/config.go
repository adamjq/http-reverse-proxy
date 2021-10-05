package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Port        string `envconfig:"PORT" default:"5000"`
	ServiceAUrl string `envconfig:"SERVICE_A_URL" default:"http://localhost:5001"`
	ServiceBUrl string `envconfig:"SERVICE_B_URL" default:"http://localhost:5002"`
}

func (c *Config) validate() error {
	if c.Port == "" {
		return errors.New("ERROR: missing PORT in config")
	}
	if c.ServiceAUrl == "" {
		return errors.New("ERROR: missing SERVICE_A_URL in config")
	}
	if c.ServiceBUrl == "" {
		return errors.New("ERROR: missing SERVICE_B_URL in config")
	}
	return nil
}

func New() (*Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config")
	}
	err = cfg.validate()
	if err != nil {
		return nil, errors.Wrap(err, "Invalid config loaded")
	}
	return &cfg, nil
}
