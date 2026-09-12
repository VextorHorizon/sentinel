package main

import (
	"fmt"
	"net"
	"strconv"
)

type IPAddress struct {
	ip   string
	port int
}

func main() {

	IPPack := []IPAddress{
		{ip: "127.0.0.1", port: 5354},
		{ip: "127.0.0.1", port: 5353},
	}
	// in the future we will append IPA to IPPack as user input

	for _, Host := range IPPack {
		target := Host.ip + ":" + strconv.Itoa(Host.port)
		// fmt.Printf("index is %d and ip is %s and port is %d \n", i, IP.ip, IP.port)

		fmt.Println(target)

		conn, err := net.Dial("tcp", target)
		fmt.Printf("Scanning %s: ", target)

		if err != nil {
			fmt.Printf("Port %d is closed! \n", Host.port)
			return
		}

		defer conn.Close()

		if conn != nil {
			fmt.Printf("Port %d is open! \n", Host.port)
		}

	}

	// target := IPPack
	// address := target.ip + ":" + strconv.Itoa(target.port)

	// Itoa = Integer to ASCII

	// 	conn, err := net.Dial("tcp", address) // connecting a connection to the port from address(ip)

	// 	if err != nil {
	// 		fmt.Printf("Port %d is closed!", target.port)
	// 	}

	// 	// fmt.Println(conn)

	// 	if conn != nil {
	// 		fmt.Printf("Port %d is open!", target.port)
	// 	}
}
