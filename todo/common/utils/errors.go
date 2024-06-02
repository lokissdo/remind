package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"github.com/jackc/pgconn"
)

// SetResponseError sets the response based on the error
func SetResponseError(c *gin.Context, err error) {
	c.Set("error", err.Error())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if strings.Contains(err.Error(), "already exists") {
		c.JSON(http.StatusConflict, gin.H{"error": "Already exists"})
		return
	}
	if strings.Contains(err.Error(), "invalid input") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if _, ok := err.(validator.ValidationErrors); ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case "23505":
			c.JSON(http.StatusConflict, gin.H{"error": "Duplicate entry"})
		case "23503":
			c.JSON(http.StatusNotFound, gin.H{"error": "Foreign key constraint fails"})
		case "22001":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data too long for column"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
}
