package routes

import (
	"github.com/gin-gonic/gin"
	"qdjr-api/controllers"
	"qdjr-api/middlewares"
)

type ArticleRoute struct{}

var articleController = new(controllers.ArticleController)

func (_ ArticleRoute) Register(group *gin.RouterGroup) {
	groupArticle := group.Group("/articles", middlewares.AuthMiddleware{}.VerifiedToken())
	{
		groupArticle.GET("", articleController.List)
		groupArticle.POST("", articleController.Create)
		groupArticle.PUT(":id", articleController.Update)
	}
}
