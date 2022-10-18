package controller

import (
	"github.com/wijtra/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST DiseaseType

func CreateDiseaseType(c *gin.Context) {

	var diseasetype entity.DiseaseType
	if err := c.ShouldBindJSON(&diseasetype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&diseasetype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diseasetype})

}

// GET /DiseaseType/:id
func GetDiseasType(c *gin.Context) {
	var diseasetype entity.DiseaseType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM DiseasTypes WHERE id = ?", id).Scan(&diseasetype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diseasetype})
}

// GET /DiseaseTypes
func ListDiseaseTypes(c *gin.Context) {
	var diseasetypes []entity.DiseaseType
	if err := entity.DB().Raw("SELECT * FROM DiseasTypes").Scan(&diseasetypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diseasetypes})
}

// PATCH /DiseaseTypes
func UpdateDiseaseType(c *gin.Context) {
	var diseasetype entity.DiseaseType
	if err := c.ShouldBindJSON(&diseasetype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", diseasetype.ID).First(&diseasetype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DiseasType not found"})
		return
	}

	if err := entity.DB().Save(&diseasetype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diseasetype})
}
