package util

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"strings"
	"time"
)

// SCP信息
type SshInfo struct {
	RemoteIp       *string `json:"remoteIp"`
	RemotePort     *int    `json:"remotePort"`
	RemoteUserName *string `json:"remoteUserName"`
	RemotePassword *string `json:"remotePassword"`
}

// 建立连接
func buildSshConnect(host *string, port *int, username *string, password *string) (*ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(*password))

	clientConfig = &ssh.ClientConfig{
		User:    *username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", *host, *port)
	return ssh.Dial("tcp", addr, clientConfig)
	//if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
	//	return nil, err
	//}
	//
	//// create session
	//if session, err = sshClient.NewSession(); err != nil {
	//	return nil, err
	//}
	//if err != nil {
	//	return nil, err
	//}
	//return session, nil
}

// 建立连接
func BuildSshConnectWithSshInfo(sshInfo *SshInfo) (*ssh.Client, error) {
	return buildSshConnect(sshInfo.RemoteIp, sshInfo.RemotePort, sshInfo.RemoteUserName, sshInfo.RemotePassword)
}

// 执行命令
func ExecuteCommand(sshClient *ssh.Client, cmd *string) (*string, error) {
	var (
		session *ssh.Session
		err     error
	)
	// 创建会话
	if session, err = sshClient.NewSession(); err != nil {
		return nil, err
	}
	defer session.Close()
	// 执行语句
	bytes, err := session.Output(*cmd)
	if err != nil {
		return nil, err
	}
	result := string(bytes)
	return &result, nil
}

// 通用的执行命令方法
func CommonExecuteCommand(sshClient *ssh.Client, cmd *string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()
	go func() {
		_, _ = ExecuteCommand(sshClient, cmd)
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("ctx.Done")
	case <-time.After(time.Duration(time.Millisecond * 5000)):
		fmt.Printf("time.After")
	}
	return nil, nil
}

// 获取进程Id
func GetProcessID(sshClient *ssh.Client, serverName *string) (*string, error) {
	cmd := fmt.Sprintf("ps -aux | grep %s | grep -v grep | awk '{print $2}'", *serverName)
	cmdResult, err := ExecuteCommand(sshClient, &cmd)
	if err != nil {
		return nil, err
	}
	pid := strings.ReplaceAll(*cmdResult, "\n", "")
	return &pid, nil
}

// 获取服务的启动命令
func GetJavaStartCommand(sshClient *ssh.Client, serverName *string) (*string, error) {
	cmd := fmt.Sprintf("ps -aux | grep %s | grep -v grep", *serverName)
	cmdResult, err := ExecuteCommand(sshClient, &cmd)
	if err != nil {
		return nil, err
	}

	var command string
	split := strings.Split(*cmdResult, "\n")
	for _, str := range split {
		index := strings.Index(str, "java -jar")
		if index > 0 {
			command = str[strings.Index(str, "java -jar"):]
			break
		}
	}
	if IsEmpty(&command) {
		return nil, errors.New("no start command found")
	}
	command = fmt.Sprintf("nohup %s &", command)
	return &command, nil
}

// 杀死进程
func KillLinuxProcess(sshClient *ssh.Client, pid *string) error {
	cmd := fmt.Sprintf("kill -9 %s", *pid)
	_, err := ExecuteCommand(sshClient, &cmd)
	if err != nil {
		return err
	}
	return nil
}
