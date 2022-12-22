package cli_usecase

import (
	"context"
	"database/sql"

	"github.com/yuuuutsk/gobase-backend/app/domain/model"
	"github.com/yuuuutsk/gobase-backend/pkg"

	"github.com/yuuuutsk/gobase-backend/app/domain/repository"
)

type TodoUseCase struct {
	db         *sql.DB
	repository repository.TodoRepository
}

func NewTodoUseCase(
	db *sql.DB,
	repo repository.TodoRepository,
) *TodoUseCase {
	return &TodoUseCase{
		db:         db,
		repository: repo,
	}
}

type TodoGetInput struct {
	ID string
}

func (uc *TodoUseCase) Get(ctx context.Context, input *TodoGetInput, clock pkg.Clock) (*model.Todo, error) {
	// TODO: implement
	return nil, nil
}
