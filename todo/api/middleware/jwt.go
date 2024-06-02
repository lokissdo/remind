package middleware

// import (
// 	"net/http"
// 	"todo/common/utils"

// 	"github.com/gin-gonic/gin"
// )

// func RequiredAuth(role string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var token string
// 		if token = c.GetHeader("Authorization"); token == "" {
// 			token = c.Query("authorization")
// 			if token == "" {
// 				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 				c.Abort()
// 				return
// 			}
// 		}
// 		claim, err := utils.ValidateToken(token, "Bearer")
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 			c.Abort()
// 			return
// 		}
// 		for _, roleJWT := range claim.UserRole {
// 			if roleJWT.RoleType == role {
// 				c.Set("user_id", roleJWT.RoleID)
// 				c.Set("school_id", roleJWT.SchoolID)
// 				c.Next()
// 				return
// 			}
// 		}
// 		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
// 		c.Abort()
// 	}
// }
