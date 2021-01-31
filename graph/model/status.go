package model

type Status struct {
	StatusID   int `gorm:"primaryKey"`
	StatusName string `json:"statusName"`
}
