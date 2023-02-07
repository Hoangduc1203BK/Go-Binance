package users

import (
	"binance/model"
	"binance/util"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func ServiceListUserByID(c *gin.Context, id string) model.User {
	user := RepositoryListUserByID(c, id)
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

func ServiceDeleteUserByID(c *gin.Context, id string) {
	RepositoryDeleteUser(id, c)
	return
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
