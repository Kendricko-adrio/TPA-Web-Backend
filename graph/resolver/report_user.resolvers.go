package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"gorm.io/gorm/clause"
)

func (r *queryResolver) InsertReport(ctx context.Context, report model.InsertReport) (*model.ReportUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllReport(ctx context.Context) ([]*model.ReportUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetReportByReported(ctx context.Context, reportedID int) ([]*model.ReportUser, error) {
	//panic(fmt.Errorf("not implemented"))
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	close, err := db.DB()
	defer close.Close()
	var report []*model.ReportUser

	db.Where("reported_id = ?", reportedID).Preload(clause.Associations).Find(&report)

	return report, nil
}
