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

	var err error

	args := os.Args

	if len(args) < 3 {
		fmt.Printf("Error, need format: sentinel <ip> <port>")
		return
	}

	var userInput IPAddress
	userInput.ip = args[1]
	userInput.port, err = strconv.Atoi(args[2]) //args[2] receive as String
	if err != nil {
		fmt.Println("Port missed type")
		return
	}

	IPPack := []IPAddress{}
	// {ip: "127.0.0.1", port: 5354},
	// {ip: "127.0.0.1", port: 5353},
	// {ip: "8.8.8.8", port: 60}, // this connecting to outside of the world, slow.

	IPPack = append(IPPack, userInput) // userInput, If there no port(args[2]) it will error. Because loop is checking target(ip, port)
	// next time it better that seperate target as ip and port. If it better! Check again!

	for _, Host := range IPPack {
		Scanner(Host)
	} // it run sequential. if 2 is connecting slow, all the process is slow. future we going to add goroutine for sure!
}

// For sure to lets you know, conn is opening!

func Scanner(target IPAddress) {

	if target.port > 65535 || target.port < 1 {
		portString := strconv.Itoa(target.port)
		fmt.Printf("Reject %s:%s, Port must between 1 and 65535", target.ip, portString)
		return
	}

	host := target.ip + ":" + strconv.Itoa(target.port) // String | Itoa = Integer to ASCII
	fmt.Printf("\n%s \n", host)

	conn, err := net.Dial("tcp", host) //Open the connection between the target and ourself
	if err != nil {
		fmt.Printf("Port %d is closed! \n \n", target.port)
	}

	if conn != nil {
		fmt.Printf("Port %d is open! \n \n", target.port)
	}

} // single scanning

// method use as sword, function use as put var in to blender
// for those who read this. I will say, YES
