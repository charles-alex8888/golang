package main

import "fmt"

//相同类型的参数 合并
func add(a, b int) (int, int, int) {
	return a, b, a + b
}

func add1(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// 可变参，切片类型
func addN(a, b int, args ...int) int {
	sum := 0
	sum += a
	sum += b
	for _, v := range args {
		sum += v
	}
	return sum
}

//
func addCalc(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

//
func subCalc(args ...int) int {
	sum := 0
	for _, v := range args {
		sum -= v
	}
	return sum
}

//
func mulCalc(args ...int) int {
	sum := 0
	for _, v := range args {
		sum *= v
	}
	return sum
}

//
func divCalc(args ...int) int {
	sum := 0
	for _, v := range args {
		sum /= v
	}
	return sum
}

// cal
func calc(op string, args ...int) int {
	switch op {
	case "add":
		return addCalc(args...)
	case "sub":
		return subCalc(args...)
	case "mul":
		return mulCalc(args...)
	case "div":
		return divCalc(args...)
	}
	return -1
}

func main() {
	a, b, result := add(3, 5)
	fmt.Println(a, b, result)
	fmt.Println(add1(4, 6))
	result = addN(1, 2, 3, 4, 5)
	fmt.Println(result)
	result = calc("add", 9, 8, 7, 6)
	fmt.Println(result)
	result = calc("sub", 9, 8, 7, 6)
	fmt.Println(result)
}
