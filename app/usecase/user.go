package usecase

import (
	"database/sql"

	"github.com/yuuuutsk/gobase-backend/pkg/logger"

	"github.com/yuuuutsk/gobase-backend/app/domain/repository"
)

type UserUseCase struct {
	db         *sql.DB
	repository repository.UserRepository
	logger     logger.Logger
}

func NewUserUseCase(
	db *sql.DB,
	repo repository.UserRepository,
	logger logger.Logger,
) *UserUseCase {
	return &UserUseCase{
		db:         db,
		repository: repo,
		logger:     logger,
	}
}
