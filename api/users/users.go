package users

import (
	"binance/database"
	"binance/model"
	"binance/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name        string `json:"name" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number" binding:"required,number,min=10"`
	Email       string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type updateUserRequest struct {
	createUserRequest
}

type GetUserDto struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	user := model.User{
		Email:       req.Email,
		Name:        req.Name,
		Password:    hashPassword,
		PhoneNumber: req.PhoneNumber}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, errorResponse(result.Error))
		return
	}

	rsq := createUserResponse{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}
	c.JSON(200, gin.H{
		"post": rsq,
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
	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	inputUser := model.User{
		Email:       req.Email,
		Name:        req.Name,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber}

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

	// var user model.User
	// database.DB.First(&user, id)
	// if user.Email == "" || user.Name == "" || user.Password == "" || user.PhoneNumber == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "User not found",
	// 	})
	// 	return
	// }

	result := database.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid User ID",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": "deleted Post ID: " + id,
	})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
