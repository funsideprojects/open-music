package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fsp/open-music/identity/graph/generated"
	"fsp/open-music/identity/graph/model"
	"log"
)

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	log.Print(ctx)
	return &model.User{
		ID: "1",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
