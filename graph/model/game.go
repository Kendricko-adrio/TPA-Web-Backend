package model

import (
	"fmt"
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"time"
)

type Game struct {
	ID                int              `json:"ID" gorm:"primaryKey"`
	Name              string           `json:"Name"`
	Description       string           `json:"Description"`
	Price             int              `json:"Price"`
	Rating            int              `json:"Rating"`
	ImageBanner       string           `json:"imageBanner"`
	Image             string           `json:"Image"`
	Tag               string           `json:"tag"`
	SystemRequirement string           `json:"systemRequirement"`
	Users             []*User          `json:"users" gorm:"many2many:users_games"`
	Wishlist          []*User          `json:"wishlist" gorm:"many2many:games_wishlist"`
	Genre             []*Genre         `json:"genre" gorm:"many2many:game_genres"`
	GameSlideShow     []*GameSlideShow `json:"gameSlideShow"  gorm:"foreignKey:GameID"`
	Post              []*Post          `json:"post" gorm:"foreignKey:GameID"`
	Items             []*Item          `json:"items" gorm:"foreignKey:GameID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt         time.Time        `json:"CreatedAt"`
	UpdatedAt         time.Time        `json:"UpdatedAt"`
	DeletedAt         *time.Time       `json:"DeletedAt"`
}

func GetAllGames() ([]*Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var games []*Game

	db.Find(&games)
	fmt.Print(games)
	return games, err
}

func SeedGames() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	db.Create(&Game{
		Name:        "Dota",
		Description: "Ini game Dota loh",
		Price:       1000000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 2,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://cdn.akamai.steamstatic.com/steam/apps/570/header.jpg?t=1612665823",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/570/ss_86d675fdc73ba10462abb8f5ece7791c5047072c.600x338.jpg?t=1612665823",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/570/ss_27b6345f22243bd6b885cc64c5cda74e4bd9c3e8.600x338.jpg?t=1612665823",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/570/ss_e0a92f15a6631a8186df79182d0fe28b5e37d8cb.600x338.jpg?t=1612665823",
			},
			{
				SlideShowURL: "https://firebasestorage.googleapis.com/v0/b/staem-a0b09.appspot.com/o/ocirdameyo.mp4?alt=media&token=677a2835-d142-416d-ab0e-bac8d1e9bdeb",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk2",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 3,
			},
			{
				GenreID: 2,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Dota 2",
		Description: "Ini game Dota 2 loh sudah berubah",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 2,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://cdn.akamai.steamstatic.com/steam/apps/570/header.jpg?t=1612665823",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/570/ss_86d675fdc73ba10462abb8f5ece7791c5047072c.600x338.jpg?t=1612665823",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/570/ss_27b6345f22243bd6b885cc64c5cda74e4bd9c3e8.600x338.jpg?t=1612665823",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/570/ss_e0a92f15a6631a8186df79182d0fe28b5e37d8cb.600x338.jpg?t=1612665823",
			},
		},
	})
	db.Create(&Game{
		Name:        "Counter-Strike: Global Offensive",
		Description: "CS:GO Kuy",
		Price:       150000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 2,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/730/ss_118cb022b9a43f70d2e5a2df7427f29088b6b191.600x338.jpg?t=1612812939",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/730/ss_34090867f1a02b6c17652ba9043e3f622ed985a9.600x338.jpg?t=1612812939",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/730/ss_34090867f1a02b6c17652ba9043e3f622ed985a9.600x338.jpg?t=1612812939",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/730/ss_60b4f959497899515f46012df805b0006ef21af6.600x338.jpg?t=1612812939",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk5",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 2,
			},
			{
				GenreID: 4,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk6",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 2,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk7",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 2,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk8",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 2,
			},
			{
				GenreID: 3,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk9",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 2,
			},
			{
				GenreID: 4,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk10",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 3,
			},
			{
				GenreID: 4,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk11",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 4,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk12",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 3,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk13",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 4,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk14",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 4,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk15",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 3,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})
	db.Create(&Game{
		Name:        "Cyber Punk16",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 2,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
			},
		},
	})

	db.Create(&Game{
		Name:        "Counter-Strike: Doesn't have counter",
		Description: "cs sdf",
		Price:       1000,
		Rating:      2,
		Genre: []*Genre{
			{
				GenreID: 1,
			},
			{
				GenreID: 2,
			},
		},
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		ImageBanner: "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",
		GameSlideShow: []*GameSlideShow{
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/730/ss_118cb022b9a43f70d2e5a2df7427f29088b6b191.600x338.jpg?t=1612812939",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/730/ss_34090867f1a02b6c17652ba9043e3f622ed985a9.600x338.jpg?t=1612812939",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/730/ss_34090867f1a02b6c17652ba9043e3f622ed985a9.600x338.jpg?t=1612812939",
			},
			{
				SlideShowURL: "https://cdn.akamai.steamstatic.com/steam/apps/730/ss_60b4f959497899515f46012df805b0006ef21af6.600x338.jpg?t=1612812939",
			},
		},
	})

}
