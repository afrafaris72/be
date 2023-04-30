package main

import (
	"dumbflix/database"
	"dumbflix/pkg/mysql"
	"dumbflix/routes"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	e := echo.New()
	mysql.DatabaseInit()
	database.RunMigrations()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	fmt.Println("Run at localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
