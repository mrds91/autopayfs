package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"autopay/graph/model"
	"autopay/graph/service"
	"context"
	"fmt"
)

// Autopay is the resolver for the autopay field.
func (r *mutationResolver) Autopay(ctx context.Context, input model.InputRequestInfo) (*model.ResponseInfo, error) {
	var res model.ResponseInfo
	res.Status = "suss"
	return service.FechCustomers(input), nil
}

// Res is the resolver for the res field.
func (r *queryResolver) Res(ctx context.Context) (*model.ResponseInfo, error) {
	panic(fmt.Errorf("not implemented: Res - res"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }