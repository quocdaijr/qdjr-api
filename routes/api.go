package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Register(*gin.RouterGroup)
}

type ApiRoute struct{}

func (_ ApiRoute) Register(*gin.RouterGroup) {}
