package post

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	db *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *Repository {
	return &Repository{db: db.Collection("posts")}
}
func (r *Repository) Create(post Post) error {
	_, err := r.db.InsertOne(context.Background(), post)
	return err
}

func (r *Repository) Update(id bson.ObjectID, post Post) error {
	_, err := r.db.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": post})
	return err
}

func (r *Repository) Delete(id bson.ObjectID) error {
	_, err := r.db.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *Repository) GetByID(id bson.ObjectID) (Post, error) {
	var post Post
	err := r.db.FindOne(context.Background(), bson.M{"_id": id}).Decode(&post)
	return post, err
}

func (r *Repository) GetAll() ([]Post, error) {
	var posts []Post
	cursor, err := r.db.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &posts)
	return posts, err
}
