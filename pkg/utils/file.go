package utils

import (
	"os"
	"path/filepath"
	"time"
)

// CreateDateDir 依照时间创建目录
func CreateDateDir(basePath string) (string, string, error) {
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		if err = os.Mkdir(basePath, 0777); err != nil {
			return "", "", err
		}
		if err = os.Chmod(basePath, 0777); err != nil {
			return "", "", err
		}
	}
	currentTime := time.Now()
	dirName1 := currentTime.Format("2006")
	dirName2 := currentTime.Format("01")
	dirName3 := currentTime.Format("02")
	foldPath := filepath.Join(basePath, dirName1)
	if _, err := os.Stat(foldPath); os.IsNotExist(err) {
		if err = os.Mkdir(foldPath, 0777); err != nil {
			return "", "", err
		}
		if err = os.Chmod(foldPath, 0777); err != nil {
			return "", "", err
		}
	}
	foldPath = filepath.Join(foldPath, dirName2)
	if _, err := os.Stat(foldPath); os.IsNotExist(err) {
		if err = os.Mkdir(foldPath, 0777); err != nil {
			return "", "", err
		}
		if err = os.Chmod(foldPath, 0777); err != nil {
			return "", "", err
		}
	}
	foldPath = filepath.Join(foldPath, dirName3)
	if _, err := os.Stat(foldPath); os.IsNotExist(err) {
		if err = os.Mkdir(foldPath, 0777); err != nil {
			return "", "", err
		}
		if err = os.Chmod(foldPath, 0777); err != nil {
			return "", "", err
		}
	}
	return foldPath, dirName3, nil
}

func CheckFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func DelFile() error {
	err := os.Remove("static/upload/2021/05/14")
	return err
}
