package model

type TransactionDetail struct {
	TransactionID int `json:"TransactionId" gorm:"primaryKey"`
	GameID        int `json:"GameId" gorm:"primaryKey"`
	Game          *Game `json:"Game" gorm:"foreignKey:GameID"`
}