package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oliver006/deco"
)

const applicationVersion = "1.0.4"

func main() {
	host := flag.String("host", "192.168.1.1", "The host address of the Deco API")
	password := flag.String("password", "", "The password for authentication")
	version := flag.Bool("version", false, "Print the application version")
	flag.Parse()

	if *version {
		fmt.Printf("Application version: %s\n", applicationVersion)
		return
	}
	if *password == "" {
		fmt.Println("Usage: usage --host 192.168.1.1 --password router_password")
		fmt.Println("Usage: usage --version")
		os.Exit(1)
	}

	c := deco.New(*host)
	err := c.Authenticate(*password)
	if err != nil {
		log.Fatal(err.Error())
	}

	printPerformance(c)
	printDevices(c)
	printDecos(c)
	printDevicesByDeco(c)
}

func printPerformance(c *deco.Client) {
	fmt.Println("[+] Performance")
	result, err := c.Performance()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Print response as json
	jsonData, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonData))
}

func printDevices(c *deco.Client) {
	fmt.Println("[+] Clients")
	result, err := c.ClientList()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, device := range result.Result.ClientList {

		fmt.Printf("%s\tOnline: %t\n", device.Name, device.Online)
	}
}

func printDevicesByDeco(c *deco.Client) {
	fmt.Println("[+] Clients by Deco")
	result, err := c.DeviceList()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, device := range result.Result.DeviceList {
		deviceName := device.Nickname
		if deviceName == "" {
			deviceName = device.DeviceIP
		}
		fmt.Printf("%s (%s)\n", deviceName, device.MAC)

		clients, err := c.ClientListForDevice(device.MAC)
		if err != nil {
			log.Printf("failed to fetch clients for %s: %s", device.MAC, err.Error())
			continue
		}
		for _, client := range clients.Result.ClientList {
			fmt.Printf("  %s\tOnline: %t\n", client.Name, client.Online)
		}
	}
}

func printDecos(c *deco.Client) {
	fmt.Println("[+] Devices")
	result, err := c.DeviceList()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, device := range result.Result.DeviceList {
		fmt.Printf("%s\tStatus: %s\n", device.DeviceIP, device.InetStatus)
	}
}
