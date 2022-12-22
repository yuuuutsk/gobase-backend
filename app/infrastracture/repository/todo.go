package repository

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/yuuuutsk/gobase-backend/app/domain"
	"github.com/yuuuutsk/gobase-backend/pkg/logger"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/yuuuutsk/gobase-backend/app/domain/model"
	"github.com/yuuuutsk/gobase-backend/app/domain/repository"
	"github.com/yuuuutsk/gobase-backend/app/infrastracture/dao"
	"github.com/yuuuutsk/gobase-backend/pkg"
)

type tagRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewTagRepository(db *sql.DB, logger logger.Logger) repository.TodoRepository {
	return &tagRepository{db: db, logger: logger}
}

func (repo *tagRepository) Create(ctx context.Context, todos []*model.Todo, clock pkg.Clock) error {
	var todoDtos dao.TodoSlice
	for _, todo := range todos {
		dto := &dao.Todo{
			Text:   todo.Text,
			Done:   todo.Done,
			UserID: 0,
		}
		todoDtos = append(todoDtos, dto)
	}

	if err := todoDtos.InsertAll(ctx, repo.db, boil.Infer()); err != nil {
		repo.logger.Warnf("todo dto.Insert")
	}

	return nil
}
func (repo *tagRepository) GetByText(ctx context.Context, text string, clock pkg.Clock) (*model.Todo, error) {
	mods := []qm.QueryMod{
		dao.TodoWhere.Text.EQ(text),
	}

	dto, err := dao.Todos(
		mods...,
	).One(ctx, repo.db)

	if err != nil {
		return nil, err
	}

	a := model.RestoreTodo(
		domain.TodoID(dto.ID),
		dto.Text,
		dto.Done,
	)

	return a, nil
}

func (repo *tagRepository) All(ctx context.Context) ([]*model.Todo, error) {
	mods := []qm.QueryMod{
		//dao.TodoWhere.Done.EQ(false),
		//dao.TodoWhere.Text.EQ(""),
	}

	dtos, err := dao.Todos(
		mods...,
	).All(ctx, repo.db)

	if err != nil {
		return nil, err
	}

	result := []*model.Todo{}
	for _, dto := range dtos {
		a := model.RestoreTodo(
			domain.TodoID(dto.ID),
			dto.Text,
			dto.Done,
		)
		result = append(result, a)
	}

	return result, nil
}
