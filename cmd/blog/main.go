package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/polaris1119/embedexample"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	tpl := &Template{
		templates: template.Must(template.New("index").ParseFS(embedexample.TemplateFS, "template/*.html")),
	}
	e.Renderer = tpl

	e.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(embedexample.StaticAsset))))

	e.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "index.html", nil)
	})

	e.Logger.Fatal(e.Start(":2020"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
