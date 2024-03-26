package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, echo.Map{
		"data":    data,
		"status":  http.StatusOK,
		"message": "Success",
	})
}

func BadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, echo.Map{
		"status":  400,
		"message": message,
	})
}

func InternalServerError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"status":  500,
		"message": "Server Error",
	})
}

func Unauthorized(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, echo.Map{
		"status":  401,
		"message": "Unauthorized",
	})
}
