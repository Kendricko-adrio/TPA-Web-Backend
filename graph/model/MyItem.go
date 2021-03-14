package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type MyItem struct {
	MyItemID  int        `json:"myItemId" gorm:"primaryKey"`
	UserID    int        `json:"userId"`
	User      *User      `json:"user" gorm:"foreignKey:UserID"`
	ItemID    int        `json:"itemId"`
	Item      *Item      `json:"item" gorm:"foreignKey:ItemID"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}


func SeedMyItem(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&MyItem{
		UserID:    2,
		ItemID:    1,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    1,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    1,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    1,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    1,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    2,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    2,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    3,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    3,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    4,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    4,
	})
	db.Create(&MyItem{
		UserID:    2,
		ItemID:    5,
	})
	db.Create(&MyItem{
		UserID:    3,
		ItemID:    1,
	})
}