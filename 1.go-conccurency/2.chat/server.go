package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// msg struct
type Message struct {
	owner   string
	content string
}

var (
	allClients    map[string]Client // store all user connections
	broadcastChan chan Message      //broadcast chan
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
	broadcastChan = make(chan Message)

	//connect to the client
	for {
		conn, err := listener.Accept()
		if !jugeErr(err, "listener.Accept") {
			continue
		}

		// handle clients connections
		go handleConn(conn)
	}

}

// user info struct
type Client struct {
	nickname   string
	ID         int
	ipAdd      string
	clientChan chan string
	joinTime   time.Time
	// conn     net.Conn
}

func (u *Client) String() string {
	return u.ipAdd + ", nickname:" + u.nickname + ", Enter At:" +
		u.joinTime.Format("2006-01-02 15:04:05+8000")
}

// handle clients' connection
func handleConn(conn net.Conn) {
	nickname := make([]byte, 2048)
	_, err := conn.Read(nickname)
	if !jugeErr(err, "server conn.Read nickname") {
		return
	}
	// fmt.Println("New client entered: " + string(nickname))

	// var client Client
	// client.nickname = string(nickname[:n])
	// client.ipAdd = conn.RemoteAddr().String()
	// client.clientChan = make(chan string)
	// client.joinTime = time.Now()

	client := Client{
		nickname:   string(nickname),
		ID:         genUserID(),
		ipAdd:      conn.RemoteAddr().String(),
		clientChan: make(chan string),
		joinTime:   time.Now(),
	}
	allClients[client.ipAdd] = client

	// 开始广播监听
	go broadcast()

	// communicate with client
	go readFromClient(client, conn)
	go writeToClient(client, conn)

	// broadcast new client
	broadcastChan <- Message{"", "New joiner, nickname:" + client.nickname + "(ip:" + client.ipAdd + ")" + ", welcome!\n"}
}

// read from client
func readFromClient(client Client, conn net.Conn) {

	fmt.Println("start read from client")
	msg := make([]byte, 2048)
	for {
		n, err := conn.Read(msg)
		if !jugeErr(err, "conn.Read") { // 断开 or 结束
			delete(allClients, client.ipAdd)
			conn.Close()
			break
		}
		fmt.Println(string(msg[:]))
		broadcastChan <- Message{client.ipAdd, "\t\t " + client.nickname + "(" + client.ipAdd + "):" + string(msg[:n])}
	}
}

// broadcast when user online and offline
func broadcast() {
	for {
		msg := <-broadcastChan
		fmt.Println("bradcast:", msg.content)
		for _, client := range allClients {
			// fmt.Println(" client.ipAdd != msg.owner", client.ipAdd != msg.owner)
			if client.ipAdd != msg.owner {
				client.clientChan <- msg.content
			}
		}
	}
}

// write to client
func writeToClient(client Client, conn net.Conn) {
	for {
		msg := <-client.clientChan
		conn.Write([]byte(msg))
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

var (
	globalID int
	idLocker sync.Mutex
)

// generate user id
func genUserID() int {
	idLocker.Lock()
	defer idLocker.Unlock()
	globalID++
	return globalID
}

func main() {

	startListenning("127.0.0.1:8000")
	// go broadcast()
}
