package router

import (
	. "GoModDemo/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	return router
}
