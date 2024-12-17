package main

import (
	"errors"
	"html/template"
	"io"
	"remember-me/handler"

	"github.com/labstack/echo/v4"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	// Echo instance
	e := echo.New()

	// Instantiate a template registry with an array of template set
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("view/home.html", "view/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Route => handler
	// Route pour la page de connexion (GET)
	e.GET("/", handler.HomeHandler)

	// Route pour le traitement du formulaire de connexion (POST)
	e.POST("/login", handler.HandleLogin)

	// Démarrage du serveur
	e.Logger.Fatal(e.Start(":1323"))
}
