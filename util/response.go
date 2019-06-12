package util

import (
	"GoModDemo/consts"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  consts.GetMsg(errCode),
		"data": data,
	})

	return
}
func (g *Gin) GetBase64Body() (string, error) {
	returnStr := ""
	buf := make([]byte, 1024)
	n, _ := g.C.Request.Body.Read(buf)
	passwordString := string(buf[0:n])
	// l := len(passwordString)
	// bufstr := make([]byte, l)
	bufstr1 := append([]byte(passwordString)[1:], []byte(passwordString)[0])
	mwString := string(bufstr1)
	// bufstr[l-1] = passwordString[0]
	// for i := 0; i < (l - 1); i++ {
	// 	bufstr[i] = passwordString[i+1]

	// }
	// mwString := string(buf)
	// s0 := string([]byte(passwordString)[:1])
	// s1 := string([]byte(passwordString)[1 : l-1])
	// mwString := s1 + s0

	decodeBytes, err := base64.StdEncoding.DecodeString(mwString)
	decodeString := string(decodeBytes)
	if err != nil {
		return returnStr, err
	}

	return decodeString, nil

}
