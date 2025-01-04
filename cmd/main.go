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
	"remember-me/internal/domain/usecases"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if data == nil {
		data = map[string]interface{}{}
	}

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("template data must be a map[string]interface{}")
	}

	// global variables to render context
	dataMap["userID"] = c.Get("userID")

	return t.templates.ExecuteTemplate(w, name, dataMap)
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

func main() {

	/*conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	log.Printf("NETWORK ACCESS : http://%v:1323", localAddr.IP)*/

	// Echo instance
	e := echo.New()
	e.Renderer = NewTemplate()
	e.Static("/images", "/cmd/web/assets/images")
	e.Static("/css", "/cmd/web/assets/css")

	db := postgre.ConnectDb()
	db = db.Debug()

	sessionStore := redis.NewRedisSessionRepository()
	sessionService := usecases.NewSessionService(sessionStore)

	userStore := postgre.NewGormUserRepository(db)
	userService := usecases.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(userService, sessionService)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Middleware for protected routes
	protected := e.Group("", middlewares.AuthMiddleware(sessionService))

	// Route => handler
	protected.GET("/", func(c echo.Context) error { return c.Render(200, "search", nil) })
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

	// delete my account
	/*e.DELETE("/users/:id", func(c echo.Context) error {
		return userHandler.DeleteUser(c)
	})*/

	e.Logger.Fatal(e.Start(":1323"))

}
