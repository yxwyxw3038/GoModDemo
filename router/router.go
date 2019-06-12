package router

import (
	"GoModDemo/service/api"
	"GoModDemo/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	loginVersionOne := router.Group("/")
	loginVersionOne.GET("Login", api.Login)
	apiVersionOne := router.Group("/api")
	apiVersionOne.Use(jwt.Jwt())
	apiVersionOne.GET("hello", api.IndexApi)
	return router
}
