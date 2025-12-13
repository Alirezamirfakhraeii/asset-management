package routes

import (
	"_AssetManagement/internal/asset/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers)          // دریافت تمامی کاربران
	r.POST("/users", controllers.CreateUser)       // ایجاد کاربر جدید
	r.GET("/users/:id", controllers.GetUserByID)   // دریافت اطلاعات کاربر بر اساس شناسه
	r.PUT("/users/:id", controllers.UpdateUser)    // بروزرسانی اطلاعات کاربر
	r.DELETE("/users/:id", controllers.DeleteUser) // حذف کاربر
}
