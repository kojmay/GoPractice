package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Listening
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	// connnected to client
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listenner.Accept err:", err)
		return
	}
	defer conn.Close()

	//1. read file name
	buf := make([]byte, 1024*4)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fileName := string(buf[:n])

	//2. return "OK"
	conn.Write([]byte("ok"))

	//3. start to recive file
	start := time.Now()

	RecieveFile(fileName, conn)

	elapsed := time.Since(start)
	fmt.Println("接收端传输耗时：", elapsed)
}

func RecieveFile(fileName string, conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File Recieved COMPLETED!")
			} else {
				fmt.Println("conn.Read err:", err)
			}
			return
		}
		if n == 0 {
			fmt.Println("File Recieved COMPLETED!")
			return
		}
		f.Write(buf[:n])
	}
}
