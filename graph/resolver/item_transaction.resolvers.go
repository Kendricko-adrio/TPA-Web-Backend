package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/mailjet"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) BuyItem(ctx context.Context, itemID int, price int) (*model.ItemTransaction, error) {
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

	var itemCheck model.ItemTransaction

	check := db.Where("buyer_pay <= ? AND item_id = ? AND transaction_type_id = 1 AND user_seller_id != ?",
		price, itemID, user.UserID).Order("buyer_pay DESC").First(&itemCheck)
	transactionStatus := 3
	sellerId := 1
	if check.RowsAffected > 0 {
		itemCheck.TransactionTypeID = 2
		//db.Exec("INSERT INTO users_item(item_items_id, user_user_id) VALUES(?,?)", itemCheck.ItemID, user.UserID)
		db.Exec("INSERT INTO my_items(item_id, user_id) VALUES(?,?)", itemCheck.ItemID, user.UserID)
		//db.Exec("DELETE FROM users_item WHERE item_items_id = ?"+
		//	" AND user_user_id = ?", itemCheck.ItemID, itemCheck.UserSellerID)
		db.Exec("DELETE FROM my_items WHERE my_item_id IN("+
			"SELECT my_item_id "+
			"FROM my_items WHERE item_id = ? AND user_id = ? LIMIT 1 "+
			")", itemCheck.ItemID, itemCheck.UserSellerID)
		transactionStatus = 2
		sellerId = itemCheck.UserSellerID
		userSeller := model.User{UserID: itemCheck.UserSellerID}
		db.First(&user)
		db.First(&userSeller)

		user.Money -= itemCheck.BuyerPay
		db.Save(&user)
		userSeller.Money += itemCheck.SellerGet
		db.Save(&userSeller)
		db.Save(&itemCheck)
		idItem := model.Item{ItemsID: itemCheck.ItemID}
		db.First(&idItem)
		mailjet.OnBuyItem(user.Email, idItem.ItemName, itemCheck.BuyerPay)
	}

	itemt := &model.ItemTransaction{
		UserSellerID:      sellerId,
		SellerGet:         price,
		BuyerPay:          price,
		UserBuyerID:       user.UserID,
		ItemID:            itemID,
		TransactionTypeID: transactionStatus,
	}

	db.Create(itemt)

	return itemt, nil
}

func (r *mutationResolver) SellItem(ctx context.Context, itemID int, price int) (*model.ItemTransaction, error) {
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

	var itemCheck model.ItemTransaction

	check := db.Where("buyer_pay >= ? AND item_id = ? AND transaction_type_id = 3 AND user_seller_id != ?",
		price, itemID, user.UserID).Order("buyer_pay DESC").First(&itemCheck)
	transactionStatus := 1
	buyerId := 1
	if check.RowsAffected > 0 {
		itemCheck.TransactionTypeID = 2
		//db.Exec("INSERT INTO users_item(item_items_id, user_user_id) VALUES(?,?)", itemCheck.ItemID, user.UserID)
		db.Exec("INSERT INTO my_items(item_id, user_id) VALUES(?,?)", itemCheck.ItemID, itemCheck.UserBuyerID)
		//db.Exec("DELETE FROM users_item WHERE item_items_id = ?"+
		//	" AND user_user_id = ?", itemCheck.ItemID, itemCheck.UserSellerID)
		db.Exec("DELETE FROM my_items WHERE my_item_id IN("+
			"SELECT my_item_id "+
			"FROM my_items WHERE item_id = ? AND user_id = ? LIMIT 1 "+
			")", itemCheck.ItemID, user.UserID)
		transactionStatus = 2
		buyerId = itemCheck.UserSellerID
		userBuyer := model.User{UserID: itemCheck.UserBuyerID}
		db.First(&user)
		db.First(&userBuyer)

		user.Money += itemCheck.BuyerPay
		db.Save(&user)
		userBuyer.Money -= itemCheck.SellerGet
		db.Save(&userBuyer)
		db.Save(&itemCheck)
		idItem := model.Item{ItemsID: itemCheck.ItemID}
		db.First(&idItem)
		mailjet.OnSellItem(user.Email, idItem.ItemName, itemCheck.BuyerPay)
	}

	itemt := &model.ItemTransaction{
		UserSellerID:      user.UserID,
		SellerGet:         price,
		BuyerPay:          price,
		UserBuyerID:       buyerId,
		ItemID:            itemID,
		TransactionTypeID: transactionStatus,
	}

	db.Create(itemt)

	return itemt, nil
}

func (r *queryResolver) GetTransactionByItem(ctx context.Context, itemID int) ([]*model.ItemTransaction, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	var transac []*model.ItemTransaction

	db.Where("item_id = ? AND transaction_type_id = 2", itemID).Preload(clause.Associations).Find(&transac)

	return transac, nil
}

func (r *queryResolver) GetRequestTransaction(ctx context.Context, page int) ([]*model.Item, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	//var transaction []*model.ItemTransaction
	var items []*model.Item

	db.Raw("SELECT i.* "+
		"FROM items AS i "+
		"JOIN item_transactions as it "+
		"ON it.item_id = i.items_id "+
		"WHERE (it.updated_at BETWEEN (now() - '1 month'::interval) AND now())"+
		"GROUP BY i.items_id "+
		"ORDER BY COUNT(it.transaction_id) DESC LIMIT 10 OFFSET ?", (page-1)*10).Preload(clause.Associations).Find(&items)

	return items, nil
}

func (r *queryResolver) GetSell(ctx context.Context, itemID int) ([]*model.ItemTransaction, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	var it []*model.ItemTransaction
	db.Raw("select * "+
		"from item_transactions AS it "+
		"WHERE it.transaction_type_id = 1 AND it.item_id = ? "+
		"ORDER BY buyer_pay DESC "+
		"LIMIT 10", itemID).Find(&it)

	return it, nil
}

func (r *queryResolver) GetBuy(ctx context.Context, itemID int) ([]*model.ItemTransaction, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	var it []*model.ItemTransaction
	db.Raw("select * "+
		"from item_transactions AS it "+
		"WHERE it.transaction_type_id = 3 AND it.item_id = ? "+
		"ORDER BY buyer_pay ASC "+
		"LIMIT 10", itemID).Find(&it)

	return it, nil
}

func (r *queryResolver) GetMyBuy(ctx context.Context, itemID int) ([]*model.ItemTransaction, error) {
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

	var myItem []*model.ItemTransaction

	db.Where("user_buyer_id = ? AND transaction_type_id = 3 AND item_id = ?", user.UserID, itemID).Preload(clause.Associations).Find(&myItem)

	return myItem, nil
}

func (r *queryResolver) GetMySell(ctx context.Context, itemID int) ([]*model.ItemTransaction, error) {
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

	var myItem []*model.ItemTransaction

	db.Where("user_seller_id = ? AND transaction_type_id = 1 AND item_id = ?", user.UserID, itemID).Preload(clause.Associations).Find(&myItem)

	return myItem, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) BuyItem(ctx context.Context, transactionID int, price int) ([]*model.ItemTransaction, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) SellItem(ctx context.Context, transactionID int, price int) ([]*model.ItemTransaction, error) {
	panic(fmt.Errorf("not implemented"))
}
