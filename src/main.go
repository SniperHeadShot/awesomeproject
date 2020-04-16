package main

import (
	"awesomeproject/src/util"
	"fmt"
	"net/http"
)

func main() {
	url := "http://192.168.9.237:10080/pmdbsvr/project/list?page=1&limit=10"
	//header := make(map[string]string)
	//header["projectUuid"] = "ccac3a91bac54c02a3e3926268d16e5b"

	httpRequestStrut := util.HttpRequestStrut{
		Method: http.MethodGet,
		Url:    &url,
		//Header: &header,
	}

	var result util.CommonResult

	err := util.SendHttpRequest(&httpRequestStrut, &result)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// java 重启demo
	//ip := "192.168.9.237"
	//port := 22
	//username := "root"
	//password := "Aa123!@#."
	//
	//sshInfo := util.SshInfo{
	//	RemoteIp:       &ip,
	//	RemotePort:     &port,
	//	RemoteUserName: &username,
	//	RemotePassword: &password,
	//}
	//
	//sshClient, err := util.BuildSshConnectWithSshInfo(&sshInfo)
	//if err != nil {
	//	fmt.Printf("sshClient create fail: msg=[%v] \n", err)
	//}
	//
	//serverName := "sppc-iac-app-service"
	//// 获取进程ID
	//pid, err := util.GetProcessID(sshClient, &serverName)
	//if err != nil {
	//	fmt.Printf("get pid fail: msg=[%v] \n", err)
	//}
	//if util.IsEmpty(pid) {
	//	fmt.Println("not find server, system exit")
	//	return
	//}
	//// 获取启动命令
	//command, err := util.GetJavaStartCommand(sshClient, &serverName)
	//if err != nil {
	//	fmt.Printf("get serverName commond fail: msg=[%v] \n", err)
	//}
	//// 杀死进程
	//err = util.KillLinuxProcess(sshClient, pid)
	//if err != nil {
	//	fmt.Printf("kill process fail: msg=[%v] \n", err)
	//}
	//// 重启服务
	//_, err = util.CommonExecuteCommand(sshClient, command)
	//if err != nil {
	//	fmt.Printf("restart server fail: msg=[%v] \n", err)
	//}
	//fmt.Printf("重启成功")
}
