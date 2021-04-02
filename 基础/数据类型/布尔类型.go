package main

import "fmt"

func main() {
	var zero bool
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)

	/*
		逻辑运算
	*/
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	/*
		关系运算
	*/
	fmt.Println(isGirl == isBoy)
	fmt.Println(isGirl != zero)
	fmt.Printf("%T, %t", isGirl, isGirl)
}
