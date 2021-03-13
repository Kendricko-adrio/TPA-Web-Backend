package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"gorm.io/gorm/clause"
)

func (r *queryResolver) GetUserItemsByGamePaginate(ctx context.Context, gameID int, userID int, page int, search *string) ([]*model.Item, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	var item []*model.Item

	//db.Joins("JOIN users_item AS ui ON ui.item_items_id = items.items_id").Where("game_id = ? AND ui.user_user_id = ? AND LOWER(items.item_name) LIKE LOWER(?)", gameID, userID, "%"+*search+"%").Offset((page - 1) * 10).Limit(10).Preload(clause.Associations).Find(&item)
	db.Joins("JOIN my_items AS ui ON ui.item_id = items.items_id").Where("game_id = ? AND ui.user_id = ? AND LOWER(items.item_name) LIKE LOWER(?)", gameID, userID, "%"+*search+"%").Offset((page - 1) * 10).Limit(10).Preload(clause.Associations).Find(&item)

	return item, nil
}

func (r *queryResolver) GetItem(ctx context.Context, itemID int) (*model.Item, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	item := model.Item{ItemsID: itemID}
	db.Preload(clause.Associations).First(&item)

	return &item, nil
}
