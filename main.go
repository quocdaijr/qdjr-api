package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"qdjr-api/initializers"
	"qdjr-api/routes"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MainDbInitializer := new(initializers.DbInitializer)
	MainDbInitializer.ConnectDataBase()

	MainRedisInitializer := new(initializers.RedisInitializer)
	MainRedisInitializer.ConnectRedis()

	gin.SetMode(gin.DebugMode)

	router := gin.Default() //new gin router initialization

	group := router.Group("/api")
	/**
	 * Auth Routes
	 */
	authRoute := new(routes.AuthRoute)
	authRoute.Register(group)

	/**
	 * User Routes
	 */
	userRoute := new(routes.UserRoute)
	userRoute.Register(group)

	/**
	 * Article Routes
	 */
	articleRoute := new(routes.ArticleRoute)
	articleRoute.Register(group)

	err = router.Run(":8000")
	if err != nil {
		return
	} //running application, Default port is 8080
}
