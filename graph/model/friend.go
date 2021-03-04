package model

import "time"

type FriendsDetail struct {
	User1          *User         `json:"user1" gorm:"foreignKey:User1id"`
	User1id        int           `json:"user1ID" gorm:"primaryKey"`
	User2          *User         `json:"user2" gorm:"foreignKey:User2id"`
	User2id        int           `json:"user2ID" gorm:"primaryKey"`
	FriendStatus   *FriendStatus `json:"friendStatus" gorm:"foreignKey:FriendStatusID"`
	FriendStatusID int           `json:"friendStatusId"`
	CreatedAt      time.Time     `json:"CreatedAt"`
	UpdatedAt      time.Time     `json:"UpdatedAt"`
	DeletedAt      *time.Time    `json:"DeletedAt"`
}
