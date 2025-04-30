package services

import (
	"github.com/sainp7/go-mongo-echo-lab/db"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"github.com/sainp7/go-mongo-echo-lab/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func CreateBook(book *models.Book) (string, error) {
	result, err := db.BookCollection.InsertOne(context.Background(), book)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error inserting book")
		return "", err
	}
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil

}
