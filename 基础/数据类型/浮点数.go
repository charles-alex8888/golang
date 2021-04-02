package main

import "fmt"

func main() {
	// float32 float64
	// 字面量：十进制表示法 科学技数表示法
	// MEN -----> M*10^N
	// 1.9E-1-----> 0.19
	var height float64 = 1.68
	fmt.Printf("%T, %f\n", height, height)
	var weight float64 = 13.05e1
	fmt.Println(weight)
}
