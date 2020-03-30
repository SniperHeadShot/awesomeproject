package util

import "strings"

// 判断字符串是否为空
func IsEmpty(str *string) bool {
	if str == nil {
		return true
	}
	emptyStr := ""
	return strings.Compare(emptyStr, *str) == 0
}

// 判断字符串是否不为空
func IsNotEmpty(str *string) bool {
	return !(IsEmpty(str))
}

// 数组是否包含元素
func IsArrContainEle(arr *[]string, str *string) bool {
	if arr == nil {
		return false
	}
	if str == nil {
		return false
	}

	for _, strTemp := range *arr {
		if strTemp == *str {
			return true
		}
	}
	return false
}
