package services

import (
	"errors"
	"github.com/sainp7/go-mongo-echo-lab/db"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"github.com/sainp7/go-mongo-echo-lab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"time"
)

func UpdateBookByID(id string, book models.Book) (*models.Book, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing id")
		return nil, errors.New("invalid id format")
	}
	book.ID = objID
	book.UpdatedAt = time.Now()

	filter := bson.M{"_id": objID}

	update := bson.M{"$set": book}
	res, err := db.BookCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		return nil, errors.New("book not found")
	}
	return &book, nil
}
