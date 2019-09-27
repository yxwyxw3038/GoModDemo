package router

import (
	"GoModDemo/middleware/jwt"
	"GoModDemo/service/api"
	_ "GoModDemo/docs"
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
	testVersionOne := router.Group("/test")
	testVersionOne.GET("hello", api.IndexApi)
	userOne := router.Group("/User")
	userOne.Use(jwt.Jwt())
	userOne.POST("GetUserByID", api.GetUserByID)
	userOne.POST("GetUserMenu", api.GetUserMenu)
	return router
}
