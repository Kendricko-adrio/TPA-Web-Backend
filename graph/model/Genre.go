package model

type Genre struct {
	GenreID   int    `json:"genreId" gorm:"primaryKey" `
	GenreName string  `json:"genreName"`
	Game      []*Game `json:"game" gorm:"many2many:game_genres"`
}
