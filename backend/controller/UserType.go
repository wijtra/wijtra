package controller

// import (
// 	"github.com/wijtra/sa-65-example/entity"

// 	"github.com/gin-gonic/gin"

// 	"net/http"
// )

// // POST User_Type

// func CreateUserType(c *gin.Context) {

// 	var usertype entity.UserType
// 	if err := c.ShouldBindJSON(&usertype); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := entity.DB().Create(&usertype).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": usertype})

// }

// // GET /user_type/:id
// func GetUserType(c *gin.Context) {
// 	var usertype entity.UserType
// 	id := c.Param("id")
// 	if err := entity.DB().Raw("SELECT * FROM user_types WHERE id = ?", id).Scan(&usertype).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": usertype})
// }

// // GET /user_types
// func ListUserTypes(c *gin.Context) {
// 	var usertypes []entity.UserType
// 	if err := entity.DB().Raw("SELECT * FROM user_types").Scan(&usertypes).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": usertypes})
// }
