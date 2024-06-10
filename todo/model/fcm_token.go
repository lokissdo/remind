package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FCMToken struct {
	ID       string `json:"id" gorm:"primaryKey"`
	UserID   string `json:"user_id" gorm:"uniqueIndex"` // Ensure UserID is unique
	FCMToken string `json:"fcm_token" gorm:"type:varchar(255)"`
}

func (c *FCMToken) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = "fcm-" + uuid.New().String()
	}
	return
}
