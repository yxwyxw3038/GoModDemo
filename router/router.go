package router

import (
	"GoModDemo/middleware/jwt"
	"GoModDemo/service/api"

	"github.com/gin-gonic/gin"
	 ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)
//InitRouter 注册路由
func InitRouter() *gin.Engine {

	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	loginVersionOne := router.Group("/")
	loginVersionOne.POST("Login", api.Login)
	apiVersionOne := router.Group("/api")
	apiVersionOne.Use(jwt.Jwt())
	apiVersionOne.GET("hello", api.IndexApi)
	return router
}
