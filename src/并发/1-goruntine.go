package main

import (
	"fmt"
	"time"
)

/*
go程===> goroutine
go程占用的系统资源远远小于线程，一个go程序大约需要4k-5k的内存资源
一个程序可以启动大量的go程，对于实现高并发性能很好
只需要在目标函数钱加上go关键字即可
*/

func dispaly(num int) {
	count := 1
	for {
		fmt.Println("当前是子go程", num, "count值为", count)
		count++
	}
}

func main() {
	// 启动子go程
	for i := 0; i < 3; i++ {
		go dispaly(i)
	}
	// 主go程
	count := 1
	for {
		fmt.Println("主go程", count)
		count++
		time.Sleep(1 * time.Second)
	}

}
