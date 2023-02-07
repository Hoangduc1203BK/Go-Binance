package users

import (
	"binance/database"
	"binance/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func RepositoryListUserByID(id string) model.User {
	var user model.User
	database.DB.First(&user, id)
	return user
}

func RepositoryUpdateUser(req *updateUserRequest, user model.User, id string, c *gin.Context) (bool, model.User) {

	inputUser := model.User{
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
