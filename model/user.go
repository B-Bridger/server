package model

import "time"

type User struct {
	UserID   string `gorm:"column:userID;primaryKey" json:"userID"`
	Password string `gorm:"column:password" json:"-"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	// TODO: Company 테이블 여부 회의 필요
	// CompanyID string    `gorm:"column:companyId" json:"companyId"`
	Language  string    `gorm:"column:language" json:"language"`
	Role      string    `gorm:"column:role" json:"role"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	FcmToken  string    `gorm:"column:fcmToken" json:"fcmToken"`
}
