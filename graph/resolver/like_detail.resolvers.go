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

func (r *mutationResolver) InsertLike(ctx context.Context, postID int) (*model.LikeDetail, error) {
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

	like := model.LikeDetail{
		PostID: postID,
		UserID: user.UserID,
		Like:   true,
	}

	test := db.Where("user_id = ? AND post_id = ?", user.UserID, postID).Find(&model.LikeDetail{})

	if test.RowsAffected > 0 {
		return nil, fmt.Errorf("already inserted")
	}

	db.Create(&like)
	post := model.Post{PostID: postID}
	db.First(&post)
	post.TotalLike += 1
	db.Save(post)
	return &like, nil
}

func (r *mutationResolver) InsertDislike(ctx context.Context, postID int) (*model.LikeDetail, error) {
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

	test := db.Where("user_id = ? AND post_id = ?", user.UserID, postID).Find(&model.LikeDetail{})

	if test.RowsAffected > 0 {
		return nil, fmt.Errorf("already inserted")
	}

	like := model.LikeDetail{
		PostID: postID,
		UserID: user.UserID,
		Like:   false,
	}
	db.Create(&like)
	post := model.Post{PostID: postID}
	db.First(&post)
	post.TotalDislike += 1
	db.Save(post)
	return &like, nil
}
