package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type CommandDetail struct {
	PostID    int        `json:"postId" `
	UserID    int        `json:"userId"`
	Command   string     `json:"command"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}

func SeedCommandDetail(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Your picture is so great!",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    1,
		UserID:    3,
		Command:   "Gajadi",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
	db.Create(&CommandDetail{
		PostID:    5,
		UserID:    2,
		Command:   "Ngetest boss",
	})
}
