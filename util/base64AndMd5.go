package util

import (	
	"encoding/base64"
	"net/url"
	"crypto/md5"
	"fmt"
	"strings"
)
func UnBase64(str string) (string, error) {
	returnStr:=""
	bufstr1 := append([]byte(str)[1:], []byte(str)[0])
	mwString := string(bufstr1)
	decodeBytes, err := base64.StdEncoding.DecodeString(mwString)
	
	if err != nil {
		return returnStr, err
	}
	decodeString := string(decodeBytes)
	unStr,err:= url.QueryUnescape(decodeString)
	if err != nil {
		return returnStr, err
	}
	return unStr, nil
}
func Md5(str string) (string) {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) 
	md5str2:=strings.ToUpper(md5str1)
	return md5str2
	}