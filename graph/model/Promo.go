package model

import (
	"gorm.io/gorm"
	"time"
)

type Promo struct {
	PromoID       int        `json:"promoId" gorm:"primaryKey"`
	PromoDiscount int        `json:"promoDiscount"`
	PromoDuration int        `json:"promoDuration"`
	GameID        int        `json:"gameId"`
	Game          *Game      `json:"game" gorm:"foreignKey:GameID"`
	CreatedAt     time.Time  `json:"CreatedAt"`
	UpdatedAt     time.Time  `json:"UpdatedAt"`
	DeletedAt 	gorm.DeletedAt
}