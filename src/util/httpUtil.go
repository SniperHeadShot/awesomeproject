package util

import (
	"awesomeproject/src/dict"
	"encoding/json"
	"errors"
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
func SendHttpGet(url *string, result interface{}) error {
	if IsEmpty(url) {
		return errors.New("url cannot be empty")
	}

	resp, err := http.Get(*url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}
