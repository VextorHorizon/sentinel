package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
)

type IPAddress struct {
	ip   string
	port int // Future this going to be []int slice, because it need to handle for multiple port scaning with single IP
}

func main() {

	UserInput, err := GetUserInput()
	if err != nil {
		return
	}
	IPPack := []IPAddress{}
	// {ip: "127.0.0.1", port: 5354},
	// {ip: "8.8.8.8", port: 60}, // this connecting to outside of the world, slow.

	IPPack = append(IPPack, UserInput) //need IPAddress struct to work

	for _, Host := range IPPack {
		Scanner(Host)
	} // it run sequential. if 2 is connecting slow, all the process is slow. future we going to add goroutine for sure
}

func GetUserInput() (IPAddress, error) {

	var err error
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("Error, need format: sentinel <ip> <port>")
		err = errors.New("mismatch user input format")
		return IPAddress{}, err
	}

	var userInput IPAddress
	userInput.ip = args[1]
	userInput.port, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Port missed type")
		return IPAddress{}, err
	}
	return userInput, err
}

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
		conn.Close() // we just checking target connection, no need to leave the door open
	}
}

// method use as sword, function use as put var in to blender
// for those who read this. I will say, YES
