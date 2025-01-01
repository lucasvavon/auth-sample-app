package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
	"remember-me/internal/adapters/http/handlers"
	"remember-me/internal/adapters/http/middlewares"
	"remember-me/internal/adapters/repositories/postgre"
	"remember-me/internal/adapters/repositories/redis"
	"remember-me/internal/domain/models"
	"remember-me/internal/domain/services"
)

type Templates struct {
	templates *template.Template
}

// Implement e.Renderer interface
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {

	tmpl, err := template.ParseGlob("cmd/web/views/*.gohtml")
	if err != nil {
		fmt.Printf("Error loading templates: %v", err)
	}
	return &Templates{
		templates: tmpl,
	}
}

var id int = 0

func NewUser(email, password string) models.User {
	id++
	return models.User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

type Data struct {
	Users models.Users
}

func (d *Data) indexOf(id int) int {
	for i, user := range d.Users {
		if user.ID == id {
			return i
		}
	}
	return -1
}

func (d *Data) hasEmail(email string) bool {
	for _, user := range d.Users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func main() {

	// Echo instance
	e := echo.New()
	e.Renderer = NewTemplate()
	e.Static("/images", "/cmd/web/assets/images")
	e.Static("/css", "/cmd/web/assets/css")

	db := postgre.ConnectDb()
	db = db.Debug()

	sessionStore := redis.NewRedisSessionRepository()
	sessionService := services.NewSessionService(sessionStore)

	userStore := postgre.NewGormUserRepository(db)
	userService := services.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(userService, sessionService)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Middleware for protected routes
	protected := e.Group("")
	protected.Use(middlewares.AuthMiddleware(sessionService))

	// Route => handler
	protected.GET("/", func(c echo.Context) error { return c.Render(200, "search", nil) })
	e.GET("/registration", func(c echo.Context) error {
		return c.Render(200, "registration", nil)
	})
	e.POST("/registration", func(c echo.Context) error {
		return userHandler.PostUser(c)
	})
	e.GET("/login", func(c echo.Context) error {
		session, err := c.Cookie("session_id") // Or check using your session management method
		if err == nil && session.Value != "" {
			// If the session exists, redirect to the homepage or dashboard
			return c.Redirect(http.StatusFound, "/") // Or "/dashboard"
		}
		return c.Render(200, "login", nil)

	})
	e.POST("/login", func(c echo.Context) error {
		return userHandler.Login(c)
	})
	e.POST("/logout", func(c echo.Context) error {
		return userHandler.Logout(c)
	})

	// delete my account
	/*e.DELETE("/users/:id", func(c echo.Context) error {
		return userHandler.DeleteUser(c)
	})*/

	e.Logger.Fatal(e.Start(":1323"))
}
