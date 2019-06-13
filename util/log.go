package util

import (
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logpath 日志文件路径
// loglevel 日志级别
func initLogger(logpath string, loglevel string) *zap.Logger {

	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    128,     // megabytes
		MaxBackups: 30,      // 最多保留300个备份
		MaxAge:     7,       // days
		Compress:   true,    // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core)
	logger.Info("DefaultLogger init success")

	return logger
}

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func InitZapLog() *zap.Logger {
	var t time.Time
	t = time.Now()
	dataStr := t.Format("2006-01-02")
	logNameStr := "./log/" + dataStr + ".log"
	return initLogger(logNameStr, "debug")
}
