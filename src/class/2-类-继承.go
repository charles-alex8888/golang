package main

import "fmt"

type Person struct {
	// 成员属性
	name   string
	age    int
	gender string
}

// 类方法
func (this *Person) Eat() {
	fmt.Printf("%s in eating\n", this.name)
}

//  继承Person类
type Teacher struct {
	Person
	subject string
}

// 嵌套Person类
type Student struct {
	per    Person
	school string
	score  float64
}

func main() {
	s1 := Student{
		per: Person{"alex",
			18,
			"male",
		},
		school: "一中",
		score:  95,
	}
	fmt.Println(s1)

	t1 := Teacher{}
	t1.subject = "english"
	t1.name = "lily"
	t1.age = 18
	t1.gender = "female"
	fmt.Println(t1)
	t1.Eat()
	fmt.Println(t1.Person.name)
}
