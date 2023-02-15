package users

import (
	"binance/model"
	"binance/util"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceListUser(payload *listUserDto) []model.User {
	users := RepositoryListUser(payload)

	return users
}

func ServiceCreateUser(req *createUserRequest, c *gin.Context) createUserResponse {

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

func ServiceGetUserByAuth(data *GetUserDto) (model.User, error) {
	result := RepositoryGetUserByAuth(&data.Email)

	if result.Id == 0 || result.DeletedAt.Valid == true {
		return result, fmt.Errorf("User not found")
	}

	err2 := util.CheckPassword(data.Password, result.Password)

	if err2 != nil {
		return result, fmt.Errorf("Invalid password")
	}

	return result, nil
}

func ServiceGetUserByID(c *gin.Context, userId uint) (interface{}, error) {
	user := RepositoryGetUserByID(c, userId)

	if user.Id == 0 || user.DeletedAt.Valid == true {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

func ServiceUpdateUserByID(c *gin.Context, id uint, req *updateUserRequest) (model.User, error) {
	var user model.User

	shouldReturn, returnValue := RepositoryUpdateUser(req, user, id, c)
	if !shouldReturn {
		err := errors.New("Cannot insert to DB")
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return model.User{}, err
	}

	return returnValue, nil
}

func ServiceDeleteUserByID(c *gin.Context, id string) {
	RepositoryDeleteUser(id, c)
	return
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
