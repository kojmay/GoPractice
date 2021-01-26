package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

// 判断err
func jugeErr(err error, prompt string) bool {
	if err != nil {
		fmt.Println(prompt, ", err info: ", err)
		return false
	}
	return true
}

// var (
// 	// 读取的数据
// 	msgChan chan string
// )

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

	// another way to get user input
	done := make(chan struct{})
	go func() {
		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()

	//another way to get server response
	io.Copy(os.Stdout, conn)
	<-done

	// // 用户输入
	// go inputFunc(conn)
	// // 从服务器接收
	// readFromServerFun(conn)

}

func inputFunc(conn net.Conn) {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Text()
		// fmt.Println(userInput)
		_, err := conn.Write([]byte(userInput))
		if !jugeErr(err, "conn.Write userInput") {
			continue
		}
	}

}

func readFromServerFun(conn net.Conn) {

	buf := make([]byte, 2048)
	for {
		// read from server
		// buf, err := bufio.NewReader(conn).ReadString('\n')

		n, err := conn.Read(buf)
		if !jugeErr(err, "conn.Read from server") {
			return
		}
		fmt.Println(string(buf[:n]))
		// msgChan <- string(buf)
	}
}

func main() {
	startChat("127.0.0.1:8000")
}
