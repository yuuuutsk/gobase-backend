package usecase

import (
	"database/sql"

	"github.com/yuuuutsk/gobase-backend/pkg/logger"

	"github.com/yuuuutsk/gobase-backend/app/domain/repository"
)

type TodoUseCase struct {
	db         *sql.DB
	repository repository.TodoRepository
	logger     logger.Logger
}

func NewTagUseCase(
	db *sql.DB,
	repo repository.TodoRepository,
	logger logger.Logger,
) *TodoUseCase {
	return &TodoUseCase{
		db:         db,
		repository: repo,
		logger:     logger,
	}
}
