package main

import "fmt"

func main() {
	A := 2
	B := A
	B = 3
	A = 4
	fmt.Println(A, B)

	C := &A
	var D *int = &A
	fmt.Printf("%T %v %p\n", C, C, D)
	*C = 20
	fmt.Println(A)

	// 用户输入
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Println("你输入的是:", name)
}
