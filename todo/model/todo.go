package model

// Array of reminder, each reminder has a time and a message
// to do including user_id, start_time, end_time
type Todo struct {
	ID                string              `json:"id" gorm:"type:char(40)"`
	Title             string              `json:"title" validate:"required" gorm:"type:varchar(100)"`
	Description       string              `json:"description" gorm:"type:text"`
	StartTime         string              `json:"start_time" validate:"required" gorm:"type:datetime"`
	EndTime           string              `json:"end_time" validate:"required" gorm:"type:datetime"`
	// UserID            string              `json:"user_id" validate:"required" gorm:"type:char(40);foreignKey:ID"`
	// User              *User               `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Reminders         []Reminder          `json:"reminders,omitempty" gorm:"foreignKey:TodoID"`
}
