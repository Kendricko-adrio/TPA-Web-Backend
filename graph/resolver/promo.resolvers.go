package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) InsertPromo(ctx context.Context, promo model.PromoInput) (*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var promoInsert = &model.Promo{}
	promoInsert.GameID = promo.GameID
	promoInsert.PromoDiscount = promo.PromoDiscount
	promoInsert.PromoDuration = promo.PromoDuration
	db.Create(&promoInsert)

	return promoInsert, nil
}

func (r *mutationResolver) UpdatePromo(ctx context.Context, promo model.PromoInput, id int) (*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var promoInsert = &model.Promo{PromoID: id}

	db.First(&promoInsert)

	promoInsert.GameID = promo.GameID
	promoInsert.PromoDiscount = promo.PromoDiscount
	promoInsert.PromoDuration = promo.PromoDuration
	db.Save(&promoInsert)

	return promoInsert, nil
}

func (r *mutationResolver) DeletePromo(ctx context.Context, id int) (*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	promo := &model.Promo{PromoID: id}

	db.First(&promo)
	db.Delete(&promo)

	return promo, nil
}

func (r *queryResolver) GetAllPromo(ctx context.Context, page int) ([]*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var allPromo []*model.Promo

	db.Preload(clause.Associations).Limit(10).Offset((page - 1) * 10).Find(&allPromo)

	return allPromo, nil
}

func (r *queryResolver) GetPromoByID(ctx context.Context, id int) (*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var promo = &model.Promo{PromoID: id}

	db.Find(promo)

	return promo, nil
}

func (r *queryResolver) GetTotalPromo(ctx context.Context) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var allPromo []*model.Promo

	var test = db.Find(&allPromo)

	return int(test.RowsAffected), nil
}

func (r *queryResolver) GetOnSale(ctx context.Context) ([]*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var allPromo []*model.Promo
	db.Where("promo_discount >= 50").Limit(5).Preload(clause.Associations).Find(&allPromo)

	return allPromo, nil
}
