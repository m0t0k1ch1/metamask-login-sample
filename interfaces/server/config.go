package server

import (
	"fmt"

	"github.com/m0t0k1ch1/metamask-login-sample/application"
)

type Config struct {
	Port          int                 `json:"port"`
	IndexFilePath string              `json:"index_file_path"`
	StaticDirPath string              `json:"static_dir_path"`
	App           *application.Config `json:"app"`
}

func (conf *Config) Address() string {
	return fmt.Sprintf(":%d", conf.Port)
}
