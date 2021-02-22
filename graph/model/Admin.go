package model

import "time"

type Admin struct {
	AdminID       int        `json:"adminId" gorm:"primaryKey"`
	AdminUsername string     `json:"adminUsername"`
	AdminPassword string     `json:"adminPassword"`
	CreatedAt     time.Time  `json:"CreatedAt"`
	UpdatedAt     time.Time  `json:"UpdatedAt"`
	DeletedAt     *time.Time `json:"DeletedAt"`
}
