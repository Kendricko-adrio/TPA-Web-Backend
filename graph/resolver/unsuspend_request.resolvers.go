package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) DeleteRequest(ctx context.Context, userID int) (*model.UnsuspendRequest, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AcceptRequest(ctx context.Context, requestID int) (*model.UnsuspendRequest, error) {
	//panic(fmt.Errorf("not implemented"))
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var request = &model.UnsuspendRequest{RequestID: requestID}
	db.First(request)

	var user = &model.User{UserID: request.SuspendedUserID}
	db.First(user)

	user.StatusID = 2
	db.Save(user)
	db.Delete(request)
	return request, nil
}

func (r *mutationResolver) RejectRequest(ctx context.Context, requestID int) (*model.UnsuspendRequest, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var request = &model.UnsuspendRequest{RequestID: requestID}
	db.First(request)
	db.Delete(request)
	return request, nil
}

func (r *queryResolver) GetUnsuspendByUserID(ctx context.Context, userID int) ([]*model.UnsuspendRequest, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var user []*model.UnsuspendRequest

	db.Where("suspended_user_id = ?", userID).Preload(clause.Associations).First(&user)

	return user, nil
}
