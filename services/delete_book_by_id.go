package services

import (
	"errors"
	"github.com/sainp7/go-mongo-echo-lab/db"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func DeleteBookByID(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing id")
		return errors.New("invalid id format")
	}
	result, err := db.BookCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error deleting book")
		return err
	}
	if result.DeletedCount == 0 {
		logger.Log.Warn().Msg("Book not found")
		return errors.New("book not found")
	}
	return nil
}
