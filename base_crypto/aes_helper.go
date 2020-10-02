package base_crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"github.com/hkxiaoyu/gobase/base_string"
)

// 去掉padding
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}

// 添加padding
func PKCS7Padding(origData []byte, blockSize int) []byte {
	padding := blockSize - len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padText...)
}

// 去掉padding(nopadding方式)
func UnPadding(origData []byte) []byte {
	length := len(origData)
	paddingChar := int(origData[length-1])
	paddingCount := 0
	for {
		if paddingChar != 0 {
			break
		} else {
			paddingCount++
			paddingChar = int(origData[length-(paddingCount+1)])
		}
	}
	return origData[:length-paddingCount]
}

// 添加padding(nopadding方式)
func Padding(origData []byte, blockSize int) []byte {
	padding := blockSize - len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(0)}, padding)
	return append(origData, padText...)
}

// AES128 CBC加密(hex)
func AESCbcEncrypt(origData []byte, key []byte) string{
	block,_:=aes.NewCipher(key)
	origData=PKCS7Padding(origData, block.BlockSize())
	blockMode:=cipher.NewCBCEncrypter(block,key[:block.BlockSize()])
	cryptData:=make([]byte,len(origData))
	blockMode.CryptBlocks(cryptData,origData)
	encodeString:=hex.EncodeToString(cryptData)
	return encodeString
}

// AES128 CBC解密(hex)
func AESCbcDecrypt(cryptData string, key []byte) []byte{
	decodeData, _ := hex.DecodeString(cryptData) //解码BASE64
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(origData, decodeData)
	origData = PKCS7UnPadding(origData)
	return origData
}

// AES128 CBC加密(base64)
func AESCbcEncryptV2(origData []byte, key []byte) string {
	block, _ := aes.NewCipher(key)
	origData = PKCS7Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	cryptData := make([]byte, len(origData))
	blockMode.CryptBlocks(cryptData, origData)
	encodeString := base64.StdEncoding.EncodeToString(cryptData)
	return encodeString
}

// AES128 CBC解密(base64)
func AESCbcDecryptV2(crpyped string, key []byte) []byte {
	decodeData, _ := base64.StdEncoding.DecodeString(crpyped) //解码BASE64
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(origData, decodeData)
	origData = PKCS7UnPadding(origData)
	return origData
}

//AES128 CTR加密(hex)
func AESCtrEncrypt(data []byte, key []byte) string {
	block, _ := aes.NewCipher(key)
	iv := []byte(base_string.StrGetRandomString(16)) //获取16个字节长度的随机IV
	blockMode := cipher.NewCTR(block, iv)
	message := make([]byte, len(data))
	blockMode.XORKeyStream(message, data)
	fullMessage := append(iv, message...)
	hexMessage := hex.EncodeToString(fullMessage)
	return hexMessage
}

//AES128 CTR解密(hex)
func AESCtrDecrypt(data string, key []byte) []byte {
	bData, err := hex.DecodeString(data)
	if err == nil {
		iv := bData[0:16]                 //取出随机生成的IV
		cryptData := bData[16:len(bData)] //取出加密后的数据
		block, _ := aes.NewCipher(key)
		blockMode := cipher.NewCTR(block, iv) //使用取出的随机IV
		message := make([]byte, len(cryptData))
		blockMode.XORKeyStream(message, cryptData) //解密真正的内容
		return message
	}
	return []byte("")
}
