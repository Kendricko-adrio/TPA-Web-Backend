package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/mailjet"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
)

func (r *mutationResolver) InsertTransaction(ctx context.Context, userID int, recepientID int, total int, paymentTypeID int, transactionDetail []int) (*model.TransactionHeader, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	for _, element := range transactionDetail {
		db.Exec("DELETE FROM carts WHERE user_id = ? AND game_id = ?", userID, element).Debug()
	}

	transaction := model.TransactionHeader{
		UserID:        userID,
		RecepientID:   recepientID,
		TotalPayment:  total,
		PaymentTypeID: paymentTypeID,
	}

	db.Create(&transaction)

	for _, element := range transactionDetail {
		db.Exec("INSERT INTO transaction_details(transaction_id, game_id) VALUES (?,?)", transaction.TransactionID, element).Debug()
		db.Exec("INSERT INTO users_games(game_id, user_user_id) VALUES (?,?)", element, recepientID).Debug()
	}

	if paymentTypeID == 1 {
		user := model.User{UserID: userID}
		db.First(&user)
		user.Money = user.Money - total
		db.Save(&user)
	}

	var detail []*model.Game
	db.Joins("JOIN transaction_details ON transaction_details.game_id = games.id").Where("transaction_details.transaction_id = ?", transaction.TransactionID).Find(&detail)

	userCurr := model.User{UserID: userID}

	db.First(&userCurr)

	mailjet.Transaction(userCurr.Email, detail)

	return &transaction, nil
}

func (r *mutationResolver) GetAllTransaction(ctx context.Context) ([]*model.TransactionHeader, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllPaymentType(ctx context.Context) ([]*model.PaymentType, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	var types []*model.PaymentType

	db.Find(&types)

	return types, nil
}
