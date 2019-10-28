package router

import (
	_ "GoModDemo/docs"
	"GoModDemo/middleware/jwt"
	"GoModDemo/service/api"

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
	userOne.POST("GetAllUserInfo", api.GetAllUserInfo)
	userOne.POST("GetAllUserViewInfo", api.GetAllUserViewInfo)
	userOne.POST("GetUserLogininfoByToken", api.GetUserLogininfoByToken)
	userOne.POST("AddUser", api.AddUser)
	userOne.POST("UpdateUser", api.UpdateUser)
	userOne.POST("DeleteUser", api.DeleteUser)
	userOne.POST("SetUserDept", api.SetUserDept)
	userOne.POST("SetUserRole", api.SetUserRole)
	menuOne := router.Group("/Menu")
	menuOne.POST("GetMenuByID", api.GetMenuByID)
	menuOne.POST("GetAllMenuInfo", api.GetAllMenuInfo)
	menuOne.POST("GetAllMenuViewInfo", api.GetAllMenuViewInfo)
	menuOne.POST("GetCascaderMenu", api.GetCascaderMenu)
	menuOne.POST("GetMenuAllCount", api.GetMenuAllCount)
	menuOne.POST("AddMenu", api.AddMenu)
	menuOne.POST("UpdateMenu", api.UpdateMenu)
	menuOne.POST("DeleteMenu", api.DeleteMenu)
	menuOne.POST("SetMenuButton", api.SetMenuButton)
	buttonOne := router.Group("/Button")
	buttonOne.POST("GetButtonByMenuIdAndUserId", api.GetButtonByMenuIdAndUserId)
	buttonOne.POST("GetAllButtonForTransfer", api.GetAllButtonForTransfer)
	buttonOne.POST("GetButtonByMenuIdnForTransfer", api.GetButtonByMenuIdnForTransfer)
	buttonOne.POST("GetAllButtonByMenuIdForTransfer", api.GetAllButtonByMenuIdForTransfer)
	buttonOne.POST("GetButtonByMenuIdRoleIdForTransfer", api.GetButtonByMenuIdRoleIdForTransfer)
	buttonOne.POST("AddButton", api.AddButton)
	buttonOne.POST("UpdateButton", api.UpdateButton)
	buttonOne.POST("DeleteButton", api.DeleteButton)
	buttonOne.POST("GetAllButtonInfo", api.GetAllButtonInfo)
	buttonOne.POST("GetButtonByID", api.GetButtonByID)

	deptOne := router.Group("/Dept")
	deptOne.POST("GetAllDeptForTransfer", api.GetAllDeptForTransfer)
	deptOne.POST("GetDeptByUserIdForTransfer", api.GetDeptByUserIdForTransfer)
	deptOne.POST("AddDept", api.AddDept)
	deptOne.POST("UpdateDept", api.UpdateDept)
	deptOne.POST("DeleteDept", api.DeleteDept)
	deptOne.POST("GetAllDeptInfo", api.GetAllDeptInfo)
	deptOne.POST("GetDeptByID", api.GetDeptByID)
	roleOne := router.Group("/Role")
	roleOne.POST("GetAllRoleForTransfer", api.GetAllRoleForTransfer)
	roleOne.POST("GetRoleByUserIdForTransfer", api.GetRoleByUserIdForTransfer)
	roleOne.POST("GetAllRoleInfo", api.GetAllRoleInfo)
	roleOne.POST("GetRoleByID", api.GetRoleByID)
	roleOne.POST("GetMenuByRoleIdForTree", api.GetMenuByRoleIdForTree)
	roleOne.POST("SetMenuRole", api.SetMenuRole)
	roleOne.POST("AddRole", api.AddRole)
	roleOne.POST("UpdateRole", api.UpdateRole)
	roleOne.POST("DeleteRole", api.DeleteRole)
	roleOne.POST("SetButtonByMenuIdRoleId", api.SetButtonByMenuIdRoleId)
	return router
}
