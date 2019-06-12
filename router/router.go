package router

import (
	"GoModDemo/middleware/jwt"
	"GoModDemo/service/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	loginVersionOne := router.Group("/")
	loginVersionOne.POST("Login", api.Login)
	apiVersionOne := router.Group("/api")
	apiVersionOne.Use(jwt.Jwt())
	apiVersionOne.GET("hello", api.IndexApi)
	return router
}
