package main

import (
	"fmt"
	"time"

	wifi "github.com/mark2b/wpa-connect"
)

var (
	ssid     string
	password string
)

func main() {

	netCount, err := listNetworks()
	must(err)

	if netCount > 0 {
		fmt.Println("Please enter the ssid of the network")
		fmt.Scanln(&ssid)

		fmt.Println("Please enter the password of the network")
		fmt.Scanln(&password)

		err := connectToNetwork(ssid, password)
		must(err)
		return
	}
	fmt.Println("There are no networks to connect to. Exiting...")
}

func listNetworks() (int, error) {
	bssList, _ := wifi.ScanManager.Scan()
	// must(err)
	fmt.Println("Network List")
	fmt.Println("--------------------------------------------------")
	for _, bss := range bssList {
		print(bss.SSID, bss.Signal, bss.KeyMgmt)
	}
	return len(bssList), nil
}

func connectToNetwork(ssid, password string) error {
	wifi.SetDebugMode()
	if conn, err := wifi.ConnectManager.Connect(ssid, password, time.Second*60); err == nil {
		fmt.Println("Connected", conn.NetInterface, conn.SSID, conn.IP4.String(), conn.IP6.String())
		return nil
	} else {
		return err
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
