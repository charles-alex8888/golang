package main

import "fmt"

func main() {
	// 1-定义
	name := "alex"
	info := `age: 20
	gender: mem
	address: wuhan`

	fmt.Printf("name: %s, info: %s\n", name, info)
	// 2-长度,访问
	fmt.Println(len(name))

	for i := 0; i < len(name); i++ {
		fmt.Printf("i: %d, v: %c\n", i, name[i])
	}
	// 3-拼接
	i, j := "hello", "world"
	fmt.Println("i+j=", i+j)

}
