package main

import (
	"net" //net.Dial
)

func main() {

	IPAddress := "127.0.0.1"

	conn, err := net.Dial("tcp", IPAddress)
}
