package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
)

func (r *queryResolver) GetAllGenre(ctx context.Context) ([]*model.Genre, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var genre []*model.Genre
	db.Find(&genre)
	return genre, nil
}
