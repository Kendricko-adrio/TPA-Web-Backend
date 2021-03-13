package model

import "github.com/kendricko-adrio/gqlgen-todos/database"

type ProfileBackground struct {
	BackgroundID  int    `json:"backgroundId" gorm:"primaryKey"`
	BackgroundURL string `json:"backgroundUrl"`
	BackgroundPoint int    `json:"backgroundPoint"`
}

func SeedProfileBackground(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	db.Create(&ProfileBackground{
		BackgroundURL: "https://wallpaperaccess.com/full/1205014.jpg",
		BackgroundPoint: 100,
	})
	db.Create(&ProfileBackground{
		BackgroundURL: "https://previews.123rf.com/images/claudiodivizia/claudiodivizia1612/claudiodivizia161203100/68733964-light-blue-paper-texture-useful-as-a-background-vertical.jpg",
		BackgroundPoint: 100,
	})
	db.Create(&ProfileBackground{
		BackgroundURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRUB1_qVkIOGEWMyw1JtiroMh3t-qP-HgJFdw&usqp=CAU",
		BackgroundPoint: 100,
	})
	db.Create(&ProfileBackground{
		BackgroundURL: "https://www.newsindiatimes.com/wp-content/uploads/2020/05/abstract-technology-particle-background_52683-25766.jpg",
		BackgroundPoint: 100,
	})


}
