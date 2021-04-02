package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum", sum)

	i, sum := 0, 0

	for i <= 100 {
		sum += i
		i++
	}

	i = 0
	fmt.Println("sum", sum)

	// 死循环
	/* 	for {
		fmt.Println("A")
		i++
	} */

	// 字符串、数组、切片、映射、管道
	desc := "i love my mother"
	for i, v := range desc {
		fmt.Println(i, v)
	}
}
