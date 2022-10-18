package controller

import (
	"github.com/wijtra/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST InpantientDepartment

func CreateInpantientDepartment(c *gin.Context) {

	var inpantientdepartment entity.InpantientDepartment
	if err := c.ShouldBindJSON(&inpantientdepartment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&inpantientdepartment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": inpantientdepartment})

}

// GET /InpantientDepartment/:id
func GetInpantientDepartment(c *gin.Context) {
	var inpantientdepartments entity.InpantientDepartment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM inpantientdepartments WHERE id = ?", id).Scan(&inpantientdepartments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inpantientdepartments})
}

// GET /inpantientdepartments
func ListInpantientDepartment(c *gin.Context) {
	var inpantientdepartments []entity.InpantientDepartment
	if err := entity.DB().Raw("SELECT * FROM inpantientdepartments").Scan(&inpantientdepartments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inpantientdepartments})
}

// PATCH /inpantientdepartment
func UpdateInpantientDepartment(c *gin.Context) {
	var inpantientdepartment entity.InpantientDepartment
	if err := c.ShouldBindJSON(&inpantientdepartment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", inpantientdepartment.ID).First(&inpantientdepartment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "zone not found"})
		return
	}

	if err := entity.DB().Save(&inpantientdepartment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inpantientdepartment})
}
