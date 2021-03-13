package model

import "time"

type TransactionDetail struct {
	TransactionID int `json:"TransactionId" gorm:"primaryKey"`
	GameID        int `json:"GameId" gorm:"primaryKey"`
	Game          *Game `json:"Game" gorm:"foreignKey:GameID"`
	CreatedAt     time.Time  `json:"CreatedAt"`
	UpdatedAt     time.Time  `json:"UpdatedAt"`
	DeletedAt     *time.Time `json:"DeletedAt"`
}