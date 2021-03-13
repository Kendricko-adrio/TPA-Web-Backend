package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type TransactionHeader struct {
	TransactionID     int                  `json:"transactionId" gorm:"primaryKey"`
	User              *User                `json:"User" gorm:"foreignKey:UserID"`
	UserID            int                  `json:"UserId" `
	Recepient         *User                `json:"Recepient" gorm:"foreignKey:RecepientID"`
	RecepientID       int                  `json:"RecepientId"`
	TotalPayment      int                  `json:"TotalPayment"`
	PaymentType       *PaymentType         `json:"PaymentType" gorm:"foreignKey:PaymentTypeID"`
	PaymentTypeID     int                  `json:"PaymentTypeId"`
	TransactionDetail []*TransactionDetail `json:"transactionDetail" gorm:"foreignKey:TransactionID"`
	CreatedAt         time.Time            `json:"CreatedAt"`
	UpdatedAt         time.Time            `json:"UpdatedAt"`
	DeletedAt         *time.Time           `json:"DeletedAt"`
}

func SeedTransaction(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&TransactionHeader{
		UserID:            5,
		RecepientID:       5,
		TotalPayment:      100,
		PaymentTypeID:     1,
		TransactionDetail: []*TransactionDetail{
			{
				GameID:        1,
			},
			{
				GameID:        2,
			},
		},
	})
	db.Create(&TransactionHeader{
		UserID:            5,
		RecepientID:       5,
		TotalPayment:      100,
		PaymentTypeID:     1,
		TransactionDetail: []*TransactionDetail{
			{
				GameID:        2,
			},
			{
				GameID:        3,
			},
		},
	})
	db.Create(&TransactionHeader{
		UserID:            5,
		RecepientID:       5,
		TotalPayment:      100,
		PaymentTypeID:     1,
		TransactionDetail: []*TransactionDetail{
			{
				GameID:        3,
			},
			{
				GameID:        1,
			},
		},
	})
	db.Create(&TransactionHeader{
		UserID:            5,
		RecepientID:       5,
		TotalPayment:      100,
		PaymentTypeID:     1,
		TransactionDetail: []*TransactionDetail{
			{
				GameID:        1,
			},
			{
				GameID:        2,
			},
		},
	})
	db.Create(&TransactionHeader{
		UserID:            5,
		RecepientID:       5,
		TotalPayment:      100,
		PaymentTypeID:     1,
		TransactionDetail: []*TransactionDetail{
			{
				GameID:        5,
			},
			{
				GameID:        6,
			},
		},
	})
}
