package main

import "fmt"

type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {

	s1 := Student{
		name:   "alex",
		age:    20,
		gender: "男",
		score:  90,
	}
	fmt.Println("s1:", s1.name, s1.age, s1.gender, s1.score)

	ptr1 := &s1
	fmt.Println("ptr1", (*ptr1).name, ptr1.age, ptr1.gender, ptr1.score)

	s2 := Student{
		name: "bill",
		age:  20,
		// gender: "男",
		// score:  95,
	}
	fmt.Println("s2:", s2.name, s2.age, s2.gender, s2.score)

	s3 := Student{
		"condy",
		20,
		"女",
		85,
	}
	fmt.Println("s3:", s3.name, s3.age, s3.gender, s3.score)

}
