package main

import "fmt"

// 定义一个接口
type IAttack interface {
	// 接口函数可以有多个，但是只能有函数类型，不可以实现
	Attack()
}

//
type HumanLowLevel struct {
	name  string
	level int
}

//
type HumanHighLevel struct {
	name  string
	level int
}

func (this *HumanLowLevel) Attack() {
	fmt.Printf("我是%s，等级为%d\n", this.name, this.level)
}

func (this *HumanHighLevel) Attack() {
	fmt.Printf("我是%s，等级为%d\n", this.name, this.level)
}

// 定义一个通用接口
func DoAttack(this IAttack) {
	this.Attack()
}

func main() {
	// 定义一个包含Attcak的接口
	// var player IAttack

	lowLevel := HumanLowLevel{
		"bill",
		1,
	}
	// lowLevel.Attack()

	// player = &lowLevel
	// player.Attack()

	highLevel := HumanHighLevel{
		"alex",
		10,
	}
	// highLevel.Attack()

	// player = &highLevel
	// player.Attack()

	DoAttack(&lowLevel)
	DoAttack(&highLevel)

	/*
		多态步骤
		1. 定义一个接口，里面设计好需要的接口，可以有多个
		2. 任何实现了这个接口的类，都可以赋值给这个接口，从而实现多态
		3. 多个类直接不需要有继承关系
		4. 如果interface中定义了多个接口，那么类必须要全部实现接口函数，才可以赋值
	*/

}
