package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Paginate(c *gin.Context) (int, int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 10 {
		size = 10
	}
	return page, size
}
