package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ashwinp15/audio-directory/graph/generated"
	"github.com/ashwinp15/audio-directory/graph/model"
)

func (r *mutationResolver) CreateNooble(ctx context.Context, input model.NewNooble) (*string, error) {
	r.nooble = &model.Nooble{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Audio:       input.File.Filename,
	}
	user, err := model.GetUserByEmail(input.Creator)
	if err != nil {
		return nil, err
	}
	r.nooble.Creator = user
	fmt.Println("mutation resolved successfully")
	r.PutNooble(input.File)
	fmt.Println("upload successful")
	return &r.nooble.ID, nil
}

func (r *mutationResolver) UpdateNooble(ctx context.Context, id string, input model.UpdateNooble) (*string, error) {
	r.nooble = &model.Nooble{ID: id}
	return r.UpdateDetails(&input)
}

func (r *mutationResolver) DeleteNooble(ctx context.Context, id string) (*string, error) {
	r.nooble = &model.Nooble{ID: id}
	return r.Delete()
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewCreator) (*model.Creator, error) {
	user := model.Creator{
		Name:  input.Name,
		Email: input.Email,
	}
	if err := user.Create(input.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *queryResolver) Noobles(ctx context.Context) ([]*model.Nooble, error) {
	return r.ReadAllNoobles()
}

func (r *queryResolver) Nooble(ctx context.Context, id string) (*model.Nooble, error) {
	return r.ReadSingleNooble(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
