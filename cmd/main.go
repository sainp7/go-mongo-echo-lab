package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/config"
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
	routes.RegisterRoutes(e)
	logger.Log.Info().Msg("Server running on :" + cfg.AppPort)
	e.Logger.Fatal(e.Start(":" + cfg.AppPort))
}
