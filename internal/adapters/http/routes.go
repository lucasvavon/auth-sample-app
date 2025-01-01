package http

import (
	"github.com/labstack/echo/v4"
	"remember-me/internal/adapters/http/handlers"
	"remember-me/internal/domain/services"
)

// TODO ALL services SERVICES, user_routes.go ....
func InitRoutes(e *echo.Echo, s *services.UserService) {

	userHandler := handlers.NewUserHandler(s)
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.PostUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
	e.PUT("/users/:id", userHandler.UpdateUser)

}
