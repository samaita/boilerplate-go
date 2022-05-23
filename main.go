package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/samaita/boilerplate-go/config"
	"github.com/samaita/boilerplate-go/internal/handlers"
	appInit "github.com/samaita/boilerplate-go/internal/init"
	"github.com/samaita/boilerplate-go/internal/repositories"
	"github.com/samaita/boilerplate-go/pkg/middleware"
)

var conf config.Config

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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
	router.Use(middleware.Recover(), middleware.Latency)
	root := router.Group("")

	// register handler to root
	healthHandler.RegisterEndpoint(root)

	router.Logger.Fatal(router.Start(conf.App.Port))
}
