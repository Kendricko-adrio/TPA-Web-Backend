package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/generated"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) Request(ctx context.Context, userID2 *int) (*model.FriendsDetail, error) {
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
	friend := model.FriendsDetail{
		User1id:        user.UserID,
		User2id:        *userID2,
		FriendStatusID: 1,
	}

	db.Create(&friend)

	return &friend, nil
}

func (r *mutationResolver) AcceptFriendRequest(ctx context.Context, userID int) (*model.FriendsDetail, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	close, err := db.DB()
	defer close.Close()

	var request model.FriendsDetail
	user1 := model.User{UserID: userID}
	user2 := model.User{UserID: user.UserID}
	db.First(&user1)
	db.First(&user2)
	db.Where("user1id = ? AND user2id = ?", userID, user.UserID).First(&request)

	user1.Friends = append(user1.Friends, &model.User{UserID: user.UserID})
	//fmt.Println(user1.Friends)

	//return nil, nil
	user2.Friends = append(user2.Friends, &model.User{UserID: userID})
	request.FriendStatusID = 2
	db.Save(&user1)
	db.Save(&user2)
	db.Save(&request)

	return &request, err
}

func (r *mutationResolver) DeclineFriendRequest(ctx context.Context, userID int) (*model.FriendsDetail, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	close, err := db.DB()
	defer close.Close()
	var request model.FriendsDetail
	db.Where("user1id = ? AND user2id = ?", userID, user.UserID).First(&request)
	request.FriendStatusID = 3

	db.Save(&request)
	return &request, nil
}

func (r *queryResolver) GetTotalRequestFriend(ctx context.Context) (int, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return 0, fmt.Errorf("access denied")
	}

	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	close, err := db.DB()
	defer close.Close()
	var request []*model.FriendsDetail

	test := db.Where("user2id = ? AND friend_status_id = 1", user.UserID).Debug().Find(&request)
	return int(test.RowsAffected), nil
}

func (r *queryResolver) GetAllRequestFriend(ctx context.Context) ([]*model.FriendsDetail, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	close, err := db.DB()
	defer close.Close()
	var request []*model.FriendsDetail

	db.Where("user2id = ? AND friend_status_id = 1", user.UserID).Preload(clause.Associations).Debug().Find(&request)
	return request, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
