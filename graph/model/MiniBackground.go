package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type MiniBackground struct {
	MiniID    int        `json:"miniId" gorm:"primaryKey"`
	MiniURL   string     `json:"miniUrl"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}

func SeedMiniBackground (){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}
	db.Create(&MiniBackground{
		MiniURL:   "https://img.freepik.com/free-photo/hand-painted-watercolor-background-with-sky-clouds-shape_24972-1095.jpg?size=626&ext=jpg",

	})
	db.Create(&MiniBackground{
		MiniURL:   "https://venngage-wordpress.s3.amazonaws.com/uploads/2018/09/Colorful-Circle-Simple-Background-Image-1.jpg",

	})
	db.Create(&MiniBackground{
		MiniURL: "https://img.rawpixel.com/s3fs-private/rawpixel_images/website_content/v462-n-130-textureidea_1.jpg?w=800&dpr=1&fit=default&crop=default&q=65&vib=3&con=3&usm=15&bg=F4F4F3&ixlib=js-2.2.1&s=9465282a2b0a375f4f5b120d7bbad882",
	})


}
