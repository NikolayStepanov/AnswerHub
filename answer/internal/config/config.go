package config

const (
	defaultHttpHost = ""
	defaultHttpPort = "8082"

	defaultLoggerConfigFileName = "configs/logger.yml"
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

	cfg.Logger.NameConfigFile = defaultLoggerConfigFileName
	cfg.Logger.Level = defaultLoggerLevel
	cfg.Logger.OutputPaths = []string{
		defaultLoggerOutputPaths,
	}
	return &cfg
}
