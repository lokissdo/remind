package service

import (
	"log"
	"net/http"
	"todo/model"
	"todo/repository"

	"github.com/gin-gonic/gin"
)

// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
    var todo model.Todo
	userID := c.Param("userID")
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse request body"})
        return
    }

	todo.UserID = userID

	for i := range todo.Reminders {
		todo.Reminders[i].UserID = userID
	}
	for _, reminder := range todo.Reminders {
		log.Printf("Reminder: %v", reminder.UserID)
	}
    createdTodo, err := repository.CreateTodo(todo)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
        return
    }
    c.JSON(http.StatusOK, createdTodo)
}

// GetTodoByID retrieves a todo by its ID
func GetTodoByID(c *gin.Context) {
    id := c.Param("id")
    todo, err := repository.GetTodoByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    c.JSON(http.StatusOK, todo)
}

// ListTodos lists all todos
func ListTodos(c *gin.Context) {
	userID := c.Param("userID")
    todos, err := repository.ListAllTodosByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
        return
    }
    c.JSON(http.StatusOK, todos)
}

// UpdateTodo updates a todo
func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var updatedTodo model.Todo
    if err := c.ShouldBindJSON(&updatedTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse request body"})
        return
    }
    updatedTodo.ID = id
    _, err := repository.UpdateTodoByID(updatedTodo.ID, updatedTodo)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

// DeleteTodo deletes a todo by its ID
func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    err := repository.DeleteTodoByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
