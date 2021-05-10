package main

import "fmt"

type Person struct {
	// 成员属性
	name   string
	age    int
	gender string
}

// 在类外面绑定方法
// func (p Person) Eat() {
// 	// 类的方法可以使用自己的成员变量
// 	fmt.Printf("%s is eating\n", p.name)
// }
func (this *Person) Eat() {
	// 类的方法可以使用自己的成员变量
	// 使用指针时 修改成员变量会修改源数据
	fmt.Printf("%s is eating\n", this.name)
}

func main() {
	// go中使用结构体实现类
	p1 := Person{
		"alex",
		20,
		"男生",
	}
	fmt.Println("p1:", p1)
	p1.Eat()
}
