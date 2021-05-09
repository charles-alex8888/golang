package main

import (
	"fmt"
	"os"
)

func readFile(filename string) {
	f1, err := os.Open(filename)

	// defer f1.Close()
	defer func() {
		_ = f1.Close()
		fmt.Println("文件已关闭")
	}()

	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	buf := make([]byte, 1024)
	n, err := f1.Read(buf)
	fmt.Println("文件长度为：", n)
	fmt.Println(string(buf))
}

func main() {
	/*
	 1。延迟: 确保某语句在当前栈退出时执行
	 2。一般用于资源管理工作
	 3。解锁、关闭文件
	 4。在同一个函数中多次调用defer时，执行类似栈的机制： 后入先出
	*/
	readFile("5-string.go")
}
