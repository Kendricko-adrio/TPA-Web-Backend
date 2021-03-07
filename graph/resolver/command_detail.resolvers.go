package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
)

func (r *mutationResolver) InsertCommand(ctx context.Context, postID int, command string) (*model.CommandDetail, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	insertCommand := model.CommandDetail{
		PostID:    postID,
		UserID:    user.UserID,
		Command:   command,
	}

	db.Create(&insertCommand)

	return &insertCommand, nil
}

func (r *queryResolver) GetCommandPaginate(ctx context.Context, postID int, page int) ([]*model.CommandDetail, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	var commands []*model.CommandDetail

	db.Where("post_id = ? ", postID).Order("created_at desc").Limit(10).Offset((page - 1) * 10).Find(&commands)
	return commands, nil
}
