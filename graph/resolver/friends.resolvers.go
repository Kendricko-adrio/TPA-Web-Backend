package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
)

func (r *mutationResolver) Request(ctx context.Context, userID1 *int, userID2 *int) (*model.FriendsDetail, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	friend := model.FriendsDetail{
		User1id:        *userID1,
		User2id:        *userID2,
		FriendStatusID: 1,
	}

	db.Create(&friend)

	return &friend, nil
}
