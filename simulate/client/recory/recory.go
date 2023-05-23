package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("netsh", "interface", "ip", "set", "address", "name=WLAN", "source=dhcp")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("IP address set to automatic (DHCP) successfully")
}
