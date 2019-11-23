package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

type App struct {
	AesKey     string
	JwtSecret  string
	JumpTime   int32
	DelayTime  int32
	OffsetTime int32
	TaskTime   int32
	TaskTime1  int32
	TaskTime2  int32
}
type Server struct {
	Ip   string
	Port string
	Url  string
}
type RedisServer struct {
	Url       string
	DbName    string
	Password  string
	MaxIdle   int32
	MaxActive int32
}
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var AppSetting = &App{}
var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var MogodbSetting = &Server{}
var RedisSetting = &RedisServer{}
var config *ini.File

func Setup() {
	var err error
	config, err = ini.Load("app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("mogodbconfig", MogodbSetting)
	mapTo("redisconfig", RedisSetting)
}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
