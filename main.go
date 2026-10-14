package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type IPAddress struct {
	ip   string
	port []int
}

func main() {

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
	return finishParser, err //[1 2 3 4 5 6]
}

func strCutParser(strCut string) []string {
	//[1-12]
	strCutal := strings.Split(strCut, "-")
	return strCutal

}

type scanResult struct {
	Port   int
	Result bool
}

func ScannerDoor(target IPAddress) ([]int, []int, error) {
	var err error
	var wg sync.WaitGroup
	ch := make(chan scanResult)

	openPort := []int{}
	closePort := []int{}

	for _, port := range target.port {
		if port > 65535 || port < 1 {
			portString := strconv.Itoa(port)
			return nil, nil, fmt.Errorf("Reject %s:%s, Port must between 1 and 65535", target.ip, portString)
		}

		host := target.ip + ":" + strconv.Itoa(port) // String | Itoa = Integer to ASCII
		wg.Add(1)
		go portScanner(host, ch, port, &wg)

	}

	// wg need to be chagning a real value number of wg in global (point to change in real one, not copy)

	go func() {
		wg.Wait()
		close(ch)
	}()
	/*
		close Channel after all goroutine done the work.
		Because of range ch will stop.
		If not, it will deadlock cuz ch waiting another value from ch<- that already finished and has no value left
	*/

	for result := range ch {

		if result.Result == false {
			closePort = append(closePort, result.Port)
		}
		if result.Result == true {
			openPort = append(openPort, result.Port)
		}
	} // future we need to arranging from smallest to largest

	return openPort, closePort, err
}

// finished feat Concurrency feature. It took me for 4 days to learn about *channel*, sync.WaitGroup and complete mental model of GOROUTINE

func portScanner(host string, ch chan scanResult, port int, wg *sync.WaitGroup) {

	defer wg.Done()

	conn, err := net.Dial("tcp", host)
	if err != nil { //Connection unsuccess
		ch <- scanResult{
			Port:   port,
			Result: false,
		}

	}
	if conn != nil { //Connection success
		conn.Close()
		ch <- scanResult{
			Port:   port,
			Result: true,
		}

	}
}

// method use as sword, function use as put var in to blender
// for those who read this. I will say, YES
