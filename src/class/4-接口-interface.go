package main

import "fmt"

func main() {
	/*

	 */
	// 定义三个接口
	var i, j, k interface{}
	names := []string{"alex", "bill"}
	i = names
	fmt.Println(i)
	age := 20
	j = age
	fmt.Println(j)
	str := "hello"
	k = str
	fmt.Println(k)

	// 判断interface的类型
	k, ok := k.(int)
	if !ok {
		fmt.Println("k不是int")
	} else {
		fmt.Println("k是int")
	}

	// 使用场景
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "hello"
	array[2] = true

	for _, v := range array {
		switch value := v.(type) {
		case int:
			fmt.Printf("当前类型为int，内容为:%d \n", value)
		case string:
			fmt.Printf("当前类型为string，内容为:%s \n", value)
		case bool:
			fmt.Printf("当前类型为bool，内容为:%v \n", value)
		default:
			fmt.Println("不是合理的数据")
		}
	}
}
