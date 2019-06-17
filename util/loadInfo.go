package util

import (
	"github.com/google/uuid"

)

func RegisterInfo() {
	cache := NewCache()
	newUuid := uuid.New()
	newUuidStr := newUuid.String()
	logger := InitZapLog()
	logger.Debug("注册设备号：" + newUuidStr + " ps当前设备号重启将发生变化")
	cache.Add("deviceId", 0, newUuidStr)
}


