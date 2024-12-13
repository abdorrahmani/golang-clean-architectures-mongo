package post

import (
	"Hexagonal/internal/domain"
	"Hexagonal/internal/ports"
	"context"
	"time"
)

type PostUseCase struct {
	repo ports.PostRepository
}

func NewPostUseCase(repo ports.PostRepository) *PostUseCase {
	return &PostUseCase{repo: repo}
}

func (uc *PostUseCase) StorePost(ctx context.Context, post *domain.Post) error {
	now := time.Now()
	post.Created = now
	post.Updated = now

	return uc.repo.Create(ctx, post)
}
