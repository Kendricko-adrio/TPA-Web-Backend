package model

type GameSlideShow struct {
	SlideShowID  int    `json:"slideShowId" gorm:"primaryKey"`
	SlideShowURL string `json:"slideShowUrl"`
	GameID       int    `json:"gameId"`
}

