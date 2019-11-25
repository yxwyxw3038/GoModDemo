package util

import (
	"GoModDemo/setting"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
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
func UnBase640(str string) (string, error) {
	returnStr := ""
	decodeString, err := AesDecryptStr(str)
	if err != nil {
		return returnStr, err
	}
	unStr, err := url.QueryUnescape(decodeString)
	if err != nil {
		return returnStr, err
	}
	return unStr, nil

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
	// block, err := aes.NewCipher(key)
	// if err != nil {
	// 	return nil, err
	// }

	// blockSize := block.BlockSize()
	// origData = PKCS5Padding(origData, blockSize)
	// blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// crypted := make([]byte, len(origData))
	// blockMode.CryptBlocks(crypted, origData)
	// return crypted, nil

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//填充原文
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	//初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(origData))
	//block大小 16
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	//block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], origData)

	return cipherText, nil

}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	// block, err := aes.NewCipher(key)
	// if err != nil {
	// 	return nil, err
	// }

	// blockSize := block.BlockSize()
	// blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	// origData := make([]byte, len(crypted))
	// blockMode.CryptBlocks(origData, crypted)
	// origData = PKCS5UnPadding(origData)
	// return origData, nil
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(crypted) < blockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := crypted[:blockSize]
	crypted = crypted[blockSize:]
	if len(crypted)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(crypted, crypted)
	//解填充
	crypted = PKCS5UnPadding(crypted)
	return crypted, nil
}
func AesEncryptIV(origData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecryptIV(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil

}

func AesEncryptStr(origStr string) (string, error) {
	keyStr := setting.AppSetting.AesKey
	keyIV := setting.AppSetting.AesIV
	keyData := []byte(keyStr)
	ivData := []byte(keyIV)
	origData := []byte(origStr)
	data, err := AesEncryptIV(origData, keyData, ivData)
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
	// cryptedData := []byte(cryptedStr)
	keyStr := setting.AppSetting.AesKey
	keyIV := setting.AppSetting.AesIV
	keyData := []byte(keyStr)
	ivData := []byte(keyIV)
	// cryptedData := []byte(cryptedStr)
	data, err := AesDecryptIV(cryptedData, keyData, ivData)
	if err != nil {
		return "", err
	}
	pass64 := string(data)
	return pass64, nil
}
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
func DesEncryptStr(origStr string) (string, error) {
	keyStr := setting.AppSetting.AesKey
	keyData := []byte(keyStr)
	origData := []byte(origStr)
	data, err := DesEncrypt(origData, keyData)
	if err != nil {
		return "", err
	}
	pass64 := base64.StdEncoding.EncodeToString(data)
	return pass64, nil
}

func DesDecryptStr(cryptedStr string) (string, error) {
	cryptedData, err := base64.StdEncoding.DecodeString(cryptedStr)
	if err != nil {
		return "", err
	}

	keyStr := setting.AppSetting.AesKey
	keyData := []byte(keyStr)
	// cryptedData := []byte(cryptedStr)
	data, err := DesDecrypt(cryptedData, keyData)
	if err != nil {
		return "", err
	}
	pass64 := string(data)
	return pass64, nil
}
