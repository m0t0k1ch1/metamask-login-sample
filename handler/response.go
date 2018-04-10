package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
)

const (
	ResponseStateSuccess = "success"
	ResponseStateError   = "error"
)

type Response struct {
	State  string      `json:"state"`
	Result interface{} `json:"result"`
}

func newSuccessResponse(result interface{}) *Response {
	return &Response{
		State:  ResponseStateSuccess,
		Result: result,
	}
}

func newErrorResponse(err *model.Error) *Response {
	return &Response{
		State:  ResponseStateError,
		Result: err,
	}
}

func jsonResponseSuccess(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, newSuccessResponse(result))
}

func jsonResponseError(c echo.Context, err error) error {
	var result *model.Error
	if re, ok := err.(*model.Error); ok {
		result = re
	} else {
		result = model.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newErrorResponse(result))
}
