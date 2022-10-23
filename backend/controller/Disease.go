package controller

import (
	"github.com/wijtra/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST Disease

func CreateDisease(c *gin.Context) {

	var disease entity.Disease

	if err := c.ShouldBindJSON(&disease); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&disease).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disease})

}

// GET /Disease/:id

func GetDisease(c *gin.Context) {

	var disease entity.Disease

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&disease).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disease})

}

// GET /Disease/:DiseaseTypeid
func GetDisease_by_DiseaseType(c *gin.Context) {
	var disease_by_diseasetype []entity.Disease
	diseasetype_id := c.Param("DiseaseTypeid")
	if err := entity.DB().Raw("SELECT * FROM beds WHERE DiseaseType_ID = ?", diseasetype_id).Scan(&disease_by_diseasetype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": disease_by_diseasetype})
}

// GET /Diseases

func ListDisease(c *gin.Context) {

	var diseases []entity.Disease

	if err := entity.DB().Raw("SELECT * FROM diseases").Find(&diseases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": diseases})

}
