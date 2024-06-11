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
	if err :=  database.Db.Where("todo_id = ? AND sent = ?", id, false).Take(&reminder).Error; err != nil {
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

func DeleteReminderByID(id string) error {
	var reminder model.Reminder
	if err :=  database.Db.Where("id = ?", id).Delete(&reminder).Error; err != nil {
		return err
	}
	return nil
}


func UpdateBatchReminderSentStatus(ctx context.Context, reminders []model.Reminder) error {
	// using a batch update to update the sent status of reminders
	query :=  database.Db.Model(&model.Reminder{}).Where("id IN ?", reminders).Update("sent", true)

	// Handle potential cancellation during execution
	if err := query.WithContext(ctx).Error; err != nil {
		return fmt.Errorf("error updating reminders: %w", err)
	}
	return nil
}