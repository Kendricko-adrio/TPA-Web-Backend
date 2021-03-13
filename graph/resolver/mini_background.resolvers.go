package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
	"gorm.io/gorm/clause"
)

func (r *queryResolver) GetPurchasableMini(ctx context.Context) ([]*model.MiniBackground, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Preload(clause.Associations).First(&user)
	var bg []*model.MiniBackground
	//arrLen := len(user.OwnProfileBackground)
	var arr []int
	for _, e := range user.OwnMiniBg {
		arr = append(arr, e.MiniID)
	}
	fmt.Println(arr)
	db.Where("mini_id NOT IN ?", arr).Find(&bg)

	return bg, nil
}
