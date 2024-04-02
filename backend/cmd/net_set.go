package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func WinNetSet(ip string) error {
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
		wifiName, secondaryDNS,
	)
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
