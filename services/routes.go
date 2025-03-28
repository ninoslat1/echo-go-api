package services

import (
	"echo-api/repositories"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(e *echo.Echo, log *logrus.Logger) {
	e.POST("/login", func(c echo.Context) error {
		return repositories.LoginHandler(c, log)
	})
}
