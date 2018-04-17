package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler"
)

type Group struct {
	group  *echo.Group
	server *Server
}

func (g *Group) Group(prefix string) *Group {
	return &Group{
		group:  g.group.Group(prefix),
		server: g.server,
	}
}

func (g *Group) Use(m ...echo.MiddlewareFunc) {
	g.group.Use(m...)
}

func (g *Group) GET(path string, h handler.HandlerFunc) {
	g.Add(http.MethodGet, path, h)
}

func (g *Group) POST(path string, h handler.HandlerFunc) {
	g.Add(http.MethodPost, path, h)
}

func (g *Group) Add(method, path string, h handler.HandlerFunc) {
	g.group.Add(method, path, func(ec echo.Context) error {
		return h(&handler.Context{
			Context: ec,
			Core:    application.NewCore(g.server.container, g.server.config.App),
		})
	})
}
