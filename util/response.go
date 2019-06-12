package util

import (
    "GoModDemo/consts"
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