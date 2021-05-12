package main

import (
	. "fmt"
	"strconv"
	. "time"
)

func Read(ch chan int) {
	value := <-ch
	Println("value: " + strconv.Itoa(value))
}

func Write(ch chan int) {
	ch <- 10
}

func Add(x, y int, quit chan int) {
	r := x + y
	Println(r)
	quit <- 1
}

func main() {
	/*
		协程间通信使用channel
	*/
	// ch := make(chan int, 10)

	// go Write(ch)
	// go Read(ch)
	Sleep(Second)

	numChan := make([]chan int, 1000)
	for i := 0; i < 1000; i++ {
		numChan[i] = make(chan int)
		go Add(i, i, numChan[i])
	}

	// 协程没有执行完成之前，以下代码会阻塞
	for _, v := range numChan {
		<-v
	}

	Println("主程序结束!")
}
