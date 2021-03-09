package model

import "github.com/kendricko-adrio/gqlgen-todos/database"

type PaymentType struct {
	TypeID   int    `json:"typeId" gorm:"primaryKey"`
	TypeName string `json:"typeName"`
}

func SeedPaymentType(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&PaymentType{
		TypeName: "Wallet",
	})
	db.Create(&PaymentType{
		TypeName: "Visa",
	})
}