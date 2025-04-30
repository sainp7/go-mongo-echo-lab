package db

import (
	"github.com/sainp7/go-mongo-echo-lab/config"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

var MongoClient *mongo.Client

func InitMongo(cfg *config.AppConfig, dbLogger *logger.MongoDBLogger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logOptions := options.Logger().
		SetSink(dbLogger).
		SetComponentLevel(options.LogComponentCommand, getLoggerLevel(cfg.LogLevel))

	clientOptions := options.Client().
		ApplyURI(cfg.MongoURI).
		SetLoggerOptions(logOptions)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}
	logger.Log.Info().Msg("Connected to MongoDB")
	logger.Log.Debug().Str("uri", cfg.MongoURI).Msg("MongoDB client connection details")
	MongoClient = client
	return nil
}

func getLoggerLevel(level string) options.LogLevel {
	switch level {
	case "info":
		return options.LogLevelInfo
	case "debug":
		return options.LogLevelDebug
	}
	return 1
}
