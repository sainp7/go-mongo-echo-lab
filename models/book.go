package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title      string             `json:"title"`
	Authors    []string           `json:"authors"`
	ISBN       string             `json:"isbn"`
	Price      float64            `json:"price"`
	Stock      int                `json:"stock"`
	Categories []string           `json:"categories"`
	CreatedAt  time.Time          `json:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt"`
}
