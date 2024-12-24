package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"remember-me/internal/adapters/handlers/http"
	"remember-me/internal/adapters/repositories/pg"
	"remember-me/internal/domain/models"
	"remember-me/internal/domain/services"
	"remember-me/internal/utils"
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
	e.Use(middleware.Logger())
	e.Renderer = NewTemplate()
	e.Static("/images", "/cmd/web/assets/images")
	e.Static("/css", "/cmd/web/assets/css")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := pg.ConnectDb()
	db = db.Debug()
	// User part
	userStore := pg.NewGormUserRepository(db)
	userService := services.NewUserService(userStore)
	userHandler := http.NewUserHandler(userService)

	/*sessionStore := redis.NewRedisSessionRepository()
	sessionService := services.NewSessionService(sessionStore)
	sessionHandler := http.NewSessionHandler(*sessionService)*/

	// Route => handler
	e.GET("/registration", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	e.POST("/registration", func(c echo.Context) error {
		return userHandler.PostUser(c)
	})

	e.POST("/login", func(c echo.Context) error {
		return userHandler.Login(c)
	})

	/*protected := e.Group("/")
	protected.Use(sessionMiddleware)
	protected.GET("home", home)*/

	/*e.DELETE("/users/:id", func(c echo.Context) error {
		return userHandler.DeleteUser(c)
	})*/

	// Démarrage du serveur
	e.Logger.Fatal(e.Start(":1323"))
}

/*func sessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "session")

		// Vérifier si l'utilisateur est authentifié
		if auth, ok := session.Values["connected"].(bool); !ok || !auth {
			return c.String(401, "Non autorisé.")
		}

		return next(c)
	}
}*/

func LoginValidationMiddleware(us services.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var inputUser models.User

			// Bind the request body to inputUser
			if err := c.Bind(&inputUser); err != nil {
				return c.String(500, err.Error())
			}

			// Check if the user exists by email
			user, err := us.GetUserByEmail(inputUser.Email)
			if err != nil {
				return c.String(401, "Invalid email or password")
			}

			// Check if the password matches
			valid := utils.ComparePassword(inputUser.Password, user.Password)
			if !valid {
				return c.String(401, "Invalid email or password")
			}

			// Save user to context for later use in the handler (e.g., setting session)
			c.Set("user", user)

			return next(c)
		}
	}
}
