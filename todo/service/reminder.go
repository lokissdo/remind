package service

import (
	"context"
	"net/http"
	"time"
	"todo/model"
	"todo/repository"

	"github.com/gin-gonic/gin"
)

func RunReminderCronJob() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Create a context for the reminder service
	ctx := context.Background()

	// Run the reminder service immediately
	repository.SendPendingReminders(ctx)

	// Run the reminder service at every tick of the ticker
	for range ticker.C {
		repository.SendPendingReminders(ctx)
	}
}


// CreateReminder creates a new reminder
func CreateReminder(c *gin.Context) {
    var reminder model.Reminder
    if err := c.ShouldBindJSON(&reminder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse request body"})
        return
    }
    createdReminder, err := repository.CreateReminder(reminder)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reminder"})
        return
    }
    c.JSON(http.StatusOK, createdReminder)
}


// ListReminders lists all reminders
func ListReminders(c *gin.Context) {
	userID := c.Param("userID")
    reminders, err := repository.ListAllTodosByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reminders"})
        return
    }
    c.JSON(http.StatusOK, reminders)
}

// UpdateReminder updates a reminder
func UpdateReminder(c *gin.Context) {
    id := c.Param("id")
    var updatedReminder model.Reminder
    if err := c.ShouldBindJSON(&updatedReminder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse request body"})
        return
    }
    updatedReminder.ID = id
    _, err := repository.UpdateReminderByID(id, updatedReminder)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reminder"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Reminder updated successfully"})
}

// DeleteReminder deletes a reminder by its ID
func DeleteReminder(c *gin.Context) {
    id := c.Param("id")
    err := repository.DeleteReminderByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reminder"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Reminder deleted successfully"})
}
