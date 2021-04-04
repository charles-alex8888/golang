package main

import "fmt"

func main() {
	/*
		切片是长度可变的数组
		array[start:end]
		array[start:end:cap]
	*/
	var sliceNum []int
	fmt.Printf("%T\n", sliceNum)
	// 字面量
	sliceNum = []int{1, 2, 3}
	fmt.Println(sliceNum)
	sliceNum = []int{1, 2, 3, 4}
	fmt.Println(sliceNum)
	// 数组切片赋值
	var arrNum [10]int = [10]int{1, 2, 3, 4, 5, 6}
	sliceNum = arrNum[1:10]
	fmt.Println(sliceNum)
	fmt.Printf("%#v %d %d\n", sliceNum, len(sliceNum), cap(sliceNum))

	// make函数
	sliceNum = make([]int, 3)
	fmt.Printf("%#v %d %d\n", sliceNum, len(sliceNum), cap(sliceNum))
	sliceNum = make([]int, 3, 5)
	fmt.Printf("%#v %d %d\n", sliceNum, len(sliceNum), cap(sliceNum))

	// 增, 容量超出后会扩展1-1.5倍
	sliceNum = append(sliceNum, 1, 2, 3, 4, 5, 6)
	fmt.Printf("%#v %d %d\n", sliceNum, len(sliceNum), cap(sliceNum))
	// 遍历
	for _, v := range sliceNum {
		fmt.Println(v)
	}

	// 切片操作
	fmt.Println(sliceNum[1:5])
	fmt.Println(sliceNum[6:10])
	sliceNum = make([]int, 3, 10)
	// 容量长度为新cap-start，如果没有cap,则容量为原cap-start
	nums := sliceNum[1:3:10]
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	// 切片操作会影响源

	// copy 删除切片
	nums01 := []int{1, 2, 3}
	nums02 := []int{10, 20, 30, 40}
	copy(nums02, nums01)
	fmt.Println(nums02)

	nums03 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	copy(nums03[2:], nums03[3:])
	fmt.Println(nums03)

	// 队列，先进先出
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)
	fmt.Println(queue)
	fmt.Println(queue[1:])

	// 堆栈, 先进后出
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack[:len(stack)-1])
}
