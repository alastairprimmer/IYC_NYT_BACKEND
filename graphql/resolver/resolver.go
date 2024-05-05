package resolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/alastairprimmer/IYC_NYT_BACKEND/graphql/generated"
	"github.com/alastairprimmer/IYC_NYT_BACKEND/graphql/model"
)

type Resolver struct{}

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, name string, email string) (*model.User, error) {
	panic("not implemented")
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*string, error) {
	panic("not implemented")
}

// AddStrand is the resolver for the addStrand field.
func (r *mutationResolver) AddStrand(ctx context.Context, name string, theme string, spanogram string, words []string) (*model.Strand, error) {
	panic("not implemented")
}

// DeleteStrand is the resolver for the deleteStrand field.
func (r *mutationResolver) DeleteStrand(ctx context.Context, id string) (*string, error) {
	panic("not implemented")
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
