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

func (r *mutationResolver) InsertWallet(ctx context.Context, walletCode string) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.First(&user)
	var code model.TopUpWallet
	test := db.Where("wallet_code LIKE ? AND top_up_type_id = 1", walletCode).First(&code)

	if test.RowsAffected == 0 {
		return nil, fmt.Errorf("No Wallet")
	}
	code.TopUpTypeID = 2
	user.Money += code.WalletMoney
	db.Save(&user)
	db.Save(&code)
	return user, nil
}
