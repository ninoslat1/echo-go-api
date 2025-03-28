package main

import (
	"echo-api/services"

	"github.com/labstack/echo/v4"
)

func main() {
	log := services.InitLogger()

	e := echo.New()

	services.SetupRoutes(e, log)

	e.Logger.Fatal(e.Start(":1323"))
}
