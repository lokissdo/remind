package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          string     `json:"id" gorm:"type:varchar(40)"`
	Title       string     `json:"title" validate:"required" gorm:"type:varchar(100)"`
	Description string     `json:"description" gorm:"type:text"`
	StartTime   time.Time  `json:"start_time" validate:"required"`
	EndTime     time.Time  `json:"end_time" validate:"required"`
	UserID      string     `json:"user_id" validate:"required" gorm:"type:varchar(40)"`
	Reminders   []Reminder `json:"reminders,omitempty" gorm:"foreignKey:TodoID;constraint:OnDelete:CASCADE"`
}
func (c *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = "tod-" + uuid.New().String()
	}
	return
}
