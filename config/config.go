package config

import "os"

const (
	DefaultSecret = "secret"

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
	Secret string
	Server *ServerConfig
}

func NewConfig() *Config {
	secret := getenv("MLS_SECRET", DefaultSecret)

	return &Config{
		Secret: secret,
		Server: NewServerConfig(),
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
