package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibnacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1

	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}

}

func main2() {
	ch := make(chan int)    // num communication
	quit := make(chan bool) // exit ?

	// consumer
	go func() {
		for i := 0; i < 10; i++ {
			num := <-ch
			fmt.Println(num)
		}
		quit <- true
	}()
	// producer
	fibnacci(ch, quit)
}

// before go run, you must hit `redis-server` to wake redis up
func main() {
	// // conn, _ := net.Dial("tcp", "localhost:6379")
	// message := "*3\r\n$3\r\nSET\r\n$1\r\na\r\n$1\r\nb\r\n"

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if ok := scanner.Scan(); !ok {
			break
		}
		fmt.Println(scanner.Text())
	}
	fmt.Println("Scanning ended")
}
