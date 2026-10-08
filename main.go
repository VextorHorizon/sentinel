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

func main() { //future refactor: sub-function should return err to decide it on main function

	UserInput, err := GetUserInput()
	if err != nil {
		fmt.Println(err)
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
		err = errors.New("Error, need format: sentinel <ip> <port>")
		return IPAddress{}, err
	} //first debugging without AI that I ever done!

	var userInput IPAddress
	userInput.ip = args[1]
	//userInputports = [1,2,3,4,5 7,8,9,6,5] string slice
	// strPort = [1 2 3 4 5 7 8 9 6 5]
	userInputports := strings.Split(args[2], ",")
	var parser []int
	for _, strPort := range userInputports {
		//str Port that is still String type
		if strings.Contains(strPort, "-") { // if the input is parser(1-12), return []int
			parser, err = strParser(strPort)
			if err != nil {
				return IPAddress{}, err
			}
			for _, i := range parser {
				userInput.port = append(userInput.port, i)
			}
			continue
		}

		intPort, err := strconv.Atoi(strPort)
		if err != nil {
			return IPAddress{}, fmt.Errorf("Invalid port: %s", strPort)
		}

		//Port is integer now

		userInput.port = append(userInput.port, intPort)

	}
	return userInput, err
}

func strParser(strPortUserInput string) ([]int, error) {
	var err error
	var parser []int
	strParser := strCutParser(strPortUserInput) // [1 12]
	for _, singleStr := range strParser {
		singleInt, err := strconv.Atoi(singleStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid port range format: %s", singleStr)
		}
		parser = append(parser, singleInt) // put int to parser[]
	}

	if len(parser) > 2 { // validate if input is more than just two number for parser
		return nil, fmt.Errorf("Invalid port range: %d", parser) // for future refactor, should return err and lets the main decide on their own
	}
	if parser[0] >= parser[1] {
		if parser[0] == parser[1] {
			return nil, fmt.Errorf("Invalid parser format(%d)", parser[0])
		}
		return nil, fmt.Errorf("Invalid parser format(%d-%d) to (%d-%d)?",
			parser[0], parser[1],
			parser[1], parser[0])
	}
	parser[1] += 1 //as parser[1] normally it be -1 IDK WHY
	var finishParser []int
	for i := parser[0]; i < parser[1]; i++ {
		finishParser = append(finishParser, i)

	}
	// fmt.Println(finishParser)
	return finishParser, err //[1 2 3 4 5 6]
}

func strCutParser(strCut string) []string {
	//[1-12]
	strCutal := strings.Split(strCut, "-")
	return strCutal

}

func Scanner(target IPAddress) {

	for _, port := range target.port {

		if port > 65535 || port < 1 {
			portString := strconv.Itoa(port)
			fmt.Printf("Reject %s:%s, Port must between 1 and 65535", target.ip, portString)
			continue

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
