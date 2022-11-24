package util

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 13:31
 * @Desc:
 */

// IsDir 判断是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateFile 创建文件
func CreateFile(name string) (io.Writer, error) {
	targetDir := path.Dir(name)
	if !IsDir(targetDir) {
		err := os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	fileName := path.Base(name)
	logFileName := fmt.Sprintf("%s/%s.%s.log", targetDir, fileName, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}
	return file, nil
}
