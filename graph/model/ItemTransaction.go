package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type ItemTransaction struct {
	TransactionID     int                  `json:"transactionId" gorm:"primaryKey"`
	UserSellerID      int                  `json:"userSellerId"`
	SellerGet         int                  `json:"sellerGet"`
	BuyerPay          int                  `json:"buyerPay"`
	UserBuyerID       int                  `json:"userBuyerId"`
	ItemID            int                  `json:"itemId"`
	Item              *Item                `json:"item" gorm:"foreignKey:ItemID"`
	TransactionType   *ItemTransactionType `json:"transactionType" gorm:"foreignKey:TransactionTypeID"`
	TransactionTypeID int                  `json:"transactionTypeId"`
	CreatedAt         time.Time            `json:"CreatedAt"`
	UpdatedAt         time.Time            `json:"UpdatedAt"`
	DeletedAt         *time.Time           `json:"DeletedAt"`
}

type ItemTransactionType struct {
	TypeID   int    `json:"typeId" gorm:"primaryKey"`
	TypeName string `json:"typeName"`
}

func SeedItemTransaction(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         230,
		BuyerPay:          300,
		UserBuyerID:       3,
		ItemID:            1,
		TransactionTypeID: 2,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         250,
		BuyerPay:          320,
		UserBuyerID:       3,
		ItemID:            1,
		TransactionTypeID: 2,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         250,
		BuyerPay:          320,
		UserBuyerID:       3,
		ItemID:            1,
		TransactionTypeID: 2,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            1,
		TransactionTypeID: 2,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      3,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       1,
		ItemID:            1,
		TransactionTypeID: 1,
	})

	db.Create(&ItemTransaction{
		UserSellerID:      1,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            1,
		TransactionTypeID: 3,
	})

	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            1,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            2,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            3,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            4,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            5,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            6,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            7,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            8,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            9,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            10,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            11,
		TransactionTypeID: 1,
	})
	db.Create(&ItemTransaction{
		UserSellerID:      2,
		SellerGet:         270,
		BuyerPay:          340,
		UserBuyerID:       3,
		ItemID:            12,
		TransactionTypeID: 1,
	})
}

func SeedItemTransactionType(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&ItemTransactionType{
		TypeName: "Sell",
	})
	db.Create(&ItemTransactionType{
		TypeName: "Purchased",
	})
	db.Create(&ItemTransactionType{
		TypeName: "Bid",
	})
}