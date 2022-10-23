package controller

import (
	"github.com/wijtra/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller User    		    ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateUser(c *gin.Context) {

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /user/:id
func GetUser(c *gin.Context) {
	var user entity.User
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /usrs
func ListUsers(c *gin.Context) {
	var users []entity.User
	if err := entity.DB().Raw("SELECT * FROM users").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller Gender    		    ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller Blood_type    		    ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller Patient		    ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// POST Patient
func CreatePatient(c *gin.Context) {

	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET /getPatient/:id
func GetPatient(c *gin.Context) {
	var GetPatient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&GetPatient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetPatient})
}

// GET /GetListPatients
func GetListPatients(c *gin.Context) {
	var getPatients []entity.Patient
	if err := entity.DB().Preload("Gender").Raw("SELECT * FROM patients").Find(&getPatients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": getPatients})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller DiseaseType		    ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
	if err := entity.DB().Raw("SELECT * FROM disease_types").Scan(&diseasetypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diseasetypes})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller Disease		    ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller InpantientDepartment	    ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller Triage	    ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

	// 15: บันทึก Triage
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
