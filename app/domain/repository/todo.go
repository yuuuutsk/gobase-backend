package repository

import (
	"context"

	"github.com/yuuuutsk/gobase-backend/pkg"

	"github.com/yuuuutsk/gobase-backend/app/domain/model"
)

type TodoRepository interface {
	Create(ctx context.Context, tag []*model.Todo, clock pkg.Clock) error
	All(ctx context.Context) ([]*model.Todo, error)
	GetByText(ctx context.Context, text string, clock pkg.Clock) (*model.Todo, error)
}
