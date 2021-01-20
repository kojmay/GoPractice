package main

import (
	"fmt"
	"net"
)

func jugeErr(err error, prompt string) bool {
	if err != nil {
		fmt.Println(prompt, ", err info: ", err)
		return false
	}
	return true
}

var (
	readChan  chan string
	writeChan chan string
)

// start to chat
func startChat(serverURL string) {
	fmt.Println("Please input your nickname:")
	var nickname string
	fmt.Scan(&nickname)

	//connect to the server
	conn, err := net.Dial("tcp", serverURL)
	if !jugeErr(err, "net.Dial") {
		return
	}

	//1. send nick name first
	_, err = conn.Write([]byte(nickname))
	if !jugeErr(err, "conn.Write nickname") {
		return
	}

	defer conn.Close()

	go readFun(conn)
	go writeFun(conn)

	// //2. chat with server
	// var userInput string
	// for {
	// 	fmt.Println("Input:")
	// 	fmt.Scan(&userInput)
	// 	//todo: sigle routine, change to mutiple routines to simulate more chat situations
	// 	_, err = conn.Write([]byte(userInput))
	// 	if !jugeErr(err, "conn.Write userInput") {
	// 		return
	// 	}

	// 	// read from server
	// 	buf := make([]byte, 1024)
	// 	_, err = conn.Read(buf)
	// 	if !jugeErr(err, "conn.Read from server") {
	// 		return
	// 	}
	// 	fmt.Println(string(buf[:]))
	// }

}

func readFunc(conn net.Conn) {
	for msg := range readChan {
		fmt.Println(msg)
	}
}

func writeFunc(conn net.Conn) {
	for msg := range writeChan {
		conn.Write([]byte(msg))
	}
}

func main() {
	startChat("127.0.0.1:8000")
}
