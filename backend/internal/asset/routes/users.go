package routes

import (
	"assetManagement/internal/asset/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	r.GET("api/v1/users", controllers.GetUsers)          // دریافت تمامی کاربران
	r.POST("api/v1/users", controllers.CreateUser)       // ایجاد کاربر جدید
	r.GET("api/v1/users/:id", controllers.GetUserByID)   // دریافت اطلاعات کاربر بر اساس شناسه
	r.PUT("api/v1/users/:id", controllers.UpdateUser)    // بروزرسانی اطلاعات کاربر
	r.DELETE("api/v1/users/:id", controllers.DeleteUser) // حذف کاربر
}
