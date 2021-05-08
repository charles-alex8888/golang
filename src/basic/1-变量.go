package main

import "fmt"

func main() {
	//  变量定义: var
	//  常量定义: const

	/*
	 01-先定义变量，再赋值
	 var var_name type
	*/
	var name string
	name = "alex"

	var age int
	age = 20

	fmt.Println("name", name)
	fmt.Printf("name is: %s, %d\n", name, age)

	/*
		02-定义时赋值
	*/
	var gender = "men"
	fmt.Println("gender", gender)

	/*
		03-定义时直接赋值，使用自动推导（最常用）
	*/
	address := "武汉"
	fmt.Printf("address is: %s\n", address)

	/*
		04- 平行赋值
	*/
	i, j := 10, 20
	fmt.Printf("i: %d,j: %d\n", i, j)

	i, j = j, i
	fmt.Printf("i: %d,j: %d\n", i, j)
}
