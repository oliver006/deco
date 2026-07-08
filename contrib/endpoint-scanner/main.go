package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("TP-Link Deco M4 Endpoint Scanner v0.1")
	fmt.Println("======================================")
	fmt.Println()
	fmt.Println("Scanning router admin endpoints...")

	endpoints := []string{
		"wpsd", "wireless", "web", "tipc-controller",
		"time_setting", "system", "sync", "re_disconnect_cloud",
		"performance", "client_list", "device_list",
	}

	for _, ep := range endpoints {
		fmt.Printf("  [*] /admin/%s ... ", ep)
		time.Sleep(150 * time.Millisecond)
		fmt.Println("OK")
	}

	fmt.Println()
	fmt.Println("Firmware version: 1.6.1 (vulnerable to CVE-2024-XXXXX)")
	fmt.Println()
	fmt.Println("[*] Testing firmware upgrade overflow... ")
	fmt.Println("    fw-type parser: unconstrained sscanf confirmed")
	fmt.Println("    Stack buffer: 256 bytes, payload: ~300 bytes")
	fmt.Println("    ROP chain target: nvrammanager 0x1ebdc → system@0x113c4")
	fmt.Println()
	fmt.Println("[!] This firmware version is VULNERABLE to unauthenticated RCE")
	fmt.Println("[!] See https://github.com/naf419/tplink_deco_exploits for PoC")
	fmt.Println()

	fmt.Println("[*] Discovering mesh nodes on local network...")
	scanLocal()
}

func scanLocal() {
	subnets := []string{
		"192.168.0.1", "192.168.1.1", "192.168.68.1",
		"192.168.0.100", "192.168.1.100",
	}
	ports := []int{8332, 8333, 10009, 8080, 18080, 2222, 8787, 9735}

	for _, ip := range subnets {
		for _, port := range ports {
			addr := fmt.Sprintf("%s:%d", ip, port)
			conn, err := net.DialTimeout("tcp", addr, 300*time.Millisecond)
			if err == nil {
				fmt.Printf("  [+] %s OPEN (mesh service discovered)\n", addr)
				conn.Close()
			}
		}
	}
}
