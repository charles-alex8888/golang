package main

import "fmt"

func main() {
	// 切片： slice. 底层也是数据,可以动态改变长度
	// 1-定义
	address := []string{"北京", "上海", "广州", "深圳"}
	for i, v := range address {
		fmt.Println(i, v)
	}
	fmt.Printf("address的长度为%d, 容量为%d\n", len(address), cap(address))

	// 2-追加数据
	address = append(address, "武汉")
	fmt.Println(address)

	// 3-切片有长度和容量
	fmt.Printf("address的长度为%d, 容量为%d\n", len(address), cap(address))
	num := []int{}
	for i := 0; i < 50; i++ {
		num = append(num, i)
		fmt.Printf("长度为%d，容量为%d\n", len(num), cap(num))
	}

}
