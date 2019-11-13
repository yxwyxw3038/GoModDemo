package util

import (
	"github.com/noaway/dateparse"
	"time"
)

func ParseAny(timeStr string) (time.Time, error) {
	t, err := dateparse.ParseAny(timeStr)
	if err != nil {
		return time.Now(), err
	}
	return t, nil
}
func ParseAnyToStr(timeStr string) (string, error) {
	if timeStr == "" {
		return "", nil
	}
	t, err := ParseAny(timeStr)
	if err != nil {
		return "", err
	}
	str := t.Format("2006-01-02 15:04:05")

	return str, nil
}
func AnyToTimeStr(obj interface{}) (string, error) {
	if obj == nil {
		return "", nil
	}
	objStr := obj.(time.Time)
	str := objStr.Format("2006-01-02 15:04:05")
	// if err != nil {
	// 	return "",err
	// }
	return str, nil
}
func GetNowStr() string {
	currentTime := time.Now()
	str := currentTime.Format("2006-01-02 15:04:05")
	return str
}
func GetNowAndStr() (time.Time, string) {
	currentTime := time.Now()
	str := currentTime.Format("2006-01-02 15:04:05")
	return currentTime, str
}
func GetMaskDataStr(maskInfo string, now time.Time) (string, error) {
	str := ""
	switch maskInfo {
	case "yyyy":
		str = now.Format("2006")
		break
	case "yyyyMM":
		str = now.Format("200601")
		break
	case "yyyyMMdd":
		str = now.Format("20060102")
		break
	}
	return str, nil
}
