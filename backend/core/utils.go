package core

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ResDir = GetResDir("")

// BaseDir 获取文件基础路径
func BaseDir() string {
	baseDir, _ := os.Getwd()
	return baseDir
}

// IsDir 判断是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		log.Println("[Warrning] 无法读取文件信息")
	}
	return s.IsDir()
}

// GetDirFNames 获取指定文件夹下所有文件
func GetDirFNames(searchDir string, Flag bool) (allFile []string, err error) {
	var fileInfo string

	files, err := os.ReadDir(searchDir)
	if err != nil {
		log.Printf("[DEBUG] Search %v", err)
		return nil, err
	}
	// 遍历获取文件
	for _, file := range files {
		fileInfo = file.Name()
		if !Flag {
			endWith := strings.Split(fileInfo, ".")
			if !IsDir(filepath.Join(ResDir, fileInfo)) || endWith[0] == "json" {
				allFile = append(allFile, fileInfo)
			}
		} else {
			allFile = append(allFile, fileInfo)
		}
	}
	return allFile, nil
}

// GetResDir 获取资源目录
func GetResDir(pathLoad string) (resDir string) {
	if pathLoad == "" {
		pathLoad = "res"
	}
	// 用户数据目录
	_, home := GetUserPath()
	// 工程路径
	resDir = filepath.Join(home, pathLoad)
	// 不存在就创建
	if _, err := os.Stat(resDir); err != nil {
		_ = os.MkdirAll(resDir, 0777)
		_ = os.Chmod(resDir, 0777)
	}
	return resDir
}

// GetUserPath 获取用户路径
func GetUserPath() (bool, string) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}

	if home == "" {
		home, _ = os.Getwd()
		return false, home
	} else {
		home = filepath.Join(home, ".config", "foxer")
		if exist, _ := PathExists(home); !exist {
			_ = os.MkdirAll(home, 0755)
		}
	}

	return true, home
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
