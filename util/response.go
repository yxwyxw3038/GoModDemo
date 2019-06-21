package util

import (
	"GoModDemo/consts"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}


func (g *Gin) Response(httpCode, errCode int, message string, data interface{}) {
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
