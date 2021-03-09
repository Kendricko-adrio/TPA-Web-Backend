package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type Post struct {
	PostID          int              `json:"postId" gorm:"primaryKey"`
	UserID          int              `json:"userId"`
	GameID          int              `json:"gameId"`
	PostType        *PostType        `json:"postType" gorm:"foreignKey:PostTypeID"`
	PostTypeID      int              `json:"postTypeId"`
	PostHelpful     bool             `json:"postHelpful"`
	PostDescription string           `json:"postDescription"`
	PostTitle       string           `json:"postTitle"`
	PostAsset       string           `json:"postAsset"`
	TotalLike       int              `json:"totalLike"`
	TotalDislike    int              `json:"totalDislike"`
	LikeDetail      []*LikeDetail    `json:"likeDetail" gorm:"foreignKey:PostID"`
	CommandDetail   []*CommandDetail `json:"commandDetail" gorm:"foreignKey:PostID"`
	CreatedAt       time.Time        `json:"CreatedAt"`
	UpdatedAt       time.Time        `json:"UpdatedAt"`
	DeletedAt       *time.Time       `json:"DeletedAt"`
}

func SeedPost() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&Post{
		UserID:          2,
		GameID:          1,
		PostTypeID:      1,
		PostHelpful:     false,
		TotalLike:       1,
		PostDescription: "Hai kamu harus bener ya",
		PostAsset:       "https://cdn-5c489073f911c8147449b474.closte.com/wp-content/uploads/2020/02/Dota-2-scaled-1-1600x901.jpg",
		DeletedAt:       nil,
	})
	db.Create(&Post{
		UserID:          2,
		GameID:          1,
		PostTypeID:      1,
		PostHelpful:     false,
		PostDescription: "Hai kamu harus bener ya",
		PostAsset:       "https://cdn-5c489073f911c8147449b474.closte.com/wp-content/uploads/2020/02/Dota-2-scaled-1-1600x901.jpg",
		DeletedAt:       nil,
	})
	db.Create(&Post{
		UserID:          2,
		GameID:          1,
		PostTypeID:      1,
		PostHelpful:     false,
		PostDescription: "Hai kamu harus bener ya",
		PostAsset:       "https://cdn-5c489073f911c8147449b474.closte.com/wp-content/uploads/2020/02/Dota-2-scaled-1-1600x901.jpg",
		DeletedAt:       nil,
	})
	db.Create(&Post{
		UserID:          2,
		GameID:          1,
		PostTypeID:      1,
		PostHelpful:     false,
		PostDescription: "Hai kamu harus bener ya",
		PostAsset:       "https://firebasestorage.googleapis.com/v0/b/staem-a0b09.appspot.com/o/ocirdameyo.mp4?alt=media&token=677a2835-d142-416d-ab0e-bac8d1e9bdeb",
		DeletedAt:       nil,
	})
	db.Create(&Post{
		UserID:          2,
		GameID:          1,
		PostTypeID:      2,
		TotalLike: 100,
		PostHelpful:     true,
		PostDescription: "Hai kamu harus bener ya",
	})
	db.Create(&Post{
		UserID:          2,
		GameID:          1,
		PostTypeID:      2,
		PostHelpful:     true,
		PostDescription: "Hai kamu harus bener ya",
	})
	db.Create(&Post{
		UserID:          3,
		GameID:          1,
		PostTypeID:      2,
		PostHelpful:     false,
		PostTitle: "ini adalah title hayo",
		PostDescription: "Hai kamu harus bener ya",
	})
	db.Create(&Post{
		UserID:          3,
		GameID:          1,
		PostTypeID:      3,
		PostHelpful:     false,
		PostTitle: "ini adalah title kamu makannya apa",
		PostDescription: "Hai kamu harus bener ya",
	})
	db.Create(&Post{
		UserID:          4,
		GameID:          1,
		PostTypeID:      3,
		PostHelpful:     false,
		PostTitle: "ini adalah contoh",
		PostDescription: "Kamu makannya apa",
	})
	db.Create(&Post{
		UserID:          4,
		GameID:          1,
		PostTypeID:      3,
		PostHelpful:     false,
		PostTitle: "kamu adalah mahluk Tuhan yang paling skesi",
		PostDescription: "Saya juru masaknya",
	})
	db.Create(&Post{
		UserID:          3,
		GameID:          2,
		PostTypeID:      3,
		PostHelpful:     false,
		PostTitle: "ini adalah title",
		PostDescription: "APA INI HAHAHAHHA",
	})
	db.Create(&Post{
		UserID:          3,
		GameID:          2,
		PostTypeID:      3,
		PostHelpful:     false,
		PostTitle: "HHHHHHHAAAAAAAAAAAAAAAAAHHHHHHHHH????",
		PostDescription: "Masa iya sih",
	})
}
