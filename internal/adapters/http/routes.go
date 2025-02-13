package http

import (
	"auth-sample-app/internal/adapters/http/handlers"
	"auth-sample-app/internal/adapters/http/middlewares"
	"auth-sample-app/internal/domain/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRoutes(e *echo.Echo, us *usecases.UserService, ss *usecases.SessionService) {

	userHandler := handlers.NewUserHandler(us, ss)
	// Middleware for protected routes
	protected := e.Group("", middlewares.AuthMiddleware(ss))

	// Route => handler
	protected.GET("/", func(c echo.Context) error {
		userID := c.Get("userID")
		data := map[string]interface{}{
			"userID": userID,
		}
		return c.Render(200, "index", data)
	})
	e.GET("/registration", func(c echo.Context) error {
		return c.Render(200, "registration", nil)
	})
	e.POST("/registration", func(c echo.Context) error {
		return userHandler.PostUser(c)
	})
	e.GET("/login", func(c echo.Context) error {
		session, err := c.Cookie("session_id")
		if err == nil && session.Value != "" {
			return c.Redirect(http.StatusFound, "/")
		}
		return c.Render(200, "login", nil)

	})
	e.POST("/login", func(c echo.Context) error {
		return userHandler.Login(c)
	})
	protected.POST("/logout", func(c echo.Context) error {
		return userHandler.Logout(c)
	})
}
