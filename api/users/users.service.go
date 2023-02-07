package users

import (
	"binance/database"
	"binance/model"
	"binance/util"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceCreateUser(req *createUserRequest, c *gin.Context) createUserResponse {

	fmt.Println("ServiceCreateUser req ", req)
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return createUserResponse{}
	}
	rsq, err := RepositoryCreateUser(req, hashPassword)
	if err != nil {
		c.JSON(http.StatusBadGateway, errorResponse(err))
		return createUserResponse{}
	}
	return rsq
}

func ServiceListUserByID(c *gin.Context, id string) model.User {
	user := RepositoryListUserByID(id)
	return user
}

func ServiceUpdateUserByID(c *gin.Context, id string, req *updateUserRequest) (model.User, error) {
	var user model.User

	shouldReturn, returnValue := RepositoryUpdateUser(req, user, id, c)
	if !shouldReturn {
		err := errors.New("Cannot insert to DB")
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return model.User{}, err
	}

	return returnValue, nil
}

func ServiceDeleteUserByID(c *gin.Context) {
	id := c.Param("id")

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
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
