package api

import (
	"github.com/gin-gonic/gin"
	"todo/service"
)

func Register(r *gin.Engine) {
	// Ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Todo routes
	todoRoutes := r.Group("/todos")
	{
		todoRoutes.POST("/:userID", service.CreateTodo)
		todoRoutes.GET("/:id", service.GetTodoByID)
		todoRoutes.GET("/user/:userID", service.ListTodos)
		todoRoutes.PUT("/:id", service.UpdateTodo)
		todoRoutes.DELETE("/:id", service.DeleteTodo)
	}

	// Reminder routes
	reminderRoutes := r.Group("/reminders")
	{
		reminderRoutes.POST("/", service.CreateReminder)
		reminderRoutes.GET("/user/:userID", service.ListReminders)
		reminderRoutes.PUT("/:id", service.UpdateReminder)
		reminderRoutes.DELETE("/:id", service.DeleteReminder)
	}

	// Run reminder cron job
	go service.RunReminderCronJob()
}
