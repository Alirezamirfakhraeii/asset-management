package main

import (
	"assetManagement/internal/asset/routes"
	"assetManagement/internal/db"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectToDatabase()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	routes.RegisterCategoryRoutes(r)
	routes.SetupUserRoutes(r)
	routes.RegisterAuthRoutes(r)
	routes.RegisterAssetsRoutes(r)

	err := r.Run(":8000") // سرور را روی پورت 8000 راه‌اندازی می‌کند
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

}
