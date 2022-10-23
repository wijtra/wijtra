package controller

import (
	"github.com/wijtra/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST/Blood_types สำหรับ สร้างข้อมูล
func CreateBlood_type(c *gin.Context) {
	var Blood_type entity.Blood_type
	if err := c.ShouldBindJSON(&Blood_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}

// GET/Blood_types/:id
func GetBlood_type(c *gin.Context) {
	var Blood_type entity.Blood_type

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Blood_types WHERE id = ?", id).Find(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}

// GET /Blood_types สำหรับเรียกดูข้อมูลทั้งหมด เอาไปใช้กับ combobox ใน fontend
func ListBlood_type(c *gin.Context) {
	var Blood_type []entity.Blood_type
	if err := entity.DB().Raw("SELECT * FROM Blood_types").Find(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}
