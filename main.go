package main

import (
	. "GoModDemo/router"
	"GoModDemo/util"
)

func main() {
	logger := util.InitZapLog()
	logger.Debug("启动服务")
	router := InitRouter()
	router.Run(":8080")
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
