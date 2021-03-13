package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) InsertReview(ctx context.Context, review string, helpful bool, gameID int) (*model.Post, error) {
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

	post := &model.Post{
		UserID:          user.UserID,
		GameID:          gameID,
		PostTypeID:      2,
		PostHelpful:     helpful,
		PostDescription: review,
		TotalLike:       0,
		TotalDislike:    0,
	}
	db.Create(post)

	return post, nil
}

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

func (r *queryResolver) GetReviewByGameRecent(ctx context.Context, gameID int) ([]*model.Post, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var posts []*model.Post
	db.Where("game_id = ? AND post_type_id = 2", gameID).Order("created_at DESC").Find(&posts)

	return posts, nil
}

func (r *queryResolver) GetReviewByGameUpvoted(ctx context.Context, gameID int) ([]*model.Post, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var posts []*model.Post
	db.Where("game_id = ? AND post_type_id = 2 AND (updated_at BETWEEN (now() - '1 month'::interval) AND now())", gameID).Order("total_like DESC").Find(&posts)

	return posts, nil
}
