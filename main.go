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
	//userInputports = [1,2,3,4,5 7,8,9,6,5] string slice
	// strPort = [1 2 3 4 5 7 8 9 6 5]
	userInputports := strings.Split(args[2], ",")
	var parser []int
	for _, strPort := range userInputports {
		//str Port that is still String type
		if strings.Contains(strPort, "-") { // if the input is parser(1-12)
			parser = strParser(strPort)
			for _, i := range parser {
				userInput.port = append(userInput.port, i)
			}
			continue
		}

		intPort, err := strconv.Atoi(strPort)
		if err != nil {
			fmt.Println("Port missed type, port need to be Integer")
			continue
		}

		//Port is integer now

		userInput.port = append(userInput.port, intPort)

	}
	return userInput, err
}

func strParser(strPortUserInput string) []int {
	var parser []int
	strParser := strCutParser(strPortUserInput) // [1 12]
	for _, singleStr := range strParser {
		singleInt, err := strconv.Atoi(singleStr)
		if err != nil {
			fmt.Println("Error: incorrect format from Parser")
			return nil
		}
		parser = append(parser, singleInt) // put int to parser[]
	}

	if len(parser) > 2 { // validate if input is more than just two number for parser
		fmt.Println("Uncorrect parser format")
		return nil // for future refactor, should return err and lets the main decide on their own
	}
	if parser[0] >= parser[1] {
		fmt.Printf("Incorrect parser format(%d-%d) to (%d-%d)?",
			parser[0], parser[1],
			parser[1], parser[0])
		return nil
	}
	parser[1] += 1 //as parser[1] normally it be -1 IDK WHY
	var finishParser []int
	for i := parser[0]; i < parser[1]; i++ {
		finishParser = append(finishParser, i)

	}
	// fmt.Println(finishParser)
	return finishParser //[1 2 3 4 5 6]
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
