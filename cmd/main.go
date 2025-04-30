package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/config"
	"github.com/sainp7/go-mongo-echo-lab/db"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"github.com/sainp7/go-mongo-echo-lab/routes"
)

func main() {
	config.LoadConfig()
	cfg := config.Get()
	logger.InitLogger(cfg)
	e := echo.New()
	e.Debug = config.Get().DebugMode
	e.Logger = logger.NewEchoLogger(logger.Log)

	e.Use(logger.MiddlewareLogger(logger.Log))
	mongoDBLogger := logger.NewMongoDBLogger(logger.Log)

	if err := db.InitMongo(cfg, mongoDBLogger); err != nil {
		e.Logger.Fatal(err)
	}
	db.InitCollections(cfg)

	routes.RegisterHealthRoute(e)
	// API V1 Group
	api := e.Group("/api/v1")
	routes.RegisterBookRoutes(api)
	// Start server
	logger.Log.Info().Msg("Server running on :" + cfg.AppPort)
	e.Logger.Fatal(e.Start(":" + cfg.AppPort))
}
