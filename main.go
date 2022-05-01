package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samaita/boilerplate-go/config"
	appInit "github.com/samaita/boilerplate-go/internal/init"
)

var conf config.Config

func init() {
	conf = config.GetConfig()
}

func main() {

	appInit.ConnectDB(conf)
	appInit.ConnectCache(conf)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(conf.App.Port))
}
