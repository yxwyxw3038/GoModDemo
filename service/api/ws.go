package api

import (
	"GoModDemo/consts"
	"GoModDemo/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsPage(c *gin.Context) {
	appG := util.Gin{C: c}
	// change the reqest to websocket model
	conn, error := upGrader.Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		appG.Response(http.StatusOK, consts.ERROR, "websocket异常", nil)
		return
	}
	newUuid := uuid.New()
	newUuidStr := newUuid.String()
	// websocket connect
	client := &util.WSClient{ID: newUuidStr, Socket: conn, Send: make(chan []byte)}
	util.WSManager.Register <- client

	go client.Read()
	go client.Write()
	// go test(client)
	s := string(newUuidStr)
	// appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
	//     "token": token,
	// })
	appG.Response(http.StatusOK, consts.SUCCESS, "WS联接成功", s)
}
