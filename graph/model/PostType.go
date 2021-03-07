package model

import "github.com/kendricko-adrio/gqlgen-todos/database"

type PostType struct {
	PostTypeID   int    `json:"postTypeId" gorm:"primaryKey"`
	PostTypeName string `json:"postTypeName"`
}

func SeedPostType(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&PostType{
		PostTypeName: "Poster",
	})
	db.Create(&PostType{
		PostTypeName: "Review",
	})
	db.Create(&PostType{
		PostTypeName: "Discussion",
	})

}