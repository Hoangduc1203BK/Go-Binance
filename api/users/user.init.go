package users

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	userRouter.POST("/", ControllerCreateUser)
	userRouter.GET("/:id", ControllerListUserByID)
	userRouter.PATCH("/:id", ControllerUpdateUserByID)
	userRouter.DELETE("/:id", ControllerDeleteUserByID)
}
