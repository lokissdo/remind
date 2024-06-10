package api

import (
	//"happy-insight/api/middleware"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// do something
	// r.POST("/login", user.Login)
	// r.POST("/refresh", user.RefreshToken)
	// r.POST("/logout", user.Logout)
	// r.POST("/password", user.ChangePassword)
	// r.POST("/avatar", user.ChangeAvatar)
	// go user.CleanUpToken()
	// apiV1 := r.Group("/api/")
	// adminHIXRoute := apiV1.Group("/admin-hix")
	// {
	// 	adminHIXRoute.Use(middleware.RequiredAuth(models.ADMIN_HIX_ROLE))
	// 	adminHIXRoute.POST("/school", admin_hix.AddSchool)
	// 	adminHIXRoute.GET("/school/:school_id", admin_hix.GetSchool)
	

	// }


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	

}
