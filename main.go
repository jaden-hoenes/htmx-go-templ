package main

import (
	// "context"
	"os"

	"github.com/jaden-hoenes/htmx-go-templ/constants/env"
	"github.com/jaden-hoenes/htmx-go-templ/constants/file"
	"github.com/jaden-hoenes/htmx-go-templ/constants/url"
	"github.com/jaden-hoenes/htmx-go-templ/templates/pages"
	// "github.com/jaden-hoenes/go-htmx-templ/utilities"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File(url.Favicon, file.Favicon)

	pages.IndexHandler.Register(e)

	port := os.Getenv(env.Port)
	e.Logger.Info(e.Start(":" + port))
}
