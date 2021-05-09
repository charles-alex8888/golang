package main

import (
	"fmt"
	"os"
)

func main() {
	// 命令行输入
	cmds := os.Args
	for k, v := range cmds {
		fmt.Println(k, v)
	}

	if len(cmds) < 2 {
		fmt.Println("请输入正确的参数！")
	}
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		// 如果想向下穿透，加上关键字 fallthrough
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default called")
	}
}
