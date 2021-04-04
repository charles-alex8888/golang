package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func tower(a, b, c string, num int) {
	if num == 1 {
		fmt.Println(a, "-->", c)
		return
	}
	tower(a, c, b, num-1)
	fmt.Println(a, "-->", c)
	tower(b, a, c, num-1)
}

func main() {
	// 阶乘
	result := fact(5)
	fmt.Println(result)
	//汉罗塔
	fmt.Println("一个盘子")
	tower("A", "B", "C", 1)
	fmt.Println("多个盘子")
	tower("A", "B", "C", 5)
}
