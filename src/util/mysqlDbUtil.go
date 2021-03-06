package util

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ConnectInfo struct {
	Host     *string `json:"host"`
	Port     *int    `json:"port"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	DbName   *string `json:"dbName"`
}

func (connectInfo ConnectInfo) verityPass() bool {
	return IsNotEmpty(connectInfo.Host) && *connectInfo.Port > 0 && IsNotEmpty(connectInfo.Username) && IsNotEmpty(connectInfo.DbName)
}

// 建立Mysql连接
func BuildDb(connInfo *ConnectInfo) (*sql.DB, error) {
	if connInfo == nil || !connInfo.verityPass() {
		return nil, errors.New("db connect info is empty")
	}
	path := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *connInfo.Username, *connInfo.Password, *connInfo.Host, *connInfo.Port, *connInfo.DbName)
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// 最大打开连接数
	db.SetMaxOpenConns(20)
	// 最大空闲连接数
	db.SetMaxIdleConns(10)
	// 连接过期时间
	db.SetConnMaxLifetime(time.Second * 10)
	return db, nil
}
