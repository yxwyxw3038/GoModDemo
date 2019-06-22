package jwt

import (
  "GoModDemo/consts"
  "GoModDemo/util"
  "github.com/gin-gonic/gin"
  "net/http"
  "time"
)

func Jwt() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code int
        msg :=""
        // var data interface{}
        appG := util.Gin{C: c}
        code = consts.SUCCESS
        token := c.Query("token")
        if token == "" {
            code = consts.ERROR
        } else {
            claims, err := util.ParseToken(token)
            if err != nil {
                code = consts.ERROR
                msg ="验证不通过"
            } else if time.Now().Unix() > claims.ExpiresAt {
                code = consts.ERROR
                msg ="验证超时"
            }
        }

        if code != consts.SUCCESS {
            // c.JSON(http.StatusUnauthorized, gin.H{
            //     "code": code,
            //     "msg":  consts.GetMsg(code),
            //     "data": data,
            // })

            // c.Abort()
            appG.Response(http.StatusUnauthorized, code,msg, nil)
            appG.C.Abort()
            return
        }

        c.Next()
    }
}