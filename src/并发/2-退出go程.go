package main

import (
	"fmt"
	"runtime"
	"time"
)

// return  返回当前函数
// exit   退出当前进程
// GOEXIT 提前退出当前go程

func main() {
	go func() {
		func() {
			fmt.Println("这是子go程")
			// return    // 退出当前函数
			// os.Exit() // 退出进程
			runtime.Goexit() // 退出当前go程
		}()
		fmt.Println("子go程结束")
	}()

	fmt.Println("这是主go程")
	time.Sleep((5 * time.Second))
	fmt.Println("over!")
}
