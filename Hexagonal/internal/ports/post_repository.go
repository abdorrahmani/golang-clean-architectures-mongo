package ports

import (
	"Hexagonal/internal/domain"
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
}
