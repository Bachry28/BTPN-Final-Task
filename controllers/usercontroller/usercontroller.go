package usercontroller

import (
	"net/http"

	"github.com/Bachry28/BTPN-Final-Task/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUser(c *gin.Context) {
	var users []model.User

	if err := model.DB.Preload("Photos").Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserById(c *gin.Context) {
	var user model.User
	id := c.Param("id")

	if err := model.DB.Preload("Photos").First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Create Success!", "user": user})
}

func UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if model.DB.Model(&model.User{}).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't Update User!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Success!", "user": user})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if model.DB.Delete(&model.User{}, id).Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't Delete User!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Success!"})
}
