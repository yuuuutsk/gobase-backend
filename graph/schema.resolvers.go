package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/yuuuutsk/gobase-backend/app/usecase"
	"github.com/yuuuutsk/gobase-backend/cmd/di"
	"github.com/yuuuutsk/gobase-backend/pkg"

	"github.com/yuuuutsk/gobase-backend/graph/generated"
	"github.com/yuuuutsk/gobase-backend/graph/gqlmodel"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input gqlmodel.NewTodo) (*gqlmodel.Todo, error) {

	in := &usecase.TodoCreateInput{
		Text: input.Text,
		//Done: input.,
	}
	clock := pkg.NewClock()
	err := usecases.TodoUseCase.Create(ctx, in, clock)
	return nil, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*gqlmodel.Todo, error) {

	in := &usecase.TodoGetInput{
		Text: "test",
	}
	clock := pkg.NewClock()
	res, err := usecases.TodoUseCase.Get(ctx, in, clock)

	if err != nil {
		return nil, err
	}

	out := []*gqlmodel.Todo{}
	if res.Todo != nil {
		out = append(out, &gqlmodel.Todo{
			Text: res.Todo.Text,
			Done: res.Todo.Done,
			ID:   fmt.Sprintf("%d", uint(res.Todo.ID)),
		})
	}

	return out, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, firstName string,
	lastName string) (*gqlmodel.User, error) {
	//TODO implement me
	// usecaseを呼び出す
	panic("implement me")
}

type queryResolver struct{ *Resolver }

var usecases *di.UseCases

func SetUsecases(u *di.UseCases) {
	usecases = u
}
