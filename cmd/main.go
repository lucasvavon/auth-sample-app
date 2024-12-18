package main

import (
	"remember-me/internal/web/views"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DB struct{}

var id int = 0

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func NewUser(username, email, password string) User {
	id++
	return User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	}
}

type Users = []User

type Data struct {
	Users Users
}

// ??????????????????????
func (d *Data) indexOf(id int) int {
	for i, user := range d.Users {
		if user.Id == id {
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

func newData() Data {
	return Data{
		Users: []User{
			NewUser("John", "j@gmail.com", "apxmcrz423"),
			NewUser("Alice", "a@gmail.com", "alicejisd90"),
			NewUser("Bob", "b@gmail.com", "bob&&isd0"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func main() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = views.NewTemplate()
	e.Static("/images", "images")
	e.Static("/css", "css")

	page := newPage()

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.POST("/users", func(c echo.Context) error {
		email := c.FormValue("email")
		username := c.FormValue("username")
		password := c.FormValue("password")

		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["username"] = username
			formData.Values["email"] = email
			formData.Values["password"] = password
			formData.Errors["email"] = "Email already exists"

			return c.Render(422, "loginForm", formData)
		}

		user := NewUser(username, email, password)

		page.Data.Users = append(page.Data.Users, user)

		c.Render(200, "loginForm", newFormData())
		return c.Render(200, "oob-user", user)
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		time.Sleep(1 * time.Second)
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(400, "Invalid id")
		}

		index := page.Data.indexOf(id)
		if index == -1 {
			return c.String(404, "User not found")
		}

		page.Data.Users = append(page.Data.Users[:index], page.Data.Users[index+1:]...)

		return c.NoContent(200)
	})

	// Route pour le traitement du formulaire de connexion (POST)
	/* e.POST("/login", handlers.HandleLogin) */

	// Démarrage du serveur
	e.Logger.Fatal(e.Start(":1323"))
}
