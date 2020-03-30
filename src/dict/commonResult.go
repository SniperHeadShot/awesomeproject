package dict

type ErrorEntity struct {
	Message string
	Code    int
}

// 全局成功
var GlobalSuccMsg = ErrorEntity{
	Message: "succ",
	Code:    0,
}

// 全局失败
var GlobalFailMsg = ErrorEntity{
	Message: "unknown fail",
	Code:    -1,
}

// 参数校验失败
var ParamVerityFailMsg = ErrorEntity{
	Message: "parameter check failed",
	Code:    1001,
}

// 解析Http请求参数失败
var ParamParseMsg = ErrorEntity{
	Message: "failed to parse HTTP request parameters",
	Code:    1002,
}

// 执行sql失败
var SqlExecuteError = ErrorEntity{
	Message: "failed to execute sql",
	Code:    2001,
}

// 获取mysql连接失败
var MysqlConnectError = ErrorEntity{
	Message: "failed to get mysql connection",
	Code:    2002,
}

// 获取文件服务器失败
var GetFileServerError = ErrorEntity{
	Message: "failed to get file server",
	Code:    3001,
}

// 上传文件到文件服务器失败
var FileServerUploadError = ErrorEntity{
	Message: "file upload to file server failed",
	Code:    3002,
}

// 文件服务器删除失败
var FileServerDeleteError = ErrorEntity{
	Message: "file server failed to delete files",
	Code:    3003,
}
