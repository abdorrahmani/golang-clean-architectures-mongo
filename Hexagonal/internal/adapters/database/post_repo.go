package database

import (
	"Hexagonal/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PostRepositoryImpl struct {
	db *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db.Collection("posts")}
}

func (r *PostRepositoryImpl) Create(ctx context.Context, post *domain.Post) error {
	_, err := r.db.InsertOne(ctx, post)
	return err
}
