package server

import "github.com/m0t0k1ch1/metamask-login-sample/application"

const (
	DefaultPort          = "1323"
	DefaultIndexFilePath = "index.html"
	DefaultStaticDirPath = "static"
)

type Config struct {
	Port          string
	IndexFilePath string
	StaticDirPath string
	App           *application.Config
}
