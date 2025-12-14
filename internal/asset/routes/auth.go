package routes

import (
	"_AssetManagement/internal/asset/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("api/v1/login", controllers.Login)
	r.POST("api/v1/logout", controllers.Logout)
}
