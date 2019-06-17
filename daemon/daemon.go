package daemon
import ("GoModDemo/util"
"time"
"fmt")
func Run ()  {
	logger := util.InitZapLog()
	logger.Debug("启动后台进程")
	go registerTask()
}
func registerTask(){
	time.Sleep(6*time.Second)

	cache := util.NewCache()
	res, err := cache.Value("deviceId")
	if err == nil {
	  str :="deviceId:"+ res.Data().(string)
	  fmt.Printf(str)
	}
    
}