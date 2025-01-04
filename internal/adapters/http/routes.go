package http

import (
	"github.com/labstack/echo/v4"
	"remember-me/internal/adapters/http/handlers"
	"remember-me/internal/domain/usecases"
)

// TODO ALL usecases SERVICES, user_routes.go ....
func InitRoutes(e *echo.Echo, s *usecases.UserService) {

	userHandler := handlers.NewUserHandler(s)
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.PostUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
	e.PUT("/users/:id", userHandler.UpdateUser)

}
