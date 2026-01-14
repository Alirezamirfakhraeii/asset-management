package controllers

import (
	"assetManagement/internal/asset/model"
	"assetManagement/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAssets(c *gin.Context) {
	var assets []model.Asset
	if err := db.DB.Find(&assets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching assets"})
		return
	}
	c.JSON(http.StatusOK, assets)
}

func CreateAsset(c *gin.Context) {
	var asset model.Asset

	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data", "error": err.Error()})
		return
	}

	if err := db.DB.Create(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating asset"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "creating asset"})

}

func UpdateAsset(c *gin.Context) {
	var asset model.Asset
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&asset).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Asset not found"})
		return
	}

	if err := db.DB.Save(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating asset"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updating asset"})

}

func DeleteAsset(c *gin.Context) {
	var asset model.Asset
	id := c.Param("id")
	if err := db.DB.Where("id = ?", id).First(&asset).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Asset not found"})
		return
	}
	if err := db.DB.Delete(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting asset"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleting asset"})
}
