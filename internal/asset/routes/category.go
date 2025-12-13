package routes

import (
	"_AssetManagement/internal/asset/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine) {
	r.GET("api/v1/categories", controllers.GetCategories)
	r.POST("api/v1/categories", controllers.CreateCategory)
	r.PUT("api/v1/categories/:id", controllers.UpdateCategory)
	r.DELETE("api/v1/categories/:id", controllers.DeleteCategory)
}
