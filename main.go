package main

import (
	"github.com/labstack/echo/v4"
	"main.go/pkg/database"
	"main.go/routes"
)

func main() {
	e := echo.New()
	database.GetDB()
	database.InitDB()
	database.Migrate()
	

	routes.SetupRoutes(e.Group(""))
	e.Logger.Fatal(e.Start(":8000"))
}