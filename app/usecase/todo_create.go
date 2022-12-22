package usecase

import (
	"context"

	"github.com/yuuuutsk/gobase-backend/app/domain/model"
	"github.com/yuuuutsk/gobase-backend/pkg"
)

type TodoCreateInput struct {
	Text string
	Done bool
}

func (uc *TodoUseCase) Create(ctx context.Context, input *TodoCreateInput, clock pkg.Clock) error {

	todo := model.NewTodo(input.Text, input.Done)

	err := uc.repository.Create(ctx, []*model.Todo{todo}, clock)
	if err != nil {
		//	todo error handling
	}

	return nil
}
