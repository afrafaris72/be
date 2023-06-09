package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func IniRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlersUser(userRepository)

	e.GET("/users", h.CariUser)
	e.GET("/profile", middleware.Auth(h.CariProfile))
	e.GET("/user/:id", h.DapatUser)
	e.POST("/user", h.MembuatUser)
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.HapusUser)
}
