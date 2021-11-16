package util

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesCBCDecrypt(t *testing.T) {
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
