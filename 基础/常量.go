package main

import "fmt"

const PI = 3.14

func main() {
	fmt.Println(PI)
	const (
		v1 = 1
		v2
		v3
		v4 = "alex"
		v5
	)
	fmt.Println(v1, v2, v3, v4, v5)
	// 枚举
	const (
		E1 int = iota
		E2
		E3
	)
	fmt.Println(E1, E2, E3)
	const (
		E4 int = iota
		E5
		E6
	)
	fmt.Println(E4, E5, E6)

	/*
		作用域： go中使用{}来定义作用域的范围
	*/

	outer := 1
	{
		inner := 2
		fmt.Println(outer, inner)
		{
			in_inner := 3
			fmt.Println(outer, inner, in_inner)
		}
	}
}
