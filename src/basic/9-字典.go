package main

import "fmt"

func main() {
	// 定义一个字典
	// var students map[int]string
	// students = make(map[int]string, 10)

	// 1. 定义时直接赋值
	students := make(map[int]string, 10)

	// 2. 赋值
	students[0] = "alex"
	students[1] = "bill"

	// 3. 遍历
	for i, v := range students {
		fmt.Printf("student id is %d, name is %s\n", i, v)
	}

	/*
	 4. map中访问不存在key不会报错，会返回零值 （数字:0  bool: false string: 空）
	 判断某个键是否存在
	*/

	value, ok := students[2]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("key不存在")
	}

	// 5. 删除key,删除不存在的key时不会报错
	delete(students, 1)
	fmt.Println(students)

}
