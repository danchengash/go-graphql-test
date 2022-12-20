package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/Genialngash/graphql-go-test/db/store"
	"github.com/Genialngash/graphql-go-test/graph/model"
)

// AddTodo is the resolver for the addTodo field.
func (r *mutationResolver) AddTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	db := store.GetStoreFromContext(ctx)
	err := db.AddTodo(&input)
	if err != nil {
		return nil, err

	}
	return &model.Todo{
		ID:   "1",
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "Danche",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	db := store.GetStoreFromContext(ctx)

	return db.Todos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	db := store.GetStoreFromContext(ctx)
	err := db.AddTodo(&input)
	if err != nil {
		return nil, err

	}
	return &model.Todo{
		ID:   "1",
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "Danche",
		},
	}, nil
}
