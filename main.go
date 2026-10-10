package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type IPAddress struct {
	ip   string
	port []int
}

func main() { //future refactor: sub-function should return err to decide it on main function

	timeStart := time.Now()

	UserInput, err := GetUserInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	IPPack := []IPAddress{} //{ip: "127.0.0.1", []port{"5354,60"}}
	// {ip: "8.8.8.8", port: 60}, // this connecting to outside of the world, slow.

	IPPack = append(IPPack, UserInput)
	var openPort []int
	var closePort []int

	for _, Host := range IPPack {
		openPort, closePort, err = ScannerDoor(Host)
		if err != nil {
			fmt.Println(err)
		}
	} // it run sequential. if 2 is connecting slow, all the process is slow. future we going to add goroutine for sure
	for _, port := range openPort {
		fmt.Printf("%d             open\n", port)
	}
	for _, port := range closePort {
		fmt.Printf("%d             closed\n", port)
	}

	fmt.Println("This operation took: ", time.Since(timeStart))
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

func ScannerDoor(target IPAddress) ([]int, []int, error) {
	var err error
	openPort := []int{}
	closePort := []int{}

	for _, port := range target.port {
		if port > 65535 || port < 1 {
			portString := strconv.Itoa(port)
			return nil, nil, fmt.Errorf("Reject %s:%s, Port must between 1 and 65535", target.ip, portString)
		}

		host := target.ip + ":" + strconv.Itoa(port) // String | Itoa = Integer to ASCII

		Result := portScanner(host) //Open the connection between the target and ourself //2
		if Result == false {        //Connection unsuccess
			closePort = append(closePort, port)
		}
		if Result == true { // Connection success
			// we just checking target connection, no need to leave the door open
			openPort = append(openPort, port)
		}
	}
	return openPort, closePort, err
}

func portScanner(host string) bool {

	conn, err := net.Dial("tcp", host)
	if err != nil {
		return false
	}
	if conn != nil {
		conn.Close()
		return true
	}

	return false
}

// method use as sword, function use as put var in to blender
// for those who read this. I will say, YES
