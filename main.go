package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type IPAddress struct {
	ip   string
	port int
}

func main() {

	args := os.Args

	var userInput IPAddress

	userInput.ip = args[1]
	userInput.port, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Port missed type")
	}

	IPPack := []IPAddress{
		{ip: "127.0.0.1", port: 5354},
		{ip: "127.0.0.1", port: 5353},
	}

	IPPack = append(IPPack)

	for _, Host := range IPPack {
		Scanner(Host)
	}
}

// For sure to lets you know, conn is opening!

func Scanner(target IPAddress) {

	host := target.ip + ":" + strconv.Itoa(target.port) // Itoa = Integer to ASCII
	fmt.Println(host)

	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Printf("Port %d is closed! \n", target.port)
	}

	if conn != nil {
		fmt.Printf("Port %d is open! \n", target.port)
	}

} // single scanning

// method use as sword, function use as put var in to blender
// for those who read this. I will say, YES
