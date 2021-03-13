package model

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type Item struct {
	ItemsID         int        `json:"itemsId" gorm:"primaryKey"`
	GameID          int        `json:"gameId"`
	Game            *Game      `json:"game" gorm:"foreignKey:GameID"`
	Users           []*User    `json:"users" gorm:"many2many:users_item"`
	ItemName        string     `json:"itemName"`
	ItemDescription string     `json:"itemDescription"`
	ItemPhoto       string     `json:"itemPhoto"`
	ItemTransaction []*ItemTransaction `json:"itemTransaction" gorm:"foreignKey:ItemID"`
	MyItem          []*MyItem          `json:"myItem" gorm:"foreignKey:ItemID"`
	CreatedAt       time.Time  `json:"CreatedAt"`
	UpdatedAt       time.Time  `json:"UpdatedAt"`
	DeletedAt       *time.Time `json:"DeletedAt"`
}

func SeedItems(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Create(&Item{
		GameID:          1,
		ItemName:        "Tangan Anti Mage",
		ItemDescription: "Ini Adalah Tangan Dewa Kematian Amat Berharga",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcUzpw9UQU3XQqu63tvQW2J8MApDs6iaKhVtwLzfeClB5OO-gY6fkuXLMqnchVRd6dd2j6eQoNyk2ALt-0s4Nmj3ctOUdwQ7Yl-Eq1K5xuy8hMTuvsiYzyNjuygk-z-DyCwC75PE/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Sword",
		ItemDescription: "Ini adalah pedang juggernaut dari kerajaan majapahit",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YipN_tj5nY2GmOXgMrfugWpD5MpjjtbS4oPnjAWLpxIuNDztIoTGJAc2aVHTqwW9lO3t15e7vciYnHI3vXJz437fy0flgh5MPO090-veFwsXMr_csw/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Topeng Jugger",
		ItemDescription: "Dipakai oleh jugger ketika dia malu sama mukanya",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3fDxR55O_mY2EheP1P4TGl3ld59d1teTA5oekt1i1uRQ5fWn7IYLHJFU4NQ6D-lW6x-jv0JPpuZXAyXYw7HQk4H-In0a_iRwfPbdxxavJt1-Cmzo/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          2,
		ItemName:        "Jugger Buttom2",
		ItemDescription: "Jugger malu kalau ga pake ini2",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
	db.Create(&Item{
		GameID:          1,
		ItemName:        "Jugger Buttom",
		ItemDescription: "Jugger malu kalau ga pake ini",
		ItemPhoto: "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KW1Zwwo4NUX4oFJZEHLbXK9QlSPcU4vBxaSV7eRvG5mMbeXlJmMQVbia2kOQtvwff3YTxM-M_5nY2GmOXgMrfugWpD5MpjjtbR9ILhmwWLpxIuNDztI9KcIANrZ1jZqAC5xu7uhcK8tJ3ByncwuCB34SvUzkHk0h9JO-E5hOveFwvCYx9lzQ/330x192",
	})
}
