package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

//func FileMD5(filePath string) (string, error) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		return "", err
//	}
//	all, err := ioutil.ReadAll(file)
//	if err != nil {
//		return "", err
//	}
//	return MD5(string(all)), nil
//}

func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
