package services

import (
	context2 "context"
	"github.com/sainp7/go-mongo-echo-lab/db"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"github.com/sainp7/go-mongo-echo-lab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

func GetAllBooks() ([]models.Book, error) {
	cursor, err := db.BookCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error closing cursor")
		}
	}(cursor, context.TODO())

	var books []models.Book
	if err := cursor.All(context.TODO(), &books); err != nil {
		return nil, err
	}
	return books, nil
}
