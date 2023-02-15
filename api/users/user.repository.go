package users

import (
	constance "binance/const"
	"binance/database"
	"binance/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RepositoryListUser(payload *listUserDto) []model.User {
	var params listUserDto
	if payload.Page > 0 {
		params.Page = payload.Page
	} else {
		params.Page = constance.DefaultPaging().Page
	}

	if payload.Limit > 0 {
		params.Limit = payload.Limit
	} else {
		params.Limit = constance.DefaultPaging().Limit
	}

	offset := (params.Page - 1) * params.Limit

	var users []model.User
	database.DB.Limit(params.Limit).Offset(offset).Find(&users)

	return users
}

func RepositoryGetUserByAuth(email *string) model.User {
	var user model.User
	database.DB.Where("email = ?", email).First(&user)

	return user
}

func RepositoryCreateUser(req *createUserRequest, hashPassword string) (createUserResponse, error) {
	user := model.User{
		Email:       req.Email,
		Name:        req.Name,
		Password:    hashPassword,
		PhoneNumber: req.PhoneNumber}

	result := database.DB.Create(&user)
	if result.Error != nil {
		return createUserResponse{}, result.Error
	}

	rsq := createUserResponse{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}
	return rsq, nil
}

func RepositoryGetUserByID(c *gin.Context, id uint) model.User {
	var user model.User
	database.DB.First(&user, id)
	return user
}

func RepositoryUpdateUser(req *updateUserRequest, user model.User, id uint, c *gin.Context) (bool, model.User) {

	inputUser := model.User{
		Id:          id,
		Email:       req.Email,
		Name:        req.Name,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber}

	fmt.Println("RepositoryUpdateUser >>>>>>>>", inputUser)
	database.DB.First(&user, id)
	if user.Email == "" || user.Name == "" || user.Password == "" || user.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"system message": "User not found",
		})
		return false, model.User{}
	}
	database.DB.First(&user, id).Updates(inputUser)
	return true, inputUser
}

func RepositoryDeleteUser(id string, c *gin.Context) {
	var user model.User
	database.DB.First(&user, id)
	if user.Email == "" || user.Name == "" || user.Password == "" || user.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

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
	return
}
