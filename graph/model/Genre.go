package model

type Genre struct {
	GenreID   int    `json:"genreId" gorm:"primaryKey"`
	GenreName string `json:"genreName"`
}
