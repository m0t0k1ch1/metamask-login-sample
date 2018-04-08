package helpers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const (
	ResponseStateSuccess = "success"
	ResponseStateError   = "error"
)

type Response struct {
	State  string      `json:"state"`
	Result interface{} `json:"result"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *ResponseError) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Message)
}

func NewResponseError(code int, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

func JSONResponseSuccess(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, &Response{
		State:  ResponseStateSuccess,
		Result: result,
	})
}

func JSONResponseError(c echo.Context, err error) error {
	var result *ResponseError
	if re, ok := err.(*ResponseError); ok {
		result = re
	} else {
		result = NewResponseError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &Response{
		State:  ResponseStateError,
		Result: result,
	})
}
