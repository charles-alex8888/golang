package main

import (
	"fmt"
	"time"
)

func main() {
	numCh := make(chan int)

	// 超时控制
	// timeout := make(chan int, 1)
	// go func() {
	// 	time.Sleep(time.Second)
	// 	timeout <- 1
	// }()

	go func(ch chan int) {
		ch <- 1
	}(numCh)

	strCh := make(chan string)
	go func(ch chan string) {
		ch <- "hello"
	}(strCh)

	time.Sleep(time.Second)
	// select 监听所有channel
	for i := 0; i < 10; i++ {
		select {
		case <-numCh:
			fmt.Println("read numCh")
		case <-strCh:
			fmt.Println("read strCh")
		// case <-timeout:
		// 	fmt.Println("timeout!")
		// 超时控制
		case <-time.After(time.Second / 2):
			fmt.Println("timeout!")
		default:
			fmt.Println("no goroutine is ready")
		}
	}
}
