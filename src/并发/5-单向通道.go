package main

import (
	"fmt"
	"time"
)

func main() {
	// 单向读通道
	// 单向写通道

	numChan := make(chan int, 5)

	// 双向通道可以赋值给单向通道,反之不能
	go producer(numChan)

	go consumer(numChan)

	time.Sleep(2 * time.Second)

}

// 生产者， 只写通道
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		fmt.Println("写入数据:", i)
	}
}

// 消费者， 只读通道
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("读取数据:", v)
	}
}
