package model

import "github.com/kendricko-adrio/gqlgen-todos/database"

type TopUpWallet struct {
	WalletID    int        `json:"walletId" gorm:"primaryKey"`
	WalletCode  string     `json:"walletCode"`
	WalletMoney int        `json:"walletMoney"`
	TopUpTypeID int        `json:"topUpTypeId"`
	TopUpType   *TopUpType `json:"topUpType" gorm:"foreignKey:TopUpTypeID"`
}

type TopUpType struct {
	TopUpTypeID     int `json:"topUpTypeId" gorm:"primaryKey"`
	TopUpTypeIDName string `json:"topUpTypeIdName"`
}

func SeedTopUpWallet(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&TopUpWallet{
		WalletCode:  "1234-5678-9123",
		WalletMoney: 50000,
		TopUpTypeID: 1,
	})
}

func SeedTopUpType(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&TopUpType{
		TopUpTypeIDName: "Available",
	})
	db.Create(&TopUpType{
		TopUpTypeIDName: "Already Used",
	})
}
