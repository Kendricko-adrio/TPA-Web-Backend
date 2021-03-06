package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type Badge struct {
	BadgeID   int        `json:"badgeId" gorm:"primaryKey"`
	BadgeName string     `json:"badgeName"`
	BadgeURL  string     `json:"badgeUrl"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}

func SeedBadge(){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}
	db.Create(&Badge{
		BadgeName: "Sharp-Eyes!",
		BadgeURL:  "https://community.akamai.steamstatic.com/public/images/badges/13_gamecollector/25_54.png?v=4",
	})
	db.Create(&Badge{
		BadgeName: "Pillars of Community",
		BadgeURL:  "https://community.akamai.steamstatic.com/public/images/badges/01_community/community02_54.png",
	})
	db.Create(&Badge{
		BadgeName: "7 years with us",
		BadgeURL:  "https://community.akamai.steamstatic.com/public/images/badges/02_years/steamyears7_54.png",
	})
	db.Create(&Badge{
		BadgeName: "winner cow winner",
		BadgeURL:  "https://community.akamai.steamstatic.com/public/images/badges/23_towerattack/300.png",
	})
	db.Create(&Badge{
		BadgeName: "CS MASTER",
		BadgeURL:  "https://cdn.akamai.steamstatic.com/steamcommunity/public/images/items/730/54e40b9e2288fbab8bd4c6537b0325d405c7e1b0.png",
	})
}
