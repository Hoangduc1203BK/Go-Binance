package auth

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	// authRouter.POST("/login", LogIn)
	authRouter.GET("/refresh-token", RefreshToken)
	authRouter.GET("/logout", LogOut)
}
