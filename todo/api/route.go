package api

import (
	//"happy-insight/api/middleware"

	admin_hix "happy-insight/admin_hix"
	"happy-insight/admin_shix"
	"happy-insight/models"
	"happy-insight/parent"
	"happy-insight/psychologist"
	"happy-insight/student"
	"happy-insight/teacher"

	"happy-insight/api/middleware"
	"happy-insight/user"

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


	r.GET()
	

}
