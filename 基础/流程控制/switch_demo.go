package main

import "fmt"

func main() {

	var score int
	fmt.Printf("请输入成绩：")
	fmt.Scan(&score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}

	var flag string = "N"
	switch flag {
	case "Y", "y":
		fmt.Println("you get a yes")
	case "N":
		fmt.Println("you get a no")
	default:
		fmt.Println("you get a no")
	}
}
