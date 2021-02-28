package model

import "time"

type ReportUser struct {
	ReportID          int        `json:"reportId" gorm:"primaryKey"`
	ReporterID        int        `json:"reporterId"`
	Reporter          *User      `json:"reporter" gorm:"foreignKey:ReporterID"`
	ReportedID        int        `json:"reportedId"`
	Reported          *User      `json:"reported" gorm:"foreignKey:ReportedID"`
	ReportDescription string     `json:"reportDescription"`
	CreatedAt         time.Time  `json:"CreatedAt"`
	UpdatedAt         time.Time  `json:"UpdatedAt"`
	DeletedAt         *time.Time `json:"DeletedAt"`
}