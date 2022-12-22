package di

import (
	"github.com/yuuuutsk/gobase-backend/app/cli_usecase"
	"github.com/yuuuutsk/gobase-backend/app/usecase"
)

type CLIUseCases struct {
	TodoUseCase *cli_usecase.TodoUseCase
}

func NewCLIUseCases(
	TodoUseCase *cli_usecase.TodoUseCase,
) *CLIUseCases {
	return &CLIUseCases{
		TodoUseCase: TodoUseCase,
	}
}

type UseCases struct {
	TodoUseCase *usecase.TodoUseCase
	UserUseCase *usecase.UserUseCase
}

func NewUseCases(
	TodoUseCase *usecase.TodoUseCase,
	UserUseCase *usecase.UserUseCase,
) *UseCases {
	return &UseCases{
		TodoUseCase: TodoUseCase,
		UserUseCase: UserUseCase,
	}
}
