package model

import "github.com/kendricko-adrio/gqlgen-todos/database"

type Avatar struct {
	AvatarID  int    `json:"avatarId" gorm:"primaryKey"`
	AvatarURL string `json:"avatarUrl"`
	AvatarPoint string `json:"avatarPoint"`
}

func SeedAvatar(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&Avatar{
		AvatarURL: "https://image.flaticon.com/icons/png/512/147/147144.png",
		AvatarPoint: "100",
	})
	db.Create(&Avatar{
		AvatarURL: "https://png.pngtree.com/png-vector/20190225/ourlarge/pngtree-vector-avatar-icon-png-image_702436.jpg",
		AvatarPoint: "200",
	})
}