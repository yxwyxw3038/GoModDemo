package util

import (
	"GoModDemo/setting"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

func UnBase64(str string) (string, error) {
	returnStr := ""
	bufstr1 := append([]byte(str)[1:], []byte(str)[0])
	mwString := string(bufstr1)
	decodeBytes, err := base64.StdEncoding.DecodeString(mwString)

	if err != nil {
		return returnStr, err
	}
	decodeString := string(decodeBytes)
	unStr, err := url.QueryUnescape(decodeString)
	if err != nil {
		return returnStr, err
	}
	return unStr, nil

	// returnStr := ""
	// if str == "" {
	// 	return returnStr, nil
	// }
	// unStr, err := url.QueryUnescape(str)
	// if err != nil {
	// 	return returnStr, err
	// }
	// decodeBytes, err := base64.StdEncoding.DecodeString(unStr)
	// if err != nil {
	// 	return returnStr, err
	// }
	// decodeString := string(decodeBytes)
	// decodeString, err = AesDecryptStr(decodeString)
	// if err != nil {
	// 	return returnStr, err
	// }
	// return decodeString, nil

}
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	md5str2 := strings.ToUpper(md5str1)
	return md5str2
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func AesEncryptStr(origStr string) (string, error) {
	keyStr := setting.AppSetting.AesKey
	keyData := []byte(keyStr)
	origData := []byte(origStr)
	data, err := AesEncrypt(origData, keyData)
	if err != nil {
		return "", err
	}
	pass64 := base64.StdEncoding.EncodeToString(data)
	return pass64, nil
}

func AesDecryptStr(cryptedStr string) (string, error) {
	cryptedData, err := base64.StdEncoding.DecodeString(cryptedStr)
	if err != nil {
		return "", err
	}

	keyStr := setting.AppSetting.AesKey
	keyData := []byte(keyStr)
	// cryptedData := []byte(cryptedStr)
	data, err := AesDecrypt(cryptedData, keyData)
	if err != nil {
		return "", err
	}
	pass64 := string(data)
	return pass64, nil
}
