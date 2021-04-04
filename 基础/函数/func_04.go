package main

import "fmt"

func print(callback func(...string), args ...string) {
	fmt.Println("输出为:")
	callback(args...)
}

func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}
}

func addBase(base int) func(int) int {
	return func(n int) int {
		return base + n
	}
}

func main() {
	print(list, "a", "b", "c")

	// 匿名函数
	sayHi := func(name string) {
		fmt.Println("hello", name)
	}
	sayHi("alex")
	// 直接调用匿名函数
	func(name string) {
		fmt.Println("Hi", name)
	}("bill")

	// 闭包
	fmt.Println(addBase(8)(10))
}
