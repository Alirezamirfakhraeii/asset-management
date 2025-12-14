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

func GetUsers(c *gin.Context) {
	var users []model.Users

	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	var user model.Users
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user model.Users
	id := c.Param("id")
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting user"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data", "error": err.Error()})
		return
	}
	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated!"})
}

func DeleteUser(c *gin.Context) {
	var user model.Users
	id := c.Param("id")
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting user"})
		return
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
