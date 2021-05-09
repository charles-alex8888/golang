package main

import "fmt"

/*
1. iota是常量计数器
2. iota从0开始，每换行+1
3. 常量组吐过不赋值，默认与上一行表达式相同
4. 如果同一行出现2个iota,那么2个iota值相同
5. 每一个iota组是独立的，新的const iota会重新清零
*/
const (
	MONDDAY = iota + 1
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	STATURDAY
	SUNDAY
)

func main() {
	// const + iota(常量计数器)
	fmt.Println(MONDDAY)
	fmt.Println(THURSDAY)
	fmt.Println(WEDNESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDAY)
	fmt.Println(STATURDAY)
	fmt.Println(SUNDAY)
}
