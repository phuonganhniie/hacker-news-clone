package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/phuonganhniie/hackernews/graph/generated"
	"github.com/phuonganhniie/hackernews/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link model.Link
	var user model.User
	link.Title = input.Title
	link.Address = input.Address
	user.Name = "test"
	link.User = &user
	// link.User.Name = "test"
	return &link, nil
}

func (r *mutationResolver) RefreshInputToken(ctx context.Context, input model.RefreshInputToken) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Link(ctx context.Context) ([]*model.Link, error) {
	var link []*model.Link
	dummyLink := model.Link{
		Title:   "Our Dummy Data",
		Address: "https://address.org",
		User:    &model.User{Name: "admin"},
	}
	links := append(link, &dummyLink)
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
