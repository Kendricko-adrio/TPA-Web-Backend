package model

import (
	"fmt"
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type Game struct {
	ID          int        `json:"ID" gorm:"primaryKey"`
	Name        string     `json:"Name"`
	Description string     `json:"Description"`
	Price       int        `json:"Price"`
	Rating      int        `json:"Rating"`
	ImageBanner       string     `json:"imageBanner"`
	Image       string     `json:"Image"`
	Tag               string     `json:"tag"`
	SystemRequirement string     `json:"systemRequirement"`
	Users       []*User    `json:"users" gorm:"many2many:users_games"`
	Genre             *Genre     `json:"genre" gorm:"foreignKey:GenreID"`
	GenreID           int        `json:"genreId"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	UpdatedAt   time.Time  `json:"UpdatedAt"`
	DeletedAt   *time.Time `json:"DeletedAt"`
}

func GetAllGames() ([]*Game, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var games []*Game

	db.Find(&games)
	fmt.Print(games)
	return games, err
}
