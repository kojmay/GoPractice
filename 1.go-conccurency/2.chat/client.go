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
	msgChan chan string
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
	// todo : 交互
	_, err = conn.Write([]byte(nickname))
	if !jugeErr(err, "conn.Write nickname") {
		return
	}
	defer conn.Close()

	// 用户输入
	go inputFunc(conn)
	// 从服务器接收
	go readFromServerFun(conn)
	// go writeFun(conn)

	for msg := <- msgChan {
		fmt.Println(msg)
	}

	// //2. chat with server
	var userInput string
	for {
		fmt.Println("Input:")
		fmt.Scan(&userInput)
		//todo: sigle routine, change to mutiple routines to simulate more chat situations
		_, err = conn.Write([]byte(userInput))
		if !jugeErr(err, "conn.Write userInput") {
			return
		}

		// read from server
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if !jugeErr(err, "conn.Read from server") {
			return
		}
		fmt.Println(string(buf[:]))
	}

}

func inputFunc(conn net.Conn) {
	// for msg := range readChan {
	// 	fmt.Println(msg)
	// }

	var userInput string
	for {
		fmt.Scan(&userInput)
		//todo: sigle routine, change to mutiple routines to simulate more chat situations
		_, err = conn.Write([]byte(userInput))
		if !jugeErr(err, "conn.Write userInput") {
			continue
		}
		// msgChan <- userInput
		// 用户输入

	}
}

func readFromServerFun(conn net.Conn) {
	for {
		// read from server
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if !jugeErr(err, "conn.Read from server") {
			return
		}
		// fmt.Println(string(buf[:]))
		msgChan <- string(buf[:])
	}
}

func main() {
	startChat("127.0.0.1:8000")
}
