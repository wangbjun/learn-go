package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	//fmt.Println(MD5("12345"))
	//fmt.Println(Sha1("12345"))
	fmt.Println(FileMD5("/home/jwang/Downloads/切尔诺贝利.Chernobyl.S01E04.720p-天天美剧字幕组.mp4"))
}

func Md51(s string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum[:])
}

func Md52(s string) string {
	sum := md5.Sum([]byte(s))
	return fmt.Sprintf("%x\n", sum)
}

func Md53(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func Md54(s string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, s)
	return hex.EncodeToString(hash.Sum(nil))
}

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func Sha1(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
