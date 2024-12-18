package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var users = map[string]string{
	"admin": "password123", // username: password
	"user":  "1234",
}

func HandleLogin(c echo.Context) error {
	// Récupération des données du formulaire
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Vérification des identifiants
	if pass, ok := users[username]; ok && pass == password {
		return c.String(http.StatusOK, "Login successful! Welcome, "+username)
	}

	// Si les identifiants sont incorrects
	return c.String(http.StatusUnauthorized, "Invalid username or password")
}

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
