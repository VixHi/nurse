package util

import (
	"fmt"
	"regexp"
)

// CheckPwd : 验证密码是否合规
func CheckPwd(pwd string) bool {
	res, err := regexp.MatchString(`[0-9]{6-16}`, pwd)
	if err != nil {
		fmt.Println("错误是", err)
	}
	return res
}
