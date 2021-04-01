package main

import "fmt"

func main() {
	// 单个变量
	var name string = "alex"
	var age int
	age = 20
	var isBoy = true
	fmt.Println(name, age, isBoy)

	// 多个变量(不推荐的写法)
	var englistScore, mathScore, musicScore int = 95, 95, 95
	fmt.Println(englistScore, mathScore, musicScore)
	var sId, sName, sAge = 1, "bill", 30
	fmt.Println(sId, sName, sAge)
	// 多个变量(推荐的写法)
	var (
		num   = 10
		price = 10.5
	)
	fmt.Println(num, price)

	// 简短声明，只能在函数内部使用
	isGirl := true
	fmt.Println(isGirl)
}
