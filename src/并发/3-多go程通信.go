package main

import "fmt"

func main() {
	// 多go程 使用管道进行通信，不需要加解锁

	// 创建管道 channel,和map一样必须使用make
	numChan := make(chan int, 10)
	// string := make(chan string)

	go func() {
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("这是子go程，读取的数据为:", data)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			// 向管道写入数据
			numChan <- i
			fmt.Println("这是主go程，写入数据:", i)
		}
	}()
	//
	for i := 20; i < 50; i++ {
		// 向管道写入数据
		numChan <- i
		fmt.Println("这是主go程，写入数据:", i)
	}

}
