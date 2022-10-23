package controller

import (
	"github.com/wijtra/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST Triages

func CreateTriages(c *gin.Context) {

	var triage entity.Triage
	var patient entity.Patient
	var disease entity.Disease
	var inpantientdepartment entity.InpantientDepartment

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 10 จะถูก bind เข้าตัวแปร Triage
	if err := c.ShouldBindJSON(&triage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา Patient ด้วย id
	if tx := entity.DB().Where("id = ?", triage.Patient_ID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// 12: ค้นหา Disease ด้วย id
	if tx := entity.DB().Where("id = ?", triage.Disease_ID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}

	//12.1: ค้นหา InpantientDepartment ด้วย id
	if tx := entity.DB().Where("id = ?", triage.InpantientDepartment_ID).First(&inpantientdepartment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "inpantientdepartment not found"})
		return
	}

	// 13: ค้นหา user ด้วย id
	// if tx := entity.DB().Where("id = ?", trige.User_ID).First(&user); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
	// 	return
	// }

	// 14: สร้าง Trige
	mb := entity.Triage{
		Patient_ID:              triage.Patient_ID,              // โยงความสัมพันธ์กับ Entity Patient
		Disease_ID:              triage.Disease_ID,              // โยงความสัมพันธ์กับ Entity Disease
		InpantientDepartment_ID: triage.InpantientDepartment_ID, // โยงความสัมพันธ์กับ Entity InpantientDepartment
		Triage_COMMENT:          triage.Triage_COMMENT,
		//User_ID: triage.User_ID,               // โยงความสัมพันธ์กับ Entity User
	}

	// 15: บันทึก
	if err := entity.DB().Create(&mb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mb})

}

// GET /Triage/:id
func GetTriage(c *gin.Context) {
	var GetTriage entity.Triage
	id := c.Param("id")
	if err := entity.DB().Preload("Patient.NAME").Preload("Patient.Gender").Preload("Patient.Blood_type").Preload("Disease.DiseaseType").Raw("SELECT * FROM triage WHERE id = ?", id).Scan(&GetTriage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": GetTriage})
}

// GET /triage
// return ไปให้เพื่อน
func GetListTriages(c *gin.Context) {
	var GetTriages []entity.Triage
	if err := entity.DB().Preload("Patient.NAME").Preload("Patient.Gender").Preload("Disease.DiseaseType").Raw("SELECT * FROM triage ").Find(&GetTriages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": GetTriages})
}
