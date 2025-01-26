package config

import "errors"

var (
	ErrFailedReadConfigFile = errors.New("failed to read config file")
	ErrFailedUnmarshalYAML  = errors.New("failed to unmarshal YAML")
)
