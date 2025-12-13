package controllers

import (
	"_AssetManagement/internal/asset/model"
	"_AssetManagement/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET

func GetCategories(c *gin.Context) {
	var categories []model.Category

	if err := db.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

//CREATE

func CreateCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data", "error": err.Error()})
		return
	}

	if err := db.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created"})
}

//UPDATE

func UpdateCategory(c *gin.Context) {
	var category model.Category

	// گرفتن شناسه از پارامتر URL
	categoryId := c.Param("id")

	// پیدا کردن دسته‌بندی با شناسه مشخص
	if err := db.DB.Where("id = ?", categoryId).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	// دریافت داده‌های جدید از بدنۀ درخواست
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data", "error": err.Error()})
		return
	}

	// بروزرسانی رکورد
	if err := db.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated"})
}

//DELETE

func DeleteCategory(c *gin.Context) {
	var category model.Category
	categoryId := c.Param("id")

	// پیدا کردن رکورد بر اساس شناسه
	if err := db.DB.Where("id = ?", categoryId).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	// حذف رکورد
	if err := db.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
