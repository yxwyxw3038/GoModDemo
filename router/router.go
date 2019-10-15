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
	userOne.POST("GetAllUserViewInfo",api.GetAllUserViewInfo)
	userOne.POST("GetUserLogininfoByToken",api.GetUserLogininfoByToken)
	userOne.POST("AddUser",api.AddUser)
	userOne.POST("UpdateUser",api.UpdateUser)
	userOne.POST("DeleteUser",api.DeleteUser)
	userOne.POST("SetUserDept",api.SetUserDept)
	userOne.POST("SetUserRole",api.SetUserRole)
	menuOne := router.Group("/Menu")
	menuOne.POST("GetMenuByID", api.GetMenuByID)
	menuOne.POST("GetAllMenuInfo",api.GetAllMenuInfo)
	menuOne.POST("GetAllMenuViewInfo",api.GetAllMenuViewInfo)
	menuOne.POST("GetCascaderMenu",api.GetCascaderMenu)
	menuOne.POST("GetMenuAllCount",api.GetMenuAllCount)
	menuOne.POST("AddMenu",api.AddMenu)
	menuOne.POST("UpdateMenu",api.UpdateMenu)
	menuOne.POST("DeleteMenu",api.DeleteMenu)
	buttonOne := router.Group("/Button")
	buttonOne.POST("GetButtonByMenuIdAndUserId", api.GetButtonByMenuIdAndUserId)
	deptOne := router.Group("/Dept")
	deptOne.POST("GetAllDeptForTransfer", api.GetAllDeptForTransfer)
	deptOne.POST("GetDeptByUserIdForTransfer", api.GetDeptByUserIdForTransfer)
	roleOne := router.Group("/Role")
	roleOne.POST("GetAllRoleForTransfer", api.GetAllRoleForTransfer)
	roleOne.POST("GetRoleByUserIdForTransfer", api.GetRoleByUserIdForTransfer)
	return router
}
