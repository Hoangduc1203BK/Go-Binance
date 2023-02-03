package users

import (
	"binance/database"
	"binance/model"

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

func ListUser(c *gin.Context) {
	id := c.Param("id")
	// get the post
	var user model.User
	database.DB.First(&user, id)
	//respond with them

	c.JSON(200, gin.H{
		"post": user,
	})
}
