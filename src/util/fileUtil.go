package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 从程序执行路径读取配置文件
func ReadConfigFromFile(result interface{}) error {
	executePath, err := os.Executable()
	if IsEmpty(&executePath) {
		return errors.New("get program path execution failed")
	}
	configFilePath := filepath.Join(filepath.Dir(executePath), "./config.json")
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("读取配置文件[%s]失败, 原因为:%v", configFilePath, err)
		return err
	}
	err = json.Unmarshal(file, result)
	if err != nil {
		return err
	}
	return nil
}
