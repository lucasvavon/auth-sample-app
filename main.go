package main

import (
	routes "auth-sample-app/internal/adapters/http"
	"auth-sample-app/internal/adapters/repositories/postgre"
	"auth-sample-app/internal/adapters/repositories/redis"
	"auth-sample-app/internal/domain/usecases"
	"auth-sample-app/web/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Echo instance
	e := echo.New()
	e.Renderer, _ = views.NewTemplate()
	e.Static("/images", "web/assets/images")
	e.Static("/css", "web/assets/css")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// conn to postgresgl instance
	db := postgre.ConnectDb()
	db = db.Debug()

	// init store repositories
	sessionStore := redis.NewRedisSessionRepository()
	userStore := postgre.NewGormUserRepository(db)

	// init services
	sessionService := usecases.NewSessionService(sessionStore)
	userService := usecases.NewUserService(userStore)

	routes.InitRoutes(e, userService, sessionService)

	e.Logger.Fatal(e.Start(":1323"))
}
