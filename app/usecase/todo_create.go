package usecase

import (
	"context"

	"github.com/yuuuutsk/gobase-backend/app/domain/models"
	"github.com/yuuuutsk/gobase-backend/pkg"
)

type TodoCreateInput struct {
	Text string
	Done bool
}

func (uc *TodoUseCase) Create(ctx context.Context, input *TodoCreateInput, clock pkg.Clock) error {

	todo := models.NewTodo(input.Text, input.Done, 1)

	err := uc.repository.Create(ctx, []*models.Todo{todo}, clock)
	if err != nil {
		//	todo error handling
	}

	return nil
}
