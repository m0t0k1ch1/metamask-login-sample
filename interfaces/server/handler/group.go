package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

type Group struct {
	*echo.Group
	Config *server.Config
	Core   *application.Core
}

func (g *Group) NewGroup(prefix string) *Group {
	return &Group{
		Group: g.Group.Group(prefix),
		Core:  g.Core,
	}
}

func (g *Group) Use(m ...echo.MiddlewareFunc) {
	g.Group.Use(m...)
}

func (g *Group) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodGet, path, h, m...)
}

func (g *Group) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodPost, path, h, m...)
}

func (g *Group) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodPut, path, h, m...)
}

func (g *Group) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	g.Add(http.MethodDelete, path, h, m...)
}

func (g *Group) Add(method, path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	g.Group.Add(method, path, func(c echo.Context) error {
		return h(NewContext(c, g.Core))
	}, m...)
}
