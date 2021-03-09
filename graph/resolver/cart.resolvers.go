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

func (r *mutationResolver) InsertCart(ctx context.Context, userID int, gameID int) (*model.Cart, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	cart := &model.Cart{
		UserID: userID,
		GameID: gameID,
	}
	test := db.Find(cart)
	if test.RowsAffected > 0 {
		return nil, fmt.Errorf("data sudah ada di db")
	}

	db.Create(cart)

	return cart, nil
}

func (r *mutationResolver) DeleteCart(ctx context.Context, gameID int) (*model.Cart, error) {
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

	cart := &model.Cart{UserID: user.UserID, GameID: gameID}

	db.First(cart)
	db.Delete(cart)

	return cart, nil
}

func (r *queryResolver) GetUserCart(ctx context.Context) ([]*model.Cart, error) {
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

	var carts []*model.Cart

	db.Where("user_id", user.UserID).Preload(clause.Associations).Find(&carts)

	return carts, nil
}
