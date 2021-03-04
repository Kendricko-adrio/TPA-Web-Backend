package model

type FriendStatus struct {
	FriendStatusID int    `json:"friendStatusId" gorm:"primaryKey"`
	StatusName     string `json:"statusName"`
}