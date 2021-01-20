package main

import (
	"fmt"
	"net"
	"time"
)

// user info struct
type Client struct {
	nickname   string
	ipAdd      string
	clientChan chan string
	joinTime   time.Time
	// conn     net.Conn
}

var (
	allClients    map[string]Client // store all user connections
	broadcastChan chan string       //broadcast chan
)

// 1.start listenning
func startListenning(serverURL string) {
	// Listenning
	listener, err := net.Listen("tcp", serverURL)
	if !jugeErr(err, "net.Listen") {
		return
	}
	defer listener.Close()

	allClients = make(map[string]Client)

	//connect to the client
	for {
		conn, err := listener.Accept()
		if !jugeErr(err, "listener.Accept") {
			continue
		}

		go handleConn(conn)
	}

}

// handle clients' connection
func handleConn(conn net.Conn) {
	defer conn.Close()

	nickname := make([]byte, 1024)
	_, err := conn.Read(nickname)
	if !jugeErr(err, "server conn.Read nickname") {
		return
	}

	var client Client
	client.nickname = string(nickname)
	client.ipAdd = conn.RemoteAddr().String()
	// client.conn = conn
	client.clientChan = make(chan string)
	client.joinTime = time.Now()
	allClients[client.ipAdd] = client

	// communicate with client
	go communicateWithClient(client, conn)

	// 开始广播监听
	go broadcast()

	// broadcast new client
	broadcastChan <- "A new client joined, \tip:" + client.ipAdd + "\tnickname" + client.nickname
}

// broadcast when user online and offline
func broadcast() {
	for {
		msg := <-broadcastChan
		for _, client := range allClients {
			// _, err := client.conn.Write([]byte(msg))
			// if !jugeErr(err, "broadcast to client "+k) {
			// 	return
			// }
			client.clientChan <- msg
		}
	}
}

// communicate with client
func communicateWithClient(client Client, conn net.Conn) {

	for msg := range client.clientChan {
		conn.Write([]byte(msg + "\n"))
	}
}

// judge err
func jugeErr(err error, prompt string) bool {
	if err != nil {
		fmt.Println(prompt, ", err info: ", err)
		return false
	}
	return true
}

func main() {

	startListenning("127.0.0.1:8000")
	// go broadcast()
}
