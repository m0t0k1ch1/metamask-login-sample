package server

import (
	"net/http"

	"github.com/labstack/echo"
)

type Controller struct {
	*echo.Group
	*Core
}

func (cntl *Controller) Child(prefix string) *Controller {
	return &Controller{
		Group: cntl.Group.Group(prefix),
		Core:  cntl.Core,
	}
}

func (cntl *Controller) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	cntl.Add(http.MethodGet, path, h, m...)
}

func (cntl *Controller) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	cntl.Add(http.MethodPost, path, h, m...)
}

func (cntl *Controller) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	cntl.Add(http.MethodPut, path, h, m...)
}

func (cntl *Controller) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	cntl.Add(http.MethodDelete, path, h, m...)
}

func (cntl *Controller) Add(method, path string, h HandlerFunc, m ...echo.MiddlewareFunc) {
	cntl.Group.Add(method, path, func(c echo.Context) error {
		return h(cntl.NewContext(c))
	}, m...)
}

func (cntl *Controller) NewContext(c echo.Context) *Context {
	return &Context{
		Context: c,
		Apps:    cntl.Apps,
	}
}
