package post

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Post struct {
	ID        bson.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title     string        `json:"title" bson:"title"`
	Body      string        `json:"body" bson:"body"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
