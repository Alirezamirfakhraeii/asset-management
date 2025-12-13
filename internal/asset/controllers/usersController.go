package controllers

import (
	"_AssetManagement/internal/asset/model"
	"_AssetManagement/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var users model.Users

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data", "error": err.Error()})
		return
	}

	if err := db.DB.Create(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
