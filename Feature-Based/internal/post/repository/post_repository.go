package repository

import (
	"Feature-Based/internal/post"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PostRepository struct {
	db *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *PostRepository {
	return &PostRepository{db: db.Collection("posts")}
}

func (r *PostRepository) Create(post post.Post) error {
	_, err := r.db.InsertOne(context.Background(), post)
	return err
}

func (r *PostRepository) Update(id bson.ObjectID, post post.Post) error {
	_, err := r.db.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": post})
	return err
}

func (r *PostRepository) Delete(id bson.ObjectID) error {
	_, err := r.db.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *PostRepository) GetByID(id bson.ObjectID) (post.Post, error) {
	var post post.Post
	err := r.db.FindOne(context.Background(), bson.M{"_id": id}).Decode(&post)
	return post, err
}

func (r *PostRepository) GetAll() ([]post.Post, error) {
	var post []post.Post
	cursor, err := r.db.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &post)
	return post, err
}
