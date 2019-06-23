package main

import (
	"GoModDemo/daemon"
	. "GoModDemo/router"
	"GoModDemo/setting"
	"GoModDemo/util"
)

// @title Rest API
// @version 1.0
// @description 目前仅仅是一个demo
func main() {
	setting.Setup()
	util.RegisterInfo()
	go daemon.Run()
	logger := util.InitZapLog()
	url := ":" + setting.ServerSetting.Port
	logger.Debug("启动服务" + url)
	router := InitRouter()
	router.Run(url)

	// address := fmt.Sprintf("%s:%s", setting.ServerSetting.Ip, setting.ServerSetting.Port)
	//   server := endless.NewServer(address, r)
	//   server.BeforeBegin = func(add string) {
	// 	  log.Printf("Actual pid is %d", syscall.Getpid())
	//   }

	// err := server.ListenAndServe()
	//   if err != nil {
	// 	  log.Printf("Server err: %v", err)
	//   }
}
