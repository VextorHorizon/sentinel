package main

import (
	"fmt"
	"net" //net.Dial
)

func main() {

	IPAddress := "127.0.0.1"

	conn, err := net.Dial("tcp", IPAddress)

	if err != nil {
		fmt.Println("Error from conn")
	}
}
