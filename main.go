package main

import (
	"GoModDemo/daemon"
	. "GoModDemo/router"
	"GoModDemo/setting"
	"GoModDemo/util"
	"fmt"
)

const NOTE_IMG = `
┏━┓　　　┏━┓　　　　   ┏━━┓┏━━┓┏┓   ┏┳┓┏┓　　┏━┓
┃┃┃┏━┓┃━┫┏━━┓   ┃┏┓┃┃━━┃┃┃   ┃┃┃┃┃　　┃┃┃
┃　┫┃┻┫┣━┃┗┓┏┛   ┃┏┓┃┃┏━┛┃┃   ┃┃┃┃┃┏┓┃┃┃
┗┻┛┗━┛┗━┛　┗┛　   ┗┛┗┛┗┛　　┗┛   ┗━┛┗┛┗┛┗━┛`
const URL_IMG = `
　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　┏┓　　　　　
┏┳┳┳┳┳┳┳┳┓┏━━┳┳┳┳┳┳━┳┳┳┳━┳━┳━┳┛┃┏━┳━┓
┃┃┃┃┃┃┃┃┃┣┫┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┣┫┣┫┃┃
┗━━┻━━┻━━┻┻┻┻╋━┣━━┻━┻━━╋┓┣━┻━┻━┻┻━┻┻┛
　　　　　　　　　　　　　┗━┛　　　　　　　┗━┛　　　　　　　　　　　　　　　　　　　　　
`

// @title Rest API
// @version 1.0
// @description 目前仅仅是一个demo
func main() {
	logger := util.InitZapLog()
	fmt.Println(NOTE_IMG)
	fmt.Println(URL_IMG)

	setting.Setup()
	util.RegisterInfo()
	go daemon.Run()
	go util.WSManager.Start()
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
