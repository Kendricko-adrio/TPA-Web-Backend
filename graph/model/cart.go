package model

type Cart struct {
	User   *User `json:"user" gorm:"foreignKey:UserID"`
	Game   *Game `json:"game" gorm:"foreignKey:GameID"`
	UserID int   `json:"userId" gorm:"primaryKey"`
	GameID int   `json:"gameId" gorm:"primaryKey"`
}