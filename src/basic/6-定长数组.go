package main

import "fmt"

func main() {
	/*
		1-定义
		nums := [10]int{1, 2, 3, 4} (常用方式)
		var arrA = [10]int{1, 2, 3, 4, 5}
		var arrB [10]int = [10]int{1, 2, 3, 4, 5, 6}
	*/
	nums := [10]int{1, 2, 3, 4}
	fmt.Println(nums)

	// 2-遍历
	for i, v := range nums {
		fmt.Println(i, v)
	}
	for _, v := range nums {
		fmt.Println(v)
	}

}
