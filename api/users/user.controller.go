package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ControllerListUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	var params listUserDto
	params.Page = page
	params.Limit = limit

	users := ServiceListUser(&params)

	c.JSON(200, gin.H{
		"result": users,
	})

}

func ControllerCreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result := ServiceCreateUser(&req, c)
	if result.Email == "" {
		c.JSON(500, gin.H{"message": "cannot create a User as incomming creatorial"})
		return
	}
	c.JSON(200, gin.H{
		"user": result,
	})
	return
}

func ControllerGetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ServiceGetUserByID(c, uint(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
	return
}

func ControllerUpdateUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	parseId := uint(id)
	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := ServiceUpdateUserByID(c, parseId, &req)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
	return
}

func ControllerDeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	ServiceDeleteUserByID(c, id)

	return
}
