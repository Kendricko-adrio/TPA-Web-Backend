package model

import "time"

type Chat struct {
	ChatID     int       `json:"chatId" gorm:"primaryKey"`
	Text       string    `json:"text"`
	ReceiverID int       `json:"receiverId"`
	Receiver   *User     `json:"receiver" gorm:"foreignKey:ReceiverID"`
	SenderID   int       `json:"senderId"`
	Sender     *User     `json:"sender" gorm:"foreignKey:SenderID"`
	CreatedAt  time.Time `json:"CreatedAt"`
}