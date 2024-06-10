package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reminder struct {
	ID      string `json:"id" gorm:"type:char(40)"`
	Start   string `json:"start" validate:"required" gorm:"type:datetime"`
	Message string `json:"message" gorm:"type:text"`
	TodoID  string `json:"todo_id" gorm:"type:char(40);foreignKey:ID"`
	Todo    *Todo  `json:"todo,omitempty" gorm:"foreignKey:TodoID"`
	UserID  string `json:"user_id" validate:"required" gorm:"type:char(40)"`
	Sent    bool   `json:"sent" gorm:"type:boolean"`
}

func (c *Reminder) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = "rem-" + uuid.New().String()
	}
	return
}
