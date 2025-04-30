package services

import (
	"errors"
	"github.com/sainp7/go-mongo-echo-lab/db"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"github.com/sainp7/go-mongo-echo-lab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func GetBookByID(id string) (*models.Book, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing id")
		return nil, errors.New("invalid id format")
	}
	var book models.Book
	if err = db.BookCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&book); err != nil {
		logger.Log.Error().Err(err).Msg("Error finding book")
		return nil, err
	}
	return &book, nil
}
