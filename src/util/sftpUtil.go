package util

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"path"
	"time"
)

// SCP信息
type ScpInfo struct {
	LocalFilePath  *string `json:"localFilePath"`
	RemoteFilePath *string `json:"remoteFilePath"`
	RemoteIp       *string `json:"remoteIp"`
	RemotePort     *int    `json:"remotePort"`
	RemoteUserName *string `json:"remoteUserName"`
	RemotePassword *string `json:"remotePassword"`
}

// 建立连接
func buildSftpConnect(host *string, port *int, username *string, password *string) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
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
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}

// 传输文件
func SftpTransferFile(scpInfo *ScpInfo) error {
	sftpClient, err := buildSftpConnect(scpInfo.RemoteIp, scpInfo.RemotePort, scpInfo.RemoteUserName, scpInfo.RemotePassword)
	if err != nil {
		return nil
	}
	srcFile, err := os.Open(*scpInfo.LocalFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(*scpInfo.LocalFilePath)
	dstFile, err := sftpClient.Create(path.Join(*scpInfo.RemoteFilePath, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf[0:n])
	}
	return nil
}
