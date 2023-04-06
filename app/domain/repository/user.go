package repository

import (
	"context"

	"github.com/yuuuutsk/gobase-backend/pkg"

	"github.com/yuuuutsk/gobase-backend/app/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, tag []*models.User, clock pkg.Clock) error
	All(ctx context.Context) ([]*models.User, error)
	//GetByText(ctx context.Context, text string, clock pkg.Clock) (*models.User, error)
}
