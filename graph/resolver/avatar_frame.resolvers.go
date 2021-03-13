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

func (r *mutationResolver) InsertFrame(ctx context.Context, frameID int) (*model.AvatarFrame, error) {
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

	db.First(&user)
	frame := model.AvatarFrame{FrameID: frameID}
	db.First(&frame)
	if user.Point < frame.FramePoint {
		return nil, fmt.Errorf("Not Enough Money")
	}
	user.Point -= frame.FramePoint
	db.Exec("INSERT INTO users_frame VALUES(?,?)", user.UserID, frameID)
	db.Save(user.OwnFrame)

	return &frame, nil
}

func (r *queryResolver) GetPurchasableFrame(ctx context.Context) ([]*model.AvatarFrame, error) {
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
	var frame []*model.AvatarFrame
	//arrLen := len(user.OwnProfileBackground)
	var arr []int
	for _, e := range user.OwnFrame {
		arr = append(arr, e.FrameID)
	}
	fmt.Println(arr)
	db.Where("frame_id NOT IN ?", arr).Find(&frame)

	return frame, nil
}
