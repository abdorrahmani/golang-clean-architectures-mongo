package post

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type UseCase struct {
	repo *Repository
}

func NewPostUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) CreatePost(post *Post) error {
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now
	return uc.repo.Create(*post)
}

func (uc *UseCase) UpdatePost(id bson.ObjectID, post *Post) error {
	post.UpdatedAt = time.Now()
	return uc.repo.Update(id, *post)
}
func (uc *UseCase) DeletePost(id bson.ObjectID) error {
	return uc.repo.Delete(id)
}

func (uc *UseCase) GetPostByID(id bson.ObjectID) (Post, error) {
	return uc.repo.GetByID(id)
}

func (uc *UseCase) GetAllPosts() ([]Post, error) {
	return uc.repo.GetAll()
}
