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
	// 工程路径
	basePath, _ := os.Getwd()
	resDir = filepath.Join(basePath, pathLoad)
	// 不存在就创建
	if _, err := os.Stat(resDir); err != nil {
		_ = os.MkdirAll(resDir, 0777)
		_ = os.Chmod(resDir, 0777)
	}
	return resDir
}

// GetFontPath 导入字体资源文件
func GetFontPath() string {
	return filepath.Join(ResDir, `fonts\sarasa-mono-sc-regular.ttf`)
}
