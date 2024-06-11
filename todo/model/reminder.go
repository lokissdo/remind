package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reminder struct {
	ID      string    `json:"id" gorm:"type:varchar(40)"`
	Start   time.Time `json:"start" validate:"required" `
	Message string    `json:"message" gorm:"type:text"`
	TodoID  string    `json:"todo_id" gorm:"type:varchar(40);foreignKey:ID"`
	Todo    *Todo     `json:"todo,omitempty" gorm:"foreignKey:TodoID"`
	UserID  string    `json:"user_id" validate:"required" gorm:"type:varchar(40)"`
	Sent    bool      `json:"sent" gorm:"type:boolean"`
}

func (c *Reminder) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = "rem-" + uuid.New().String()
	}
	return
}
