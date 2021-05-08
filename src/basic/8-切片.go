package main

import "fmt"

func main() {
	address := [7]string{"北京", "上海", "深圳", "广州", "武汉", "重庆", "南京"}

	addressA := [3]string{}
	addressA[0] = address[0]
	addressA[1] = address[2]
	addressA[2] = address[2]
	fmt.Println(addressA)
	// 修改切片会使源同步修改
	addressB := address[0:3]
	fmt.Println(addressB)

	// 1. 从0开始截取，冒号左边的数字可以省略
	addressC := address[:5]
	fmt.Println(addressC)

	// 2. 如果截取到数组末尾，冒号右边的数字可以省略
	addressD := address[3:]
	fmt.Println(addressD)

	// 3. 如果截取整个数组
	addressE := address[:]
	fmt.Println(addressE)

	// 4. 创建切片的时候，指定切片的容量,这样可以提高运行效率
	str := make([]string, 10)
	fmt.Printf("%d,%d\n", len(str), cap(str))

	//5. 深拷贝： 使用copy()函数可以是切片独立于原始数据（即修改切片不会影响到原始数组的数组）
	addressCopy := make([]string, len(address))
	copy(addressCopy, address[:])
	addressCopy[0] = "香港"
	fmt.Println(addressCopy)
	fmt.Println("原始数组：", address)
}
