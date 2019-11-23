package util

import (
	"GoModDemo/consts"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, message string, data interface{}) {

	g.C.Header("Access-Control-Allow-Origin", "*") // 这是允许访问所有域
	g.C.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	// g.C.Header("Access-Control-Allow-Credentials", "true")
	g.C.Header("Access-Control-Allow-Headers", "Content-Type, Accept")
	g.C.Header("Access-Control-Max-Age", "1728000")
	g.C.JSON(httpCode, gin.H{
		"Code":    consts.GetMsg(errCode),
		"Message": message,
		"Data":    data,
	})

	return
}
func (g *Gin) Response1(httpCode, errCode int, message string, data interface{}, dataCount int) {

	g.C.Header("Access-Control-Allow-Origin", "*") // 这是允许访问所有域
	g.C.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	// g.C.Header("Access-Control-Allow-Credentials", "true")
	g.C.Header("Access-Control-Allow-Headers", "Content-Type, Accept")
	g.C.Header("Access-Control-Max-Age", "1728000")
	g.C.JSON(httpCode, gin.H{
		"Code":      consts.GetMsg(errCode),
		"Message":   message,
		"Data":      data,
		"DataCount": dataCount,
	})

	return
}

func (g *Gin) GetBase64Body() (string, error) {
	returnStr := ""
	count := 1024 * 100
	buf := make([]byte, count)
	n, _ := g.C.Request.Body.Read(buf)
	passwordString := string(buf[0:n])
	fmt.Println(passwordString)
	decodeString, err := UnBase64(passwordString)
	if err != nil {
		return returnStr, err
	}
	return decodeString, nil

}
func (g *Gin) ParseQuery() (map[string][]string, error) {
	str, err := g.GetBase64Body()
	if err != nil {

		return nil, err
	}
	dists, err := url.ParseQuery(str)
	if err != nil {

		return nil, err
	}
	return dists, nil
}
