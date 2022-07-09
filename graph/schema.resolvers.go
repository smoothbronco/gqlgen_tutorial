package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/smoothbronco/gqlgen_tutorial/entity"
	"github.com/smoothbronco/gqlgen_tutorial/graph/generated"
	"github.com/smoothbronco/gqlgen_tutorial/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	record := entity.Todo{
		Text:   input.Text,
		UserID: input.UserID,
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewTodoFromEntity(&record)
	return res, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	record := entity.User{
		Name: input.Name,
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewUserFromEntity(&record)
	return res, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	id_number, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	var user entity.User
	if err := r.DB.Find(&user, id_number).Error; err != nil {
		return nil, err
	}
	return &model.User{
		ID:   fmt.Sprintf("%d", user.ID),
		Name: user.Name,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	var records []entity.Todo
	if err := r.DB.Where("user_id", obj.ID).Find(&records).Error; err != nil {
		return nil, err
	}

	todos := []*model.Todo{}
	for _, record := range records {
		todos = append(todos, model.NewTodoFromEntity(&record))
	}

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
