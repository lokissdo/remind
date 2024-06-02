package repository

// CRUD to do

import (
	"context"
	"fmt"
	"time"
	"todo/database"
	"todo/model"
)



func CreateReminder(reminder model.Reminder) (model.Reminder, error) {
	if err := database.Db.Create(&reminder).Error; err != nil {
		return reminder, err
	}
	return reminder, nil
}

func GetRemindersByTodoID(id string) (model.Reminder, error) {
	var reminder model.Reminder
	if err :=  database.Db.Where("todo_id = ?", id).Take(&reminder).Error; err != nil {
		return reminder, err
	}
	return reminder, nil
}


func UpdateReminderByID(id string, reminder model.Reminder) (model.Reminder, error) {
	if err :=  database.Db.Where("id = ?", id).Updates(&reminder).Error; err != nil {
		return reminder, err
	}
	return reminder, nil
}


func GetNeedSentReminders(ctx context.Context) ([]model.Reminder, error) {
	now := time.Now()

	var reminders []model.Reminder
	query :=  database.Db.Where("start <= ?", now).Find(&reminders)

	// Handle potential cancellation during execution
	if err := query.WithContext(ctx).Error; err != nil {
		return nil, fmt.Errorf("error fetching reminders: %w", err)
	}

	return reminders, nil
}