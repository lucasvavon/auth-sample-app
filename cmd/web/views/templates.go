package views

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
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
