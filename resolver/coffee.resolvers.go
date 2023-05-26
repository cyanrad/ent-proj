package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"main/ent"
	"main/resolver/generated"
)

func (r *mutationResolver) CreateCoffee(ctx context.Context, input ent.CreateCoffeeInput) (*ent.Coffee, error) {
	return r.client.Coffee.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteCoffee(ctx context.Context, id int) (*int, error) {
	return &id, r.client.Coffee.DeleteOneID(id).Exec(ctx)
}

func (r *mutationResolver) UpdateCoffee(ctx context.Context, id int, input ent.UpdateCoffeeInput) (*ent.Coffee, error) {
	return r.client.Coffee.UpdateOneID(id).SetInput(input).Save(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
