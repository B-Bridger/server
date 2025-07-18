package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// `ChatRoom` belongs to `User`, `UserID` is the foreign key
type ChatRoom struct {
	ChatRoomID    string    `gorm:"column:chatRoomID;primaryKey;" json:"chatRoomID"`
	UserID        string    `gorm:"column:ownerUserID" json:"-"`
	Owner         User      `gorm:"foreignKey:UserID;references:UserID" json:"owner"`
	LastMessage   string    `gorm:"column:lastMessage" json:"lastMessage"`
	LastMessageAt time.Time `gorm:"column:lastMessageAt" json:"lastMessageAt"`
	CreatedAt     time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
}

type CreateChatRoomModel struct {
	InviteUserIDS []string `json:"inviteUserIDs"`
}

func (cr *ChatRoom) BeforeCreate(tx *gorm.DB) (err error) {
	if cr.ChatRoomID == "" {
		cr.ChatRoomID = uuid.NewString()
	}

	return
}
