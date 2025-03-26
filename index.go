package main

import (
	"echo-api/configs"
	"echo-api/services"

	"github.com/labstack/echo/v4"
)

func main() {
	log := services.InitLogger()

	e := echo.New()

	configs.RunDatabase(log)
	services.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
