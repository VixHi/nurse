/**
name : vapi
date : 2020-08-14 16:61
autor: vixliu
describe: 项目帮助小工具

*/

package vutil

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"time"
)

// CheckPwd : 验证密码是否合规
func CheckPwd(pwd string) bool {
	res, err := regexp.MatchString(`[0-9]{6-16}`, pwd)
	if err != nil {
		fmt.Println("错误是", err)
	}
	return res
}

// RandNum : 生成随机数 min 生成的最小数 max生成的最大数
func RandNum(min, max int) (randNum int) {

	rand.Seed(time.Now().UnixNano())
	randNum = rand.Intn(max-min) + min
	return
}

// VHashValue : 获取文件的hash值
func VHashValue(filtpath string) (hashvalue string, err error) {

	file, err := os.Open(filtpath)
	if err != nil {
		return "", err
	}

	sha_256 := sha256.New()
	_, err = io.Copy(sha_256, file)
	if err != nil {
		return "", err
	}
	hash := sha_256.Sum(nil)
	hashvalue = hex.EncodeToString(hash)
	return hashvalue, nil

}
