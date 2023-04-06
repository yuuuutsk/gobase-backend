package repository

import (
	"context"

	"github.com/yuuuutsk/gobase-backend/pkg"

	"github.com/yuuuutsk/gobase-backend/app/domain/models"
)

type TodoRepository interface {
	Create(ctx context.Context, tag []*models.Todo, clock pkg.Clock) error
	All(ctx context.Context) ([]*models.Todo, error)
	GetByText(ctx context.Context, text string, clock pkg.Clock) (*models.Todo, error)
}
