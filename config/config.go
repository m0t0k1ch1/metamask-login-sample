package config

import (
	"os"
)

const (
	DefaultAppSecret = "secret"

	DefaultServerPort          = "1323"
	DefaultServerIndexFilePath = "index.html"
	DefaultServerStaticDirPath = "static"
)

func getenv(key, defaultValue string) string {
	s := os.Getenv(key)
	if s == "" {
		s = defaultValue
	}

	return s
}

type Config struct {
	App    *AppConfig
	Server *ServerConfig
}

func NewConfig() *Config {
	return &Config{
		App:    NewAppConfig(),
		Server: NewServerConfig(),
	}
}

type AppConfig struct {
	Secret string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Secret: getenv("MLS_APP_SECRET", DefaultAppSecret),
	}
}

type ServerConfig struct {
	Port          string
	IndexFilePath string
	StaticDirPath string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:          getenv("MLS_SERVER_PORT", DefaultServerPort),
		IndexFilePath: getenv("MLS_SERVER_INDEX_FILE_PATH", DefaultServerIndexFilePath),
		StaticDirPath: getenv("MLS_SERVER_STATIC_DIR_PATH", DefaultServerStaticDirPath),
	}
}
