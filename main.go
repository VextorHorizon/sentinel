package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

/* Future:

 */

type IPAddress struct {
	ip   string
	port []int
}

func main() {

	timeStart := time.Now()

	UserInput, err, timeoutDuration, worker, openportCommand := GetUserInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	IPPack := []IPAddress{} //{ip: "127.0.0.1", []port{"5354,60"}}
	// {ip: "8.8.8.8", port: 60}, // this connecting to outside of the world, slow.

	IPPack = append(IPPack, UserInput)
	var openPort []int
	var closePort []int

	for _, Host := range IPPack { //Multi Host
		openPort, closePort, err = ScannerDoor(Host, timeoutDuration, worker)
		if err != nil {
			fmt.Println(err)
		}
	}
	slices.Sort(openPort)
	slices.Sort(closePort)

	for _, port := range openPort {
		fmt.Printf("%d             open\n", port)
	}
	if openportCommand == true {
	}

	if openportCommand == false {
		for _, port := range closePort {
			fmt.Printf("%d             closed\n", port)
		}
	}

	fmt.Println("This operation took: ", time.Since(timeStart))
}

func GetUserInput() (IPAddress, error, int, int, bool) {

	var err error

	/* Flag Argument: go run . -p 3241 192.0.0.1
	Idex = flag.Args()
	0 = 192.0.0.1
	*/

	p := flag.String("p", "", "port number: 1-65535")
	t := flag.Int("t", 3, "timeout duration in second")
	w := flag.Int("w", 500, "amount of Goroutine workers")
	o := flag.Bool("o", false, "command that allow to see just open port")

	flag.Parse()
	args := flag.Args()

	ports := *p
	timeoutDuration := *t
	worker := *w
	openportCommand := *o

	if len(args) < 1 || ports == "" {
		err = errors.New("Error, need format: sentinel <option> <port> <ip>")
		return IPAddress{}, err, timeoutDuration, worker, openportCommand
	} //first debugging without AI that I ever done!

	var userInput IPAddress
	userInput.ip = args[0]
	//userInputports = [1,2,3,4,5 7,8,9,6,5] string slice
	// strPort = [1 2 3 4 5 7 8 9 6 5]
	userInputports := strings.Split(ports, ",")
	var parser []int
	for _, strPort := range userInputports {
		//str Port that is still String type
		if strings.Contains(strPort, "-") { // if the input is parser(1-12), return []int
			parser, err = strParser(strPort)
			if err != nil {
				return IPAddress{}, err, timeoutDuration, worker, openportCommand
			}
			for _, i := range parser {
				userInput.port = append(userInput.port, i)
			}
			continue
		}

		intPort, err := strconv.Atoi(strPort)
		if err != nil {
			return IPAddress{}, fmt.Errorf("Invalid port: %s", strPort), timeoutDuration, worker, openportCommand
		}

		//Port is integer now

		userInput.port = append(userInput.port, intPort)

	}
	return userInput, err, timeoutDuration, worker, openportCommand
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
		return nil, fmt.Errorf("Invalid port range: %d", parser)
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

func ScannerDoor(target IPAddress, timeoutDuration int, worker int) ([]int, []int, error) {
	var err error
	var wg sync.WaitGroup
	ch := make(chan scanResult)

	limit := make(chan bool, worker)

	openPort := []int{}
	closePort := []int{}

	for _, port := range target.port {
		if port > 65535 || port < 1 {
			portString := strconv.Itoa(port)
			return nil, nil, fmt.Errorf("Reject %s:%s, Port must between 1 and 65535", target.ip, portString)
		}

		limit <- true

		host := target.ip + ":" + strconv.Itoa(port) // String | Itoa = Integer to ASCII
		wg.Add(1)
		go portScanner(host, ch, port, &wg, limit, timeoutDuration) // wg need to be chagning a real value number of wg in global (point to change in real one, not copy)

	}

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
	}

	return openPort, closePort, err
}

// finished feat Concurrency feature. It took me for 4 days to learn about *channel*, sync.WaitGroup and complete mental model of GOROUTINE

func portScanner(host string, ch chan scanResult, port int, wg *sync.WaitGroup, limit chan bool, timeoutDuration int) {

	defer wg.Done()
	conn, err := net.DialTimeout("tcp", host, time.Duration(timeoutDuration)*time.Second)
	<-limit
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
