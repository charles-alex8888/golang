package main

import "fmt"

func main() {
	/*
		go语言在使用指针时会使用内部的垃圾回收机制(gc: garbage collector)，开发人员不需要手动释放内存
		go语言可以返回栈上的指针，程序会在编译的时候就确定变量的分配位置
	*/
	name := "alex"
	ptr := &name
	fmt.Println("name", *ptr)
	fmt.Println("name ptr", ptr)

	// 02-使用new关键字定义
	namePtr := new(string)
	*namePtr = "bill"

	fmt.Println("name", *namePtr)
	fmt.Println("namePtr", namePtr)

	// 空指针，nil
	if namePtr == nil {
		fmt.Println("namePtr是空")
	} else {
		fmt.Println("namePtr 非空")
	}

}
