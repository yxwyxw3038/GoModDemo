package util

import (
	"GoModDemo/consts"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}


func (g *Gin) Response(httpCode, errCode int, message string, data interface{}) {

	g.C.Header("Access-Control-Allow-Origin", "*")		// 这是允许访问所有域
	g.C.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	g.C.JSON(httpCode, gin.H{
		"Code": consts.GetMsg(errCode),
		"Message": message ,
		"Data": data,
	})

	return
}
func (g *Gin) GetBase64Body() (string, error) {
	returnStr := ""
	buf := make([]byte, 1024)
	n, _ := g.C.Request.Body.Read(buf)
	passwordString := string(buf[0:n])
	decodeString,err:=UnBase64(passwordString)
	if err != nil {
		return returnStr, err
	}
	return decodeString, nil

}
