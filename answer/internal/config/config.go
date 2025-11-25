package config

import (
	"os"
	"path/filepath"
)

const (
	defaultHttpHost = ""
	defaultHttpPort = "8082"

	defaultConfigDir            = "configs/"
	defaultLoggerConfigFileName = "/configs/logger.yml"
	defaultLoggerLevel          = "debug"
	defaultLoggerOutputPaths    = "stdout"
)

type (
	Config struct {
		HTTP   HTTPConfig
		Logger LoggerConfig
	}
	HTTPConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
	LoggerConfig struct {
		NameConfigFile string   `mapstructure:"name_config_file"`
		Level          string   `mapstructure:"level" yaml:"level"`
		OutputPaths    []string `mapstructure:"output_paths" yaml:"output_paths"`
		EncoderTime    string   `mapstructure:"encode_time" yaml:"encode_time"`
	}
)

func Init() *Config {
	cfg := Config{}

	cfg.HTTP.Port = defaultHttpPort
	cfg.HTTP.Host = defaultHttpHost

	configDir := os.Getenv("CONFIG_DIR")
	if configDir == "" {
		configDir = defaultConfigDir
	}

	loggerConfigPath := filepath.Join(configDir, "logger.yml")
	cfg.Logger.NameConfigFile = loggerConfigPath
	cfg.Logger.Level = defaultLoggerLevel
	cfg.Logger.OutputPaths = []string{
		defaultLoggerOutputPaths,
	}
	return &cfg
}
