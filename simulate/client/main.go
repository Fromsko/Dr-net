//go:generate goversioninfo
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"client/check"
)

type JSONData struct {
	Content Content `json:"content"`
	Time    string  `json:"time"`
}

type Content struct {
	ScanIP string `json:"scan_ip"`
	Mac    string `json:"mac"`
	TTL    string `json:"ttl"`
}

func Login(urlIp string) (string, error) {
	var jsonData JSONData

	// 自定义URL
	if urlIp == "" {
		urlIp = "http://localhost:9000/api/v1/ip"
	}

	// 获取数据
	resp, err := http.Get(urlIp)
	if err != nil {
		return "", err
	}

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// 读取数据
	info, _ := io.ReadAll(resp.Body)

	// JSON 格式解析
	if err := json.Unmarshal([]byte(info), &jsonData); err != nil {
		return "", err
	}

	// 获取IP地址
	scanIP := jsonData.Content.ScanIP

	// 输出
	fmt.Println(scanIP)

	return scanIP, nil
}

func NetSet(ip string) (bool, error) {
	subnetMask := "255.255.240.0"
	gateway := "10.25.128.1"
	wifiName := "WLAN"
	primaryDNS := "180.76.76.76"
	secondaryDNS := "119.29.29.29"

	// 设置本地 IP 地址
	cmd := exec.Command("netsh", "interface", "ip", "set", "address",
		"name="+wifiName, "source=static", ip, subnetMask, gateway)
	err := cmd.Run()
	if err != nil {
		log.Fatalln("设置失败: " + err.Error())
	}

	// 设置主 DNS 服务器
	cmd = exec.Command("netsh", "interface", "ip", "set", "dns",
		wifiName, "static", primaryDNS, "primary")
	err = cmd.Run()
	if err != nil {
		log.Fatalln("设置失败: " + err.Error())
	}

	// 添加备用 DNS 服务器
	cmd = exec.Command("netsh", "interface", "ip", "add", "dns",
		wifiName, secondaryDNS)
	err = cmd.Run()
	if err != nil {
		log.Fatalln("设置失败: " + err.Error())
	}

	// TODO: 刷新DNS设置
	cmd = exec.Command("ipconfig", "/flushdns")
	err = cmd.Run()
	if err != nil {
		log.Fatalln("设置失败: " + err.Error())
	}

	return true, nil
}

// Setting 主程序设置
func Setting(url string) bool {
	// 获取IP
	ip, err := Login(url)
	if err != nil {
		log.Fatalln("设置失败: " + err.Error())
	}

	// 设置 IP
	result, err := NetSet(ip)
	if err != nil {
		log.Fatalln("设置失败: " + err.Error())
	}
	return result
}

func main() {
	// 获取并设置
	var url string

	// 参数判断
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	if flag := Setting(url); flag {
		log.Println("Set successful!")
		if loginFlag, Info, err := check.LoginCheck(); loginFlag {
			fmt.Println(Info.IPAddress)
		} else {
			fmt.Println(err.Error())
		}
	}
}
