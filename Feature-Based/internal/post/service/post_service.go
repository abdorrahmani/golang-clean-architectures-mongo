package service

import (
	"Feature-Based/internal/post"
	"Feature-Based/internal/post/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post post.Post) error {
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now
	return s.repo.Create(post)
}

func (s *PostService) UpdatePost(id bson.ObjectID, post *post.Post) error {
	existingPost, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	post.CreatedAt = existingPost.CreatedAt
	post.UpdatedAt = time.Now()

	return s.repo.Update(id, *post)
}

func (s *PostService) DeletePost(id bson.ObjectID) error {
	return s.repo.Delete(id)
}

func (s *PostService) GetPostByID(id bson.ObjectID) (post.Post, error) {
	return s.repo.GetByID(id)
}

func (s *PostService) GetAllPosts() ([]post.Post, error) {
	return s.repo.GetAll()
}
