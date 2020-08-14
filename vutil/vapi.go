/**
name : vapi
date : 2020-08-14 16:61
autor: vixliu
describe: 协助处理http请求小工具

*/

package vutil

// ResponseWith : http数据返回
func ResponseWith(code int, message string, data interface{}) map[string]interface{} {

	return map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
}
