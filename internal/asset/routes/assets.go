package routes

import (
	"_AssetManagement/internal/asset/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAssetsRoutes(r *gin.Engine) {
	r.GET("/api/v1/assets", controllers.GetAssets)
	r.POST("/api/v1/assets", controllers.CreateAsset)
	r.PUT("/api/v1/assets/:id", controllers.UpdateAsset)
	r.DELETE("/api/v1/assets/:id", controllers.DeleteAsset)
}
