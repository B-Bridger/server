package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID   string `gorm:"column:userID;primaryKey;" json:"userID"`
	Password string `gorm:"column:password" json:"-"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email;unique" json:"email"`
	// TODO: Company 테이블 여부 회의 필요
	// CompanyID string    `gorm:"column:companyId" json:"companyId"`
	Language  string    `gorm:"column:language" json:"language"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	FcmToken  string    `gorm:"column:fcmToken" json:"fcmToken"`
}

type CreateUserModel struct {
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Language string `gorm:"column:language" json:"language"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserID == "" {
		u.UserID = uuid.NewString()
	}
	return
}
