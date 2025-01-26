package config

import (
	"os"
	"time"

	yaml "gopkg.in/yaml.v3"
)

type HTTPServerConfig struct {
	Address     string        `yaml:"address"`
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idletimeout"`
}

type Config struct {
	HTTPServerConfig HTTPServerConfig `yaml:"httpserver"`
}

func MustLoadConfig(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, ErrFailedReadConfigFile
	}
	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, ErrFailedUnmarshalYAML
	}

	return &cfg, nil
}
