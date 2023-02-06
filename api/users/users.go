package users

import (
	"binance/database"
	"binance/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var body struct {
		Email       string
		Name        string
		Password    string
		PhoneNumber string
	}
	c.Bind(&body)
	user := model.User{
		Email:       body.Email,
		Name:        body.Name,
		Password:    body.Password,
		PhoneNumber: body.PhoneNumber}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": user,
	})
}

func ListUserByID(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	database.DB.First(&user, id)
	c.JSON(200, gin.H{
		"post": user,
	})
}

func UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Email       string
		Name        string
		Password    string
		PhoneNumber string
	}
	c.Bind(&body)
	inputUser := model.User{
		Email:       body.Email,
		Name:        body.Name,
		Password:    body.Password,
		PhoneNumber: body.PhoneNumber}

	var user model.User

	database.DB.First(&user, id)
	if user.Email == "" || user.Name == "" || user.Password == "" || user.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}
	database.DB.First(&user, id).Updates(inputUser)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")

	var user model.User
	database.DB.First(&user, id)
	if user.Email == "" || user.Name == "" || user.Password == "" || user.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	database.DB.Delete(&model.User{}, id)

	c.JSON(200, gin.H{
		"post": "deleted Post ID: " + id,
		// "Deleted User": user,
	})
}
