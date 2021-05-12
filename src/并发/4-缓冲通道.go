package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		1. 当缓冲写满的时候，写阻塞，当被读取后，恢复写入
		2. 当缓冲区读完毕，读阻塞
		3. 如果管道没有使用make分配空间，那么管道默认是nil的，读取写入都会阻塞
		4. 对于一个管道 读与写必须对等
		5. 如果阻塞在主go程，程序回崩溃
		6. 如果阻塞在子fo程，会出现内存泄漏
		7. 关闭管道的一方必须是写数据
		8。 关闭一个已经关闭的管道/向关闭的管道写数据，程序会崩溃
	*/

	namesChan := make(chan string, 10)

	go func() {
		fmt.Println("name:", <-namesChan)
	}()

	namesChan <- "hello"
	time.Sleep(1 * time.Second)

	numsChan := make(chan int, 10)
	go func() {
		for i := 0; i < 50; i++ {
			numsChan <- i
			fmt.Println("写入数据：", i)
		}
		fmt.Println("数据写入完毕")
		close(numsChan)
	}()

	for v := range numsChan {
		fmt.Println("读取数据：", v)
	}

	// 判断管道是否关闭

	_, ok := <-numsChan
	if !ok {
		fmt.Println("管道已经关闭")
	}

}
