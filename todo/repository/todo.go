package repository

import (
	"remind/todo/database"
	"remind/todo/model"
)

// CreateTodo creates a new todo
func CreateTodo(todo model.Todo) (model.Todo, error) {
	if err := database.Db.Create(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func ListAllTodosByUserID(userID string) ([]model.Todo, error) {

	var todos []model.Todo
	if err := database.Db.Preload("Reminders").Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// GetTodoByID retrieves a todo by its ID along with its reminders
func GetTodoByID(id string) (model.Todo, error) {
	var todo model.Todo
	if err := database.Db.Where("id = ?", id).Preload("Reminders").First(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

// UpdateTodoByID updates a todo by its ID
func UpdateTodoByID(id string, updatedTodo model.Todo) (model.Todo, error) {
	var todo model.Todo
	if err := database.Db.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	}

	if err := database.Db.Model(&todo).Updates(&updatedTodo).Error; err != nil {
		return todo, err
	}

	return updatedTodo, nil
}

// DeleteTodoByID deletes a todo by its ID
func DeleteTodoByID(id string) error {
	var todo model.Todo
	if err := database.Db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}
