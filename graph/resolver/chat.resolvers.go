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

func (r *mutationResolver) AddMessage(ctx context.Context, receiverID int, text string) (*model.Chat, error) {
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

	db.First(&user)
	friend := model.User{UserID: receiverID}
	db.First(&friend)

	msg := model.Chat{
		Text:       text,
		ReceiverID: friend.UserID,
		SenderID:   user.UserID,
	}

	db.Create(&msg)

	if friendSocket, ok := r.ChatSockets[friend.UserID]; ok {
		friendSocket <- &msg
	}

	return &msg, nil

}

func (r *queryResolver) GetChat(ctx context.Context, receiverID int) ([]*model.Chat, error) {
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

	var chat []*model.Chat

	db.
		Where("sender_id = ? OR receiver_id = ?", receiverID, receiverID).
		Where("sender_id = ? OR receiver_id = ?", user.UserID, user.UserID).
		Order("created_at DESC").
		Preload(clause.Associations).
		Find(&chat)
	return chat,nil
}

func (r *subscriptionResolver) NewMessageAdded(ctx context.Context) (<-chan *model.Chat, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	socket := make(chan *model.Chat, 1)

	r.ChatSockets[user.UserID] = socket

	go func() {
		<-ctx.Done()
		delete(r.ChatSockets, user.UserID)
	}()

	return socket, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
