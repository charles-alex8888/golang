package main

import "fmt"

// 无参函数
func sayHello() {
	fmt.Println("hello, world!")
}

// 有参函数(形参)
func sayHi(name string) {
	fmt.Println("你好:", name)
}

// 有参数有返回值
func add(a int, b int) int {
	return a + b
}

func main() {
	// 调用函数
	sayHello()

	// 有参函数(实参）
	sayHi("alex")

	//
	result := add(3, 5)
	fmt.Println(result)

}
