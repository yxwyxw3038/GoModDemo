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

	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()
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
	userOne.POST("GetAllUserInfo",api.GetAllUserInfo)
	userOne.POST("GetUserLogininfoByToken",api.GetUserLogininfoByToken)
	menuOne := router.Group("/Menu")
	menuOne.POST("GetMenuByID", api.GetMenuByID)
	menuOne.POST("GetAllMenuInfo",api.GetAllMenuInfo)
	menuOne.POST("GetCascaderMenu",api.GetCascaderMenu)
	buttonOne := router.Group("/Button")
	buttonOne.POST("GetButtonByMenuIdAndUserId", api.GetButtonByMenuIdAndUserId)

	return router
}
