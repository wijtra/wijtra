package controller

import (
	"github.com/wijtra/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST inpantientdepartment
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

// GET /getinpantientdepartment/:id
func GetInpantientDepartment(c *gin.Context) {
	var Getinpantientdepartment entity.InpantientDepartment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM inpantient_departments WHERE id = ?", id).Scan(&Getinpantientdepartment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Getinpantientdepartment})
}

// GET /GetListIPDs
func GetListInpantientDepartments(c *gin.Context) {
	var getinpantientdepartments []entity.InpantientDepartment
	if err := entity.DB().Raw("SELECT * FROM inpantient_departments").Scan(&getinpantientdepartments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": getinpantientdepartments})
}
