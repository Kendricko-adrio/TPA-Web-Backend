package model

import "github.com/kendricko-adrio/gqlgen-todos/database"

type Theme struct {
	ThemeID   int    `json:"themeId" gorm:"primaryKey"`
	ThemeName string `json:"themeName"`
}

func SeedTheme(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&Theme{
		ThemeName: "red",
	})
	db.Create(&Theme{
		ThemeName: "blue",
	})
	db.Create(&Theme{
		ThemeName: "yellow",
	})
}