package domain

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Post struct {
	ID      bson.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title   string        `json:"title" bson:"title"`
	Body    string        `json:"body" bson:"body"`
	Created time.Time     `json:"created" bson:"created"`
	Updated time.Time     `json:"updated" bson:"updated"`
}
