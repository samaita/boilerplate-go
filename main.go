package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samaita/boilerplate-go/config"
	"github.com/samaita/boilerplate-go/internal/handlers"
	appInit "github.com/samaita/boilerplate-go/internal/init"
	"github.com/samaita/boilerplate-go/internal/repositories"
)

var conf config.Config

func init() {
	conf = config.GetConfig()
}

func main() {

	// initiate DB & Cache connection
	DBConn := appInit.ConnectDB(conf)
	RedisConn := appInit.ConnectCache(conf)

	// initiate repositories
	healthRepo := repositories.NewHealthRepo(DBConn.MainDB, RedisConn.MainCache)

	// initiate handlers
	healthHandler := handlers.NewHealthHandler(healthRepo)

	router := echo.New()
	router.Use(middleware.Recover())
	root := router.Group("")

	// register handler to root
	healthHandler.RegisterEndpoint(root)

	router.Logger.Fatal(router.Start(conf.App.Port))
}
