package main

import "fmt"

/*
	1. init函数没有参数和返回值
	2， 一个包中包含多个init时，调用顺序时不确定的
	3. init函数是不允许显示调用的
	4. 如果只想引用某个包的init函数，可以使用_形式来处理
	   import _ "xxx/xx/xx"
*/
func init() {
	fmt.Println("init called")
}

func main() {
	fmt.Println("init函数")
}
