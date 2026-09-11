package main

import (
	"fmt"
	"net" //net.Dial
	"strconv"
)

type IPAddress struct {
	ip   string
	port int
}

func main() {

	IPPack := []IPAddress{
		{ip: "127.0.0.1", port: 5354},
	}
	// in the future we will append IPA to IPPack as user input

	target := IPPack
	address := target.ip + ":" + strconv.Itoa(target.port)

	// SetofIPAddress := []string{}
	// SetofIPAddress = append(SetofIPAddress, address)
	// fmt.Println(SetofIPAddress)
	// Itoa = Integer to ASCII

	conn, err := net.Dial("tcp", address) // connecting a connection to the port from address(ip)

	if err != nil {
		fmt.Printf("Port %d is closed!", target.port)
	}

	// fmt.Println(conn)

	if conn != nil {
		fmt.Printf("Port %d is open!", target.port)
	}
}
