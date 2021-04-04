package main

import "fmt"

func changeInt(n int) {
	n = 100
}

func changeIntBYPopinter(p *int) {
	*p = 100
}

func changeSlice(s []int) {
	s[0] = 100
}

func main() {
	/*
		值类型   ： int bool float  pointer
		引用类型 ： slice map
		将变量赋值给一个值，如果变量发生改变对旧变量有影响就是引用类型，无影响就是值类型
	*/
	array1 := [3]string{"A", "B", "C"}
	slice1 := []string{"A", "B", "C"}
	array2 := array1
	slice2 := slice1
	array2[0] = "Z"
	slice2[0] = "Z"
	fmt.Println(array1, array2)
	fmt.Println(slice1, slice2)

	m := map[string]string{}
	m1 := m
	m1["alex"] = "武汉"
	fmt.Println(m1, m)

	// 值传递
	num := 1
	changeInt(num)
	fmt.Println(num)

	nums := []int{1, 2, 3}
	changeSlice(nums)
	fmt.Println(nums)

	changeIntBYPopinter(&num)
	fmt.Println(num)
}
