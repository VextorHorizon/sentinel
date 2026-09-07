package main

import (
	"fmt"
	"net" //net.Dial
	"strconv"
)

func main() {

	ip := "127.0.0.1"
	port := 5354

	address := ip + ":" + strconv.Itoa(port)

	SetofIPAddress := []string{}

	SetofIPAddress = append(SetofIPAddress, address)

	fmt.Println(SetofIPAddress)

	// Itoa = Integer to ASCII

	// fmt.Println(address)

	conn, err := net.Dial("tcp", address) // connecting a connection to the port from address(ip)

	if err != nil {
		fmt.Printf("Port %d is closed!", port)
	}

	// fmt.Println(conn)

	if conn != nil {
		fmt.Printf("Port %d is open!", port)
	}
}
