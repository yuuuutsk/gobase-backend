package usecase

import (
	"context"

	"github.com/yuuuutsk/gobase-backend/app/domain/models"
	"github.com/yuuuutsk/gobase-backend/pkg"
)

type TodoGetInput struct {
	Text string
}

type TodoGetOutput struct {
	Todo *models.Todo
}

func (uc *TodoUseCase) Get(ctx context.Context, input *TodoGetInput, clock pkg.Clock) (*TodoGetOutput, error) {

	todo, err := uc.repository.GetByText(ctx, input.Text, clock)
	if err != nil {
		//	todo error handling
	}

	output := &TodoGetOutput{
		Todo: todo,
	}
	return output, nil
}
