package util

import (
	"awesomeproject/src/dict"
)

type CommonResult struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	ErrCode int         `json:"errCode"`
	Msg     string      `json:"msg"`
}

type PageInfo struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

func BuildSuccDataRsp(data interface{}) *CommonResult {
	res := CommonResult{
		Data:    data,
		Success: true,
		ErrCode: dict.GlobalSuccMsg.Code,
		Msg:     dict.GlobalSuccMsg.Message,
	}
	return &res
}

func BuildSuccPageInfoRsp(data interface{}, total int) *CommonResult {
	var pageInfo = PageInfo{
		List:  data,
		Total: total,
	}

	res := CommonResult{
		Data:    pageInfo,
		Success: true,
		ErrCode: dict.GlobalSuccMsg.Code,
		Msg:     dict.GlobalSuccMsg.Message,
	}
	return &res
}

func BuildStandardRsp(succ bool, resCode int, msg string, data interface{}) *CommonResult {
	res := CommonResult{
		Data:    data,
		Success: succ,
		ErrCode: resCode,
		Msg:     msg,
	}
	return &res
}

func BuildFailConstantRsp(constant dict.ErrorEntity) *CommonResult {
	res := CommonResult{
		Data:    nil,
		Success: false,
		ErrCode: constant.Code,
		Msg:     constant.Message,
	}
	return &res
}

func BuildFailMsgRsp(msg string) *CommonResult {
	res := CommonResult{
		Data:    nil,
		Success: false,
		ErrCode: dict.GlobalFailMsg.Code,
		Msg:     msg,
	}
	return &res
}
