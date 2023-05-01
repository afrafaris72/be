package main

import (
	"dumbflix/database"
	"dumbflix/pkg/mysql"
	"dumbflix/routes"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	e := echo.New()
	mysql.DatabaseInit()
	database.RunMigrations()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                                 // mengijinkan akses semuanya
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},        // mengijinkan method apa aja yg bisa digunakan
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"}, // mengijinkan headers apa aja yg bisa digunakan
	}))

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	fmt.Println("Run at localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
