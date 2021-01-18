package main

import (
	"fmt"
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

func main() {
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
