package utils

import (
	"strconv"
	"strings"
)

type ErrStruct struct {
	Code int32
	Msg  string
}

// GetErrorMap 读取 error 文本，解析为 struct 返回给前端
func GetErrorMap(errStr string) ErrStruct {
	// 定义 err 变量
	var err ErrStruct
	// 切割错误文案
	errStrSplice := strings.Split(errStr, ",")
	// 遍历切片
	for _, value := range errStrSplice {
		valueSplice := strings.Split(value, ":")
		switch true {
		// 赋值 code
		case strings.Contains(value, "code"):
			code, errConv := strconv.ParseInt(valueSplice[1], 10, 32)
			if errConv != nil {
				err = ErrStruct{
					Code: -1,
					Msg:  errStr,
				}
				return err
			}
			err.Code = int32(code)
		case strings.Contains(value, "msg"):
			err.Msg = valueSplice[1]
		}
	}

	return err
}
