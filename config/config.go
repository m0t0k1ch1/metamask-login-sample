package config

import (
	"os"

	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func getenv(key, defaultValue string) string {
	s := os.Getenv(key)
	if s == "" {
		s = defaultValue
	}

	return s
}

func NewServerConfig() *server.Config {
	return &server.Config{
		Port:          getenv("MLS_SERVER_PORT", server.DefaultPort),
		IndexFilePath: getenv("MLS_SERVER_INDEX_FILE_PATH", server.DefaultIndexFilePath),
		StaticDirPath: getenv("MLS_SERVER_STATIC_DIR_PATH", server.DefaultStaticDirPath),
		App:           NewApplicationConfig(),
	}
}

func NewApplicationConfig() *application.Config {
	return &application.Config{
		Secret: getenv("MLS_APP_SECRET", application.DefaultSecret),
	}
}
