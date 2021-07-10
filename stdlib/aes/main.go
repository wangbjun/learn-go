package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	// 需要被加密的内容，需要填充
	var src = "Hello，我是一个测试加密内容你知道吗？？？"
	// key必须是16\24\32位
	var key = "1234567890123456"

	// 使用了PKCS7填充法
	cipherText := AesCBCEncrypt(PKCS7Padding([]byte(src), aes.BlockSize), []byte(key))

	// 为方便展示，用base64编码
	fmt.Printf("cipherText text is %s\n", base64.StdEncoding.EncodeToString(cipherText))

	// 解密
	plainText := AesCBCDecrypt(cipherText, []byte(key))
	fmt.Printf("plain text is %s\n", PKCS7UnPadding(plainText))
}

func AesCBCEncrypt(plainText []byte, key []byte) []byte {
	// 生成加密用的block
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 对IV有随机性要求，但没有保密性要求，所以常见的做法是将IV包含在加密文本当中
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText
}

func AesCBCDecrypt(cipherText []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(cipherText) < aes.BlockSize {
		panic("cipher text too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipher text is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	return cipherText
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ZeroPadding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(src, padText...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
