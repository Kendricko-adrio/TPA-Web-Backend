package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/generated"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) InsertGame(ctx context.Context, game model.GameInput) (*model.Game, error) {
	//panic(fmt.Errorf("not implemented"))
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var slideArr []*model.GameSlideShow
	var genreArr []*model.Genre

	for _, x := range game.Genre {
		genreArr = append(genreArr, &model.Genre{GenreID: x})
	}

	for _, x := range game.GameSlideShow {
		slideArr = append(slideArr, &model.GameSlideShow{SlideShowURL: x})
	}

	fmt.Println(game.Genre)
	fmt.Println(game.GameSlideShow)
	//fmt.Println(game.Genre)
	//return nil, err
	insertedGame := &model.Game{
		Name:              game.Name,
		Description:       game.Description,
		Price:             game.Price,
		Rating:            game.Rating,
		ImageBanner:       game.ImageBanner,
		Image:             game.Image,
		Genre:             genreArr,
		GameSlideShow:     slideArr,
		SystemRequirement: game.SystemRequirement,
	}
	fmt.Println("insert gan")
	db.Create(insertedGame)
	//fmt.Println(gameDb)

	return insertedGame, nil
}

func (r *mutationResolver) UpdateGame(ctx context.Context, game model.GameInput) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var updateGame = &model.Game{ID: game.ID}
	db.Find(updateGame)

	updateGame.SystemRequirement = game.SystemRequirement
	//updateGame.Tag = game.Tag
	updateGame.Image = game.Image
	updateGame.ImageBanner = game.ImageBanner
	updateGame.Rating = game.Rating
	updateGame.Price = game.Price
	updateGame.Description = game.Description
	updateGame.Name = game.Name

	db.Save(updateGame)
	//fmt.Println(gameDb)

	return updateGame, nil
}

func (r *queryResolver) AllGame(ctx context.Context) ([]*model.Game, error) {
	games, err := model.GetAllGames()

	return games, err
}

func (r *queryResolver) GetAllGamesPaginated(ctx context.Context, page int) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var games []*model.Game

	db.Offset((page - 1) * 10).Limit(10).Find(&games)
	return games, err
}

func (r *queryResolver) GetGameByID(ctx context.Context, id int) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	game := &model.Game{ID: id}
	db.Preload(clause.Associations).First(game)
	return game, nil
}

func (r *queryResolver) SearchGameByTitle(ctx context.Context, title string) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var games []*model.Game
	db.Where("LOWER(name) LIKE LOWER(?) ", "%"+title+"%").Debug().Limit(5).Find(&games)

	return games, nil
}

func (r *queryResolver) SearchGameInfinite(ctx context.Context, title string, page int) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var searchGame []*model.Game

	//db.Where("LOWER(name) = LOWER(?)", "%" + title + "%").Limit(10 * page).Find(&searchGame)
	db.Where("LOWER(name) LIKE LOWER(?) ", "%"+title+"%").Debug().Limit(10 * page).Find(&searchGame)
	return searchGame, nil
}

func (r *queryResolver) GetTotalGame(ctx context.Context) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var searchGame []*model.Game
	test := db.Find(&searchGame)
	return int(test.RowsAffected), nil
}

func (r *queryResolver) GetFilterGame(ctx context.Context, genre int, price int, title string) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var searchGame []*model.Game

	db.Where("price <= ? AND LOWER(name) LIKE LOWER(?)", price, "%"+title+"%").Debug().Find(&searchGame)

	return searchGame, nil
}

func (r *queryResolver) GetNewGame(ctx context.Context) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var searchGame []*model.Game
	db.Order("created_at desc").Limit(10).Preload(clause.Associations).Find(&searchGame)

	return searchGame, nil
}

func (r *queryResolver) GetTopSeller(ctx context.Context) ([]*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type gameResolver struct{ *Resolver }
