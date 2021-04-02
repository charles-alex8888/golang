package main

import "fmt"

func main() {
	// ""  可解释字符串
	// ``  原生字符串
	// 特殊字符 \n \f \t \r \b \v
	var name string = "al\tex"
	var desc = `我来自\t中国`
	fmt.Println(name)
	fmt.Printf("%T, %s\n", desc, desc)

	// 字符串连接
	name = "alex,"
	desc = "我来自中国"
	fmt.Println(name + desc)

	// 关系运算（== != < <= > >=)

	// 索引 0-（n-1) (n为字符串长度)
	// 字符串定义内容必须为ascii
	desc = "abcdef"
	fmt.Printf("%T %c\n", desc[0], desc[0])

	// 切片
	// 字符串定义内容必须为ascii
	fmt.Printf("%T %s\n", desc[0:2], desc[0:2])
	fmt.Println(len(desc))
}
