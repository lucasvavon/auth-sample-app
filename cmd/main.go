package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"remember-me/internal/adapters/handlers/http"
	"remember-me/internal/adapters/repositories/postgres"
	"remember-me/internal/domain/models"
	"remember-me/internal/domain/usecases"
)

type Templates struct {
	templates *template.Template
}

// Implement e.Renderer interface
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("cmd/web/views/*.html")),
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

	db := postgres.ConnectDb()
	db = db.Debug()
	// User part
	userStore := postgres.NewUserGORMRepository(db)
	userService := usecases.NewUserService(userStore)
	http.InitRoutes(e, userService)

	userHandler := http.NewUserHandler(userService)

	users := []models.User{
		NewUser("j@gmail.com", "apxmcrz423"),
		NewUser("a@gmail.com", "alicejisd90"),
		NewUser("b@gmail.com", "bob&&isd0"),
	}

	for _, user := range users {
		err := userStore.CreateUser(&user)
		if err != nil {
			fmt.Printf("%s", err)
		}
	}

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return userHandler.GetUsers(c)
	})

	e.POST("/registration", func(c echo.Context) error {

		/*if userHandler.hasEmail(email) {
			formData := newFormData()
			formData.Values["email"] = email
			formData.Values["password"] = password
			formData.Errors["email"] = "Email already exists"

			return c.Render(422, "loginForm", formData)
		}*/

		/*return c.Render(200, "oob-user", user)*/
		return userHandler.PostUser(c)
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		return userHandler.DeleteUser(c)
	})
	
	// Route pour le traitement du formulaire de connexion (POST)
	/* e.POST("/login", handlers.HandleLogin) */

	// Démarrage du serveur
	e.Logger.Fatal(e.Start(":1323"))
}
