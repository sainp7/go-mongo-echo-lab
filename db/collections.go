package db

import (
	"github.com/sainp7/go-mongo-echo-lab/config"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

var BookCollection *mongo.Collection

func InitCollections(cfg *config.AppConfig) {
	if MongoClient == nil {
		logger.Log.Fatal().Msg("Mongo client not initialized")
	}
	db := MongoClient.Database(cfg.DBName)
	BookCollection = db.Collection("books")
}
