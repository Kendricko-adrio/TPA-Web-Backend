package model

import "github.com/kendricko-adrio/gqlgen-todos/database"

type AvatarFrame struct {
	FrameID  int    `json:"frameId" gorm:"primaryKey"`
	FrameURL string `json:"frameUrl"`
	FramePoint int    `json:"framePoint"`
}


func SeedAvatarFrame(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&AvatarFrame{
		FrameURL: "https://cdn.akamai.steamstatic.com/steamcommunity/public/images/items/1492660/06bb85cd5f39a963a39ae9327ea4eb7da5cd30d4.png",
		FramePoint: 100,
	})
	db.Create(&AvatarFrame{
		FrameURL: "https://cdn.akamai.steamstatic.com/steamcommunity/public/images/items/730/c2a822b6422d7a2c9e72ec74bf54139db8723ef0.png",
		FramePoint: 100,
	})
	db.Create(&AvatarFrame{
		FrameURL: "https://cdn.akamai.steamstatic.com/steamcommunity/public/images/items/730/c30260bb120bf1379f075802653c8eb86da7a7e9.png",
		FramePoint: 100,
	})
	db.Create(&AvatarFrame{
		FrameURL: "https://cdn.akamai.steamstatic.com/steamcommunity/public/images/items/322330/46461aaea39b18a4a3da2e6d3cf253006f2d6193.png",
		FramePoint: 100,
	})

}