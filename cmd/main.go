package main

import (
	"_AssetManagement/internal/asset/routes"
	"_AssetManagement/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectToDatabase()
	r := gin.Default()
	routes.RegisterCategoryRoutes(r)
	routes.SetupUserRoutes(r)
	routes.RegisterAuthRoutes(r)
	routes.RegisterAssetsRoutes(r)
	err := r.Run(":8000") // سرور را روی پورت 8000 راه‌اندازی می‌کند
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

}
