package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type JSONData struct {
	Content Content `json:"content"`
	Limit   int     `json:"limit"`
}

type Content struct {
	ScanIP string `json:"ip"`
	Mac    string `json:"mac"`
	TTL    string `json:"ttl"`
}

func Login(ip *string) (err error) {
	var jsonData JSONData

	urlIp := "http://localhost:8080/api/v1/ip/search"

	// 获取数据
	resp, err := http.Get(urlIp)
	if err != nil {
		return err
	}

	// 延迟关闭
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// 读取数据
	info, _ := io.ReadAll(resp.Body)

	// JSON 格式解析
	if err := json.Unmarshal([]byte(info), &jsonData); err != nil {
		return err
	}

	// 获取IP地址
	scanIP := jsonData.Content.ScanIP

	// 输出
	*ip = string(scanIP)
	return nil
}

func winNetSet(ip string) error {
	var area string
	areaInfo := strings.Split(ip, ".")[1]

	if info, _ := strconv.Atoi(areaInfo); info > 25 {
		area = "25"
	} else {
		area = "15"
	}

	subnetMask := "255.255.240.0"
	gateway := fmt.Sprintf("10.%s.128.1", area)
	wifiName := "WLAN"
	primaryDNS := "10.100.10.100"
	secondaryDNS := "119.29.29.29"

	// 设置本地 IP 地址
	cmd := exec.Command("netsh", "interface", "ip", "set", "address",
		"name="+wifiName, "source=static", ip, subnetMask, gateway)
	err := cmd.Run()
	if err != nil {
		return errors.New("设置失败")
	}

	// 设置主 DNS 服务器
	cmd = exec.Command("netsh", "interface", "ip", "set", "dns",
		wifiName, "static", primaryDNS, "primary")
	err = cmd.Run()
	if err != nil {
		return errors.New("设置失败")
	}

	// 添加备用 DNS 服务器
	cmd = exec.Command("netsh", "interface", "ip", "add", "dns",
		wifiName, secondaryDNS)
	err = cmd.Run()
	if err != nil {
		return errors.New("设置失败")
	}

	// TODO: 刷新DNS设置
	cmd = exec.Command("ipconfig", "/flushdns")
	err = cmd.Run()
	if err != nil {
		return errors.New("设置失败")
	}

	return nil
}

// NetSet 主程序设置
func NetSet(ip string) string {
	// 获取IP
	if err := Login(&ip); err != nil {
		return "伪装失败"
	}
	// 设置 IP
	if err := winNetSet(ip); err == nil {
		return "伪装成功"
	} else {
		return "伪装失败"
	}
}
