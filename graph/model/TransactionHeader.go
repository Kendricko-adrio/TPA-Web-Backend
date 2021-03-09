package model

import "time"

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
