package model


type Reminder struct {
	ID        string `json:"id" gorm:"type:char(40)"`
	Start      string `json:"start" validate:"required" gorm:"type:datetime"`
	Message   string `json:"message" gorm:"type:text"`
	TodoID    string `json:"todo_id" gorm:"type:char(40);foreignKey:ID"`
	Todo      *Todo  `json:"todo,omitempty" gorm:"foreignKey:TodoID"`
}