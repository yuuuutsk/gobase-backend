package usecase

import (
	"context"

	"github.com/yuuuutsk/gobase-backend/app/domain/models"
	"github.com/yuuuutsk/gobase-backend/pkg"
)

type UserCreateInput struct {
	ID        int
	FirstName string
	LastName  string
	Birthday  string
}

func (uc *UserUseCase) Create(ctx context.Context, input *UserCreateInput, clock pkg.Clock) error {

	user := models.NewUser(input.FirstName, input.LastName)

	err := uc.repository.Create(ctx, []*models.User{user}, clock)
	if err != nil {
		//	todo error handling
	}

	return nil
}
