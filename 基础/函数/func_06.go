package main

import "fmt"

func main() {
	/*
		延迟执行(函数退出时执行),堆栈： 先进后去
		作用： 资源释放、日志记录
	*/
	defer func() {
		fmt.Println("defer01")
	}()
	defer func() {
		fmt.Println("defer02")
	}()
	/*
		recover 在defer中使用,用于处理panic的
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	/*
		panic
		panic之后的代码不会执行，程序退出
	*/
	panic("error occured!")
	fmt.Println("main over")

}
