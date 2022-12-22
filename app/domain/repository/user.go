package repository

import (
	"context"

	"github.com/yuuuutsk/gobase-backend/pkg"

	"github.com/yuuuutsk/gobase-backend/app/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, tag []*model.User, clock pkg.Clock) error
	All(ctx context.Context) ([]*model.User, error)
	//GetByText(ctx context.Context, text string, clock pkg.Clock) (*model.User, error)
}
