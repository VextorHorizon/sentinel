package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type IPAddress struct {
	ip   string
	port []int
}

func main() {

	UserInput, err := GetUserInput()
	if err != nil {
		return
	}

	IPPack := []IPAddress{} //{ip: "127.0.0.1", []port{"5354,60"}}
	// {ip: "8.8.8.8", port: 60}, // this connecting to outside of the world, slow.

	IPPack = append(IPPack, UserInput)

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
	} //first debugging without AI that I ever done!

	var userInput IPAddress
	userInput.ip = args[1]
	//userInput.port = [1,2,3,4,5 7,8,9,6,5] string slice
	//strPortsSlice = [1,2,3,4,5] [7,8,9,6,5]
	// strPort = [1 2 3 4 5]
	userInputports := []string{args[2]}
	for _, strPortsSlice := range userInputports {
		strPorts := strings.Split(strPortsSlice, ",")

		for _, strPort := range strPorts {
			intPort, err := strconv.Atoi(strPort)

			if err != nil {
				fmt.Println("Port missed type, port need to be Integer")
				return IPAddress{}, err
			}

			userInput.port = append(userInput.port, intPort)

		}
	}
	return userInput, err
}

func Scanner(target IPAddress) {

	for _, port := range target.port {

		if port > 65535 || port < 1 {
			portString := strconv.Itoa(port)
			fmt.Printf("Reject %s:%s, Port must between 1 and 65535", target.ip, portString)

		}

		host := target.ip + ":" + strconv.Itoa(port) // String | Itoa = Integer to ASCII
		fmt.Printf("\n%s \n", host)

		conn, err := net.Dial("tcp", host) //Open the connection between the target and ourself
		if err != nil {                    //Connection unsuccess
			fmt.Printf("Port %d is closed! \n \n", port)
		}

		if conn != nil { // Connection success
			fmt.Printf("Port %d is open! \n \n", port)
			conn.Close() // we just checking target connection, no need to leave the door open
		}

	}

}

// method use as sword, function use as put var in to blender
// for those who read this. I will say, YES
