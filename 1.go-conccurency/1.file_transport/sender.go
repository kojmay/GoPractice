package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func SendFile(path string, conn net.Conn) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File tansport COMPLETE!")
			} else {
				fmt.Println("f.Read err:", err)
			}
			return
		}
		conn.Write(buf[:n])

		if n == 0 {
			fmt.Println("File Recieved COMPLETED!")
			return
		}
	}

}

func main() {

	fmt.Println("Please input the file path:")
	var path string
	fmt.Scan(&path)

	// get file name
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}

	// dail to server
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	defer conn.Close()

	// send file name to server
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	var n int
	buf := make([]byte, 1024)

	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	if string(buf[:n]) == "ok" {
		// send file
		start := time.Now()
		SendFile(path, conn)
		elapsed := time.Since(start)
		fmt.Println("发送端传输耗时：", elapsed)
	}

}
