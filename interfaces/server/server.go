package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

type Server struct {
	*Controller
	e *echo.Echo
}

func New(core *Core) *Server {
	srv := &Server{
		e: echo.New(),
	}
	srv.Controller = &Controller{
		Group: srv.e.Group(""),
		Core:  core,
	}

	srv.e.Logger.SetLevel(core.Config.LogLvl())
	srv.e.HTTPErrorHandler = srv.httpErrorHandler

	return srv
}

func (srv *Server) Base() *Controller {
	return srv.Controller
}

func (srv *Server) Logger() echo.Logger {
	return srv.e.Logger
}

func (srv *Server) Start() error {
	return srv.e.Start(srv.Config.Address())
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.e.Shutdown(ctx)
}

func (srv *Server) httpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if httpErr, ok := err.(*echo.HTTPError); ok {
		code = httpErr.Code
		msg = fmt.Sprintf("%v", httpErr.Message)
	}

	appErr := domain.NewError(code, msg)
	srv.e.Logger.Error(appErr)

	if !c.Response().Committed {
		if err := c.JSON(http.StatusOK, NewErrorResponse(appErr)); err != nil {
			srv.e.Logger.Error(err)
		}
	}
}
