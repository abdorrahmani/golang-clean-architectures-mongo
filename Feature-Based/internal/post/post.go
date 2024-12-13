package post

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Post struct {
	ID        bson.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title     string        `json:"title" bson:"title"`
	Body      string        `json:"body" bson:"body"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}
