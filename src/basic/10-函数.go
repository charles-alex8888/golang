package main

import "fmt"

/*
1. 函数返回值在参数列表之后
2. 如果有多个函数返回值，需要使用()，多个参数之间用,分割
3. 当返回值有名字切使用了，return不用加参数
*/
func add(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

func addA(a, b int, c string) (sum int, str string, bl bool) {
	sum = a + b
	str = c
	bl = true
	return
}
func main() {
	r, c, _ := add(1, 2, "hello")
	fmt.Println(r, c)
	r1, c1, b1 := addA(4, 5, "hello")
	fmt.Println(r1, c1, b1)
}
