package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"
	"time"
)

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(crypted)%blockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	if origData == nil {
		return nil, errors.New("index out of range")
	}
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	if unpadding < 0 {
		return nil
	}
	return origData[:(length - unpadding)]
}

//根据日期生成key
func get_key() []byte {
	key := "gfe03f" + time.Now().Format("02") + "_9fd&wfl"
	return []byte(key)
}

//加密字符串
func Encription(bt []byte) string {
	result, err := AesEncrypt(bt, get_key())
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(result)
}

//解密字符串
func Decription(str string) string {
	result := ""
	defer func() {
		if r := recover(); r != nil {
			log.Println("[W]", r)
		}
	}()
	encodeStr, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	bts, err := AesDecrypt(encodeStr, get_key())
	if err != nil {
		println("token:" + str + " err:" + err.Error())
		return ""
	}
	result = string(bts)
	return result
}
