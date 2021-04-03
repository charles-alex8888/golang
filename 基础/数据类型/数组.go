package main

import "fmt"

func main() {
	// 声明(定长)
	var arrNum [10]int
	var arrBool [5]bool
	var arrstring [3]string
	fmt.Printf("%T\n", arrNum)
	fmt.Println(arrNum)
	fmt.Println(arrBool)
	fmt.Printf("%T %q\n", arrstring, arrstring)

	// 字面量
	arrNum = [10]int{10, 20, 30}
	fmt.Println(arrNum)
	arrNum = [10]int{0: 10, 9: 90}
	fmt.Println(arrNum)
	arrNum = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arrNum)
	arrNum1 := [5]int{1, 2, 3, 4}
	fmt.Printf("%T %#v\n", arrNum1, arrNum1)

	// 比较(长度、类型不同，不能比较)
	arrNum2 := [3]int{2, 3, 6}
	arrNum3 := [3]int{2, 3, 4}
	fmt.Println(arrNum2 == arrNum3)

	// 获取数组长度
	fmt.Println(len(arrNum2))

	// 索引
	arrStr1 := [3]string{"alex", "bill", "chris"}
	fmt.Println(arrStr1[2])
	arrStr1[1] = "adm"
	fmt.Println(arrStr1)

	// 遍历
	for i := 0; i < len(arrStr1); i++ {
		fmt.Println(arrStr1[i])
	}
	for i, v := range arrStr1 {
		fmt.Println(i, v)
	}
	for _, v := range arrStr1 {
		fmt.Println(v)
	}

	// 切片
	fmt.Printf("%T\n", arrStr1[0:2])
	fmt.Printf("%T\n", arrStr1[0:2:3])

	// 多维数组
	var marrays [3][2]int
	fmt.Println(marrays)
	fmt.Println(marrays[1][1])
	marrays[1][1] = 1000
	marrays[0] = [2]int{1, 3}
	fmt.Println(marrays)
}
