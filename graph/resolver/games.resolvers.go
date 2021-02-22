package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/generated"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
)

func (r *mutationResolver) InsertGame(ctx context.Context, game model.GameInput) (*model.Game, error) {
	//panic(fmt.Errorf("not implemented"))
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var gameDb model.Game
	gameDb.ID = game.ID
	testDb := db.First(&gameDb)

	// insert
	fmt.Println("test row ", testDb.RowsAffected)
	if testDb.RowsAffected == 0 {
		fmt.Println("insert gan")
		db.Create(&model.Game{
			Name:              game.Name,
			Description:       game.Description,
			Price:             game.Price,
			Rating:            game.Rating,
			ImageBanner:       game.ImageBanner,
			Image:             game.Image,
			Tag:               game.Tag,
			SystemRequirement: game.SystemRequirement,
		})
	}
	//fmt.Println(gameDb)

	return &gameDb, nil
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

	var games []*model.Game

	db.Offset((page - 1) * 10).Limit(10).Find(&games)
	return games, err
}

func (r *queryResolver) GetGameByID(ctx context.Context, id int) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	game := &model.Game{ID: id}
	db.First(game)

	return game, nil
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
