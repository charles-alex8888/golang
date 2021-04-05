package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var arrA [10]int
	for i := 0; i < 10; i++ {
		arrA[i] = rand.Intn(100)
	}
	fmt.Println(arrA)
	for j := 0; j < len(arrA)-1; j++ {
		for j := 0; j < len(arrA)-1-j; j++ {
			if arrA[j] > arrA[j+1] {
				arrA[j], arrA[j+1] = arrA[j+1], arrA[j]
			}
		}
		fmt.Println(arrA)
	}
}
