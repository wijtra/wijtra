package controller

import (
	"github.com/wijtra/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /สำหรับสร้างข้อมูล
func CreateGender(c *gin.Context) {
	var Gender entity.Gender
	if err := c.ShouldBindJSON(&Gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Gender})
}

// GET / แบบเฉพาะเจาะจง
func GetGender(c *gin.Context) {
	var Gender entity.Gender

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Genders WHERE id = ?", id).Find(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Gender})
}

// GET /
func ListGender(c *gin.Context) {
	var Gender []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM Genders").Find(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Gender})
}
