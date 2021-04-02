package main

import "fmt"

func main() {
	// int
	var age int = 30
	fmt.Printf("%T, %d\n", age, age)

	/*
		算术运算(+ - * / % ++ --)
	*/
	age++
	fmt.Println(age)
	age--
	fmt.Println(age)

	/*
		关系运算(== != > >= < <=)
	*/
	fmt.Println(2 == 3)
	fmt.Println(2 != 3)
	fmt.Println(2 >= 3)
	fmt.Println(2 <= 3)
	fmt.Println(2 > 3)
	fmt.Println(2 < 3)
	fmt.Println(2 == 3)

	/*
		位运算(& ｜ ^ << >> &^)
	*/

	/*
		赋值(= += -+ *= /= %= &= |= ^= <<= >>= &^=)
	*/

	// 类型转换
	var intA = 10
	var uintB uint = 20
	fmt.Println(uint(intA) + uintB)
}
