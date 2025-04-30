package db

import (
	"github.com/sainp7/go-mongo-echo-lab/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var BookCollection *mongo.Collection

func InitCollections(cfg *config.AppConfig) {
	db := MongoClient.Database(cfg.DBName)
	BookCollection = db.Collection("books")
}
