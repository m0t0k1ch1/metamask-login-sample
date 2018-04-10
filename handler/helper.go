package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
)

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
