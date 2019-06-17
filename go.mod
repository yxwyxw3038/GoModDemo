module GoModDemo

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net => github.com/golang/net v0.0.0-20190607181551-461777fb6f67
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190610200419-93c9922d18ae
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190610231749-f8d1dee965f7
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/astaxie/beego v1.11.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.4.0
	github.com/google/uuid v1.1.1
	github.com/jinzhu/configor v1.0.0
	github.com/muesli/cache2go v0.0.0-20190609140403-5eb79359852d
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	gopkg.in/ini.v1 v1.42.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
