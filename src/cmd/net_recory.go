package cmd

import (
	"os/exec"
)

func DefaultNet() string {
	// 执行 netsh 命令以将 IP 设置恢复为 DHCP
	cmd := exec.Command("netsh", "interface", "ip", "set", "address", "name=WLAN", "source=dhcp")
	if err := cmd.Run(); err != nil {
		return "恢复失败"
	}

	// 执行 netsh 命令以将 DNS 设置恢复为 DHCP
	cmd1 := exec.Command("netsh", "interface", "ip", "set", "dns", "name=WLAN", "source=dhcp")
	if err := cmd1.Run(); err != nil {
		return "恢复失败"
	}

	return "恢复成功"
}
