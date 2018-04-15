package response

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

func JSONSuccess(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, newSuccessResponse(result))
}

func JSONError(c echo.Context, err error) error {
	var result *common.Error
	if commonErr, ok := err.(*common.Error); ok {
		result = commonErr
	} else {
		c.Logger().Error(err)
		result = common.NewUnexpectedError()
	}

	return c.JSON(http.StatusOK, newErrorResponse(result))
}
