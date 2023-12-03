package routes

import (
	"github.com/gin-gonic/gin"
	"qdjr-api/controllers"
)

type AuthRoute struct {
	Route
}

var authController = new(controllers.AuthController)

func (_ AuthRoute) Register(group *gin.RouterGroup) {
	groupAuth := group.Group("/auth")
	{
		groupAuth.POST("/register", authController.Register)
		groupAuth.POST("/login", authController.Login)
		groupAuth.POST("/refresh", authController.RefreshToken)
	}
}
