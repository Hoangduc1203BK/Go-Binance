package users

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	userRouter.POST("", CreateUser)
	userRouter.GET("/:id", ListUserByID)
	userRouter.PATCH("/:id", UpdateUserByID)
	userRouter.DELETE("/:id",DeleteUserByID)
}
