package model

import "time"

type UnsuspendRequest struct {
	RequestID       int                   `json:"requestId" gorm:"primaryKey"`
	SuspendedUser   *User                 `json:"suspendedUser" gorm:"foreignKey:SuspendedUserID"`
	SuspendedUserID int                   `json:"suspendedUserId"`
	SuspendedTypeID int                   `json:"suspendedTypeId"`
	SuspendedType   *UnsuspendRequestType `json:"suspendedType" gorm:"foreignKey:SuspendedTypeID"`
	CreatedAt       time.Time             `json:"CreatedAt"`
	UpdatedAt       time.Time             `json:"UpdatedAt"`
	DeletedAt       *time.Time            `json:"DeletedAt"`
}