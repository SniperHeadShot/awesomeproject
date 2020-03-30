package util

import (
	"awesomeproject/src/dict"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"runtime"
)

// 判断网络是否正常
func CheckNetworkReachable(ip *string) bool {
	var cmd *exec.Cmd

	if dict.SystemWindows == runtime.GOOS {
		cmd = exec.Command("ping", *ip)
	}
	if dict.SystemLinux == runtime.GOOS {
		cmd = exec.Command("ping", *ip, "-c", "1", "-W", "5")
	}
	if cmd == nil {
		return false
	}
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

// 判断服务是否启动
func CheckServerPortUse(ip *string, port *int) bool {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	return err == nil
}

// 发送Get请求
func SendHttpGet(url *string, result interface{}) *string {
	var msg string
	if url == nil || IsEmpty(url) {
		msg = fmt.Sprint("请求参数不合法")
		return &msg
	}

	resp, err := http.Get(*url)
	if err != nil {
		msg = fmt.Sprintf("获取本地服务列表失败 [%v]", err)
		return &msg
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg = fmt.Sprintf("读取响应数据失败 [%v]", err)
		return &msg
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		msg = fmt.Sprintf("json解析失败 [%v]", err)
		return &msg
	}
	return nil
}
