package auth

import (
	"binance/api/users"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	var body users.GetUserDto
	c.Bind(&body)

	result, err := LogInService(&body)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	maxAge := time.Hour * 24 * 30
	c.SetCookie("refreshToken", result.refreshToken, int(maxAge), "", "", true, true)

	c.JSON(200, gin.H{
		"accessToken":  result.accessToken,
		"refreshToken": result.refreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	fmt.Println(c.Keys)
	refreshToken, _ := c.Cookie("refreshToken")
	result, err := RefreshTokenService(&refreshToken)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	maxAge := time.Hour * 24 * 30
	c.SetCookie("refreshToken", result.refreshToken, int(maxAge), "", "", true, true)

	c.JSON(200, gin.H{
		"accessToken":  result.accessToken,
		"refreshToken": result.refreshToken,
	})
}

func LogOut(c *gin.Context) {
	refreshToken, _ := c.Cookie("refreshToken")

	err := LogOutService(&refreshToken)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("refreshToken", "", -1, "", "", true, true)

	c.JSON(200, gin.H{
		"success": true,
	})
}

