package core

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type NetworkInfo struct {
	MAC string `json:"mac"`
	IP  string `json:"ip"`
	TTL string `json:"ttl"`
}

type NetworkPair struct {
	Sect     map[string][]*NetworkInfo `json:"content"`
	AreaList []string
}

type JsonFile struct {
	Local string
	Date  string
}

// Sort 对Json文件进行排序
func (info *JsonFile) Sort() {
	info.Date = NowTime()
	resDir := GetResDir("")
	jsonFiles, _ := GetDirFNames(resDir, false)

	for _, file := range jsonFiles {
		fileInfo := strings.Split(file, ".")
		if info.Date == fileInfo[0] {
			info.Local = file
			continue
		}
		if fileInfo[1] == "exe" {
			continue
		}
		if err := os.Remove(filepath.Join(resDir, file)); err != nil {
			log.Printf("Error removing file %v", file)
		}
	}
	if info.Local == "" {
		color.Red("[DEBUG] 数据文件过期了, 请执行 snet-tools.exe web")
		os.Exit(0)
	}
}
