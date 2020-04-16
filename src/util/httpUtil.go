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

type HttpRequestStrut struct {
	Method string
	Url    *string
	Header *map[string]string
}

func (httpRequestStrut HttpRequestStrut) verityPass() bool {
	return IsNotEmpty(&httpRequestStrut.Method) && IsNotEmpty(httpRequestStrut.Url)
}

// 发送Http请求
func SendHttpRequest(httpRequestStrut *HttpRequestStrut, result interface{}) error {
	var (
		request       *http.Request
		response      *http.Response
		responseBytes []byte
		err           error
	)
	if httpRequestStrut == nil || !httpRequestStrut.verityPass() {
		return errors.New("url cannot be empty")
	}

	// 构造请求
	request, err = http.NewRequest(httpRequestStrut.Method, *httpRequestStrut.Url, nil)
	if err != nil {
		return nil
	}

	// 设置header
	if httpRequestStrut.Header != nil && len(*httpRequestStrut.Header) > 0 {
		for key, value := range *httpRequestStrut.Header {
			request.Header[key] = []string{value}
		}
	}

	// 请求
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 读取结果
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseBytes, result)
	if err != nil {
		return err
	}

	return nil
}
