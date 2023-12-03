package routes

import (
	"github.com/gin-gonic/gin"
	"qdjr-api/controllers"
	"qdjr-api/middlewares"
)

type UserRoute struct{}

var userController = new(controllers.UserController)

func (_ UserRoute) Register(group *gin.RouterGroup) {
	groupUser := group.Group("/users", middlewares.AuthMiddleware{}.VerifiedToken())
	{
		groupUser.GET("", userController.Search)
		groupUser.GET("/:id", userController.Detail)
	}
}
