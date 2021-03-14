package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"github.com/kendricko-adrio/gqlgen-todos/graph/myredis"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
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

	db.Preload(clause.Associations).Limit(10).Offset((page - 1) * 10).Debug().Find(&allPromo)

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
	user := middleware.ForContext(ctx)
	if user == nil {
		return 0, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	var allPromo []*model.Promo
	cached, err := myredis.UseCache().Get(ctx, strconv.Itoa(user.UserID)).Result()

	if err == nil {
		fmt.Println("Masuk sini")
		hasilnya, _ := strconv.Atoi(cached)
		return hasilnya, nil
	}

	var test = db.Find(&allPromo)

	if err := myredis.UseCache().Set(ctx, strconv.Itoa(user.UserID), int(test.RowsAffected), 10*time.Second).Err(); err != nil {
		log.Fatal(err)
	}

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
	db.Unscoped().Debug().Limit(5).Preload(clause.Associations).Find(&allPromo)

	return allPromo, nil
}
