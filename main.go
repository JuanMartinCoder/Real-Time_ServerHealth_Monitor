package main

import (
	"io"
	"log"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func RenderRoutes(e *echo.Echo) {
	e.GET(RoutesInstance.MAIN, HandlerMain)
	e.GET(RoutesInstance.MONITOR, HandlerMonitor)
}

func main() {
	log.Println("Server Health Monitor is running...")

	t := &Template{
		templates: template.Must(template.ParseGlob("html/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "html")

	RenderRoutes(e)

	e.Logger.Fatal(e.Start(":1313"))
}
