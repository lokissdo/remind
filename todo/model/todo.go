package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID                string              `json:"id" gorm:"type:char(40)"`
	Title             string              `json:"title" validate:"required" gorm:"type:varchar(100)"`
	Description       string              `json:"description" gorm:"type:text"`
	StartTime         string              `json:"start_time" validate:"required" gorm:"type:datetime"`
	EndTime           string              `json:"end_time" validate:"required" gorm:"type:datetime"`
	UserID            string              `json:"user_id" validate:"required" gorm:"type:char(40)"`
	Reminders         []Reminder          `json:"reminders,omitempty" gorm:"foreignKey:TodoID"`
}


func (c *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = "tod-" + uuid.New().String()
	}
	return
}
