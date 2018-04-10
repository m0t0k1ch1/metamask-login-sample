package response

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

func JSONSuccess(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, newSuccessResponse(result))
}

func JSONError(c echo.Context, err error) error {
	var result *domain.Error
	if domainErr, ok := err.(*domain.Error); ok {
		result = domainErr
	} else {
		result = domain.NewUnexpectedError(err.Error())
	}

	return c.JSON(http.StatusOK, newErrorResponse(result))
}
