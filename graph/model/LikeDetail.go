package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type LikeDetail struct {
	PostID    int        `json:"postId" gorm:"primaryKey"`
	UserID    int        `json:"userId" gorm:"primaryKey"`
	Like      bool       `json:"like"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}

func SeedLikeDetail(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&LikeDetail{
		PostID:    1,
		UserID:    2,
		Like:      true,
	})
}