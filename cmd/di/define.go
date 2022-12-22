//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"

	"github.com/yuuuutsk/gobase-backend/pkg/logger"

	"github.com/yuuuutsk/gobase-backend/app/usecase"

	"github.com/yuuuutsk/gobase-backend/cmd"

	"github.com/google/wire"
	"github.com/yuuuutsk/gobase-backend/app/cli_usecase"
	"github.com/yuuuutsk/gobase-backend/app/infrastracture/repository"
	"github.com/yuuuutsk/gobase-backend/pkg"
)

func InitCLIUseCases(
	db *sql.DB,
	config *cmd.Config,
	twClient *cmd.TwitterClients,
	logger logger.Logger,
	clock pkg.Clock,
) *CLIUseCases {
	wire.Build(
		NewCLIUseCases,
		cli_usecase.NewTodoUseCase,
		repository.NewTagRepository,
	)
	return nil
}

func InitUseCases(
	db *sql.DB,
	config *cmd.Config,
	twClient *cmd.TwitterClients,
	logger logger.Logger,
	clock pkg.Clock,
) *UseCases {
	wire.Build(
		NewUseCases,
		usecase.NewTagUseCase,
		usecase.NewUserUseCase,
		repository.NewTagRepository,
		repository.NewUserRepository,
	)
	return nil
}
