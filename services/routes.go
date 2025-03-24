package services

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		userAgent := c.Request().Header.Get("User-Agent")
		response := fmt.Sprintf("Hello, World! from %s", userAgent)
		Log.Infof("API Access from %s", userAgent)
		return c.String(http.StatusOK, response)
	})
}
