package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



func ControllerCreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("ControllerCreateUser req ", req)
	result := ServiceCreateUser(&req, c)
	c.JSON(200, gin.H{
		"user": result,
	})
	return
}

func ControllerListUserByID(c *gin.Context) {
	id := c.Param("id")
	user := ServiceListUserByID(c, id)
	c.JSON(200, gin.H{
		"user": user,
	})
	return
}

func ControllerUpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	user, err := ServiceUpdateUserByID(c, id, &req)
	if  err!= nil  {
		return 
	}
	c.JSON(200, gin.H{
		"user": user,
	})
	return
}

func ControllerDeleteUserByID(c *gin.Context) {

}
