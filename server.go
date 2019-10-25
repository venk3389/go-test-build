package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
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
	return tmpl.ExecuteTemplate(w, "index.html", data)
}

func main() {
	// Echo instance
	e := echo.New()
	d := data{}
	d.New()
	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)
	templates["index.html"] = template.Must(template.ParseFiles("view/index.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}, host=${host}, method=${method}, uri=${uri}, status=${status},latency=${latency_human}\n",
	}))
	// Route => handler
	e.GET("/", HomeHandler)
	e.POST("/url", func(c echo.Context) error {
		url := c.FormValue("url")
		hashStr := randomBase64String(6)
		d.Add(hashStr, url)
		u := &Url{
			UrlString: fmt.Sprintf("http://localhost:8000/%s", hashStr),
		}
		return c.JSON(http.StatusOK, u)
	})
	e.GET("/*", func(c echo.Context) error {
		key := c.Request().URL.Path
		return c.Redirect(http.StatusMovedPermanently, d.Get(key[1:]))
	})

	// Start the Echo server
	e.Logger.Fatal(e.Start(":8000"))
}
