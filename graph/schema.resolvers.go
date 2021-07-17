package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/ashwinp15/audio-directory/graph/generated"
	"github.com/ashwinp15/audio-directory/graph/model"
)

func (r *mutationResolver) CreateNooble(ctx context.Context, input model.NewNooble) (*model.Nooble, error) {
	id := strconv.Itoa(rand.Intn(1000000))
	nooble := &model.Nooble{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Audio:       input.File.Filename,
	}
	fmt.Printf("%v\n", nooble)
	r.noobles = append(r.noobles, nooble)
	return nooble, nil
}

func (r *queryResolver) Noobles(ctx context.Context) ([]*model.Nooble, error) {
	return r.noobles, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
