package controller

// import (
// 	"github.com/wijtra/sa-65-example/entity"

// 	"github.com/gin-gonic/gin"

// 	"net/http"
// )

// // POST User

// func CreateUser(c *gin.Context) {

// 	var user entity.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := entity.DB().Create(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": user})

// }

// // GET /user/:id
// func GetUser(c *gin.Context) {
// 	var user entity.Triage
// 	id := c.Param("id")
// 	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

// // GET /users
// func ListUsers(c *gin.Context) {
// 	var users []entity.Triage
// 	if err := entity.DB().Raw("SELECT * FROM users").Scan(&users).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": users})
// }
