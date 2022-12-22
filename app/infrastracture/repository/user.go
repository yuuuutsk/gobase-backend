package repository

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/yuuuutsk/gobase-backend/app/domain"
	"github.com/yuuuutsk/gobase-backend/app/domain/model"
	"github.com/yuuuutsk/gobase-backend/app/domain/repository"
	"github.com/yuuuutsk/gobase-backend/app/infrastracture/dao"
	"github.com/yuuuutsk/gobase-backend/pkg"
	"github.com/yuuuutsk/gobase-backend/pkg/logger"
)

type userRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewUserRepository(db *sql.DB, logger logger.Logger) repository.UserRepository {
	return &userRepository{db: db, logger: logger}
}

func (repo *userRepository) Create(ctx context.Context, users []*model.User, clock pkg.Clock) error {
	var userDtos dao.UserSlice
	for _, user := range users {
		dto := &dao.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		userDtos = append(userDtos, dto)
	}

	if err := userDtos.InsertAll(ctx, repo.db, boil.Infer()); err != nil {
		repo.logger.Warnf("user dto.Insert")
	}

	return nil
}

func (repo *userRepository) All(ctx context.Context) ([]*model.User, error) {
	mods := []qm.QueryMod{
		//dao.TodoWhere.Done.EQ(false),
		//dao.TodoWhere.Text.EQ(""),
	}

	dtos, err := dao.Users(
		mods...,
	).All(ctx, repo.db)

	if err != nil {
		return nil, err
	}

	result := make([]*model.User, 0, len(dtos))
	for i, dto := range dtos {
		result[i] = model.RestoreUser(
			domain.UserID(dto.ID),
			dto.FirstName,
			dto.LastName)
	}
	return result, nil
}
