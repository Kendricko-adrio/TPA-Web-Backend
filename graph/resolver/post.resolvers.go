package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"gorm.io/gorm/clause"
)

func (r *queryResolver) GetAllPost(ctx context.Context) ([]*model.Post, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	var posts []*model.Post
	db.Preload(clause.Associations).Find(&posts)

	return posts, nil
}

func (r *queryResolver) GetPost(ctx context.Context, postID int) (*model.Post, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	posts := model.Post{PostID: postID}

	db.Preload(clause.Associations).First(&posts)

	return &posts, nil
}
