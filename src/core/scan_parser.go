package core

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

// AppLocal 获取可执行文件路径
func AppLocal(appName string) (appLocal string) {
	if appName == "" {
		appName = "scan.exe"
	}
	resDir := GetResDir("")
	appLocal = filepath.Join(resDir, appName)
	return appLocal
}

// NowTime 获取现在时间
func NowTime() string {
	return time.Now().Local().Format("2006-01-02")
}

// StartApp 启动扫描程序
func StartApp(area string) string {
	appLocal := AppLocal("")

	color.Green(fmt.Sprintf("[INFO] Will Scan area => %s", area))

	cmd := exec.Command(appLocal, "-t", area)
	color.Blue("[INFO] Please wait for some time...")

	result, err := cmd.Output()
	if err != nil {
		color.Red("[cmd Error] " + err.Error())
		os.Exit(0)
	}
	return string(result)
}

// InitNetWorkPair 数据的构造函数
func InitNetWorkPair() *NetworkPair {
	return &NetworkPair{
		Sect:     make(map[string][]*NetworkInfo, 0),
		AreaList: make([]string, 0),
	}
}

// 单次解析结果
func noeParser(OneScanData string) (*NetworkInfo, string) {
	splitData := strings.Split(OneScanData, " ")
	ttlParser := strings.Split(splitData[6], ".")
	scanArea := strings.Split(splitData[4], ".")[2]

	// 构造结构体
	var info = &NetworkInfo{
		MAC: splitData[2],
		IP:  splitData[4],
		TTL: ttlParser[0],
	}

	return info, scanArea
}

// ParserResult 最终解析结果
func (netPir *NetworkPair) ParserResult(result string) {
	var sectList []string

	for _, line := range strings.Split(result, "\n") {
		if line == "" {
			break
		}
		// 解析 扫描 结果
		info, sect := noeParser(line)
		sectList = append(sectList, sect)

		if _, ok := netPir.Sect[sect]; !ok {
			// 若网段不属于Sect则为该网段创建一个空数组
			netPir.Sect[sect] = make([]*NetworkInfo, 0)
		} else {
			netPir.Sect[sect] = append(netPir.Sect[sect], info)
		}
	}

	// 遍历map，删除值为空的键值对
	for _, info := range sectList {
		if len(netPir.Sect[info]) == 0 {
			delete(netPir.Sect, info)
		}
	}
}

// GetAreaList 获取区域 列表
func (netPir *NetworkPair) GetAreaList() {
	ipList := make([]string, len(netPir.Sect))

	count := 0
	for index := range netPir.Sect {
		ipList[count] = index
		count++
	}
	netPir.AreaList = ipList
}

// SaveJsonFile 存储JSON文件
func (netPir *NetworkPair) SaveJsonFile() {
	// 资源路径
	resDir := GetResDir("")
	saveName := filepath.Join(resDir, "result.json")

	// 存储文件
	fileObj, err := os.Create(saveName)
	if err != nil {
		color.Red("[CreateFile]" + err.Error())
		os.Exit(0)
	}

	defer func() {
		_ = fileObj.Close()
	}()

	// Json 格式 转 字节码
	byteData, err := json.Marshal(&netPir)
	if err != nil {
		color.Red("[JsonChange] " + err.Error())
		os.Exit(0)
	}
	// 写入文件
	_, err = fileObj.Write(byteData)
	if err != nil {
		color.Red("[WriteFile] " + err.Error())
		os.Exit(0)
	}
	color.Blue("[INFO] SaveJsonFile succeeded!")
}
