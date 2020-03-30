package util

import (
	"fmt"
	"strings"
)

// 获取总数 PO
type TotalStruct struct {
	Total *int `json:"total" field:"total"`
}

// SQL提升为分页
func BuildPageDataSql(sql *string, page int, limit int) *string {
	pagingDataSql := fmt.Sprintf("select * from (%s) _paging limit %d,%d", *sql, (page-1)*limit, limit)
	return &pagingDataSql
}

// SQL提升为查询总数
func BuildPageTotalSql(sql *string) *string {
	pagingTotalSql := fmt.Sprintf("select count(1) as total from (%s) _paging", *sql)
	return &pagingTotalSql
}

// 如果字符串使用in关键字查询，需要给每一项都加上单引号
func BuildInString(param string) *string {
	if param == "" {
		return nil
	}
	arr := strings.Split(param, ",")
	return BuildInStringWithArr(arr)
}

// 将一个字符串数组转换为'',''的形式
func BuildInStringWithArr(strArr []string) *string {
	if strArr == nil || len(strArr) == 0 {
		return nil
	}
	var result string
	paramLen := len(strArr)
	for k := 0; k < paramLen; k++ {
		temp := strings.Trim(strArr[k], "")
		if temp != "" {
			if k == paramLen-1 {
				result += fmt.Sprintf("'%s'", temp)
			} else {
				result += fmt.Sprintf("'%s',", temp)
			}
		}
	}
	return &result
}
