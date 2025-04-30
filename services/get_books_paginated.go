package services

import (
	context2 "context"
	"github.com/sainp7/go-mongo-echo-lab/db"
	"github.com/sainp7/go-mongo-echo-lab/logger"
	"github.com/sainp7/go-mongo-echo-lab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func GetBooksPaginated(pageNumber, pageLimit int) ([]models.Book, int64, error) {
	skip := pageNumber * pageLimit

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageLimit))
	findOptions.SetSort(bson.D{{"createdAt", -1}})

	cursor, err := db.BookCollection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error finding books")
		return nil, 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error closing cursor")
		}
	}(cursor, context.TODO())

	var books []models.Book
	if err := cursor.All(context.TODO(), &books); err != nil {
		logger.Log.Error().Err(err).Msg("Error fetching books")
		return nil, 0, err
	}
	total, err := db.BookCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error counting books")
		return nil, 0, err
	}
	return books, total, nil
}
