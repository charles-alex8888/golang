package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	// rand_num := rand.Int() % 100
	rand_num := rand.Intn(100)
	fmt.Println(rand_num)
	var guess_num int
	var count int = 0
	for count < 5 {
		fmt.Printf("请输入一个100以内的整数:")
		fmt.Scan(&guess_num)
		switch {
		case guess_num > rand_num:
			fmt.Printf("你猜大了，你还有%d次机会\n", 4-count)
		case guess_num < rand_num:
			fmt.Printf("你猜小了，你还有%d次机会\n", 4-count)
		case guess_num == rand_num:
			fmt.Printf("你猜对了了，你真聪明,用了%d次就猜对了\n", 4-count)
			count = 5
		}
		count++
	}
	if count == 5 {
		fmt.Println("你真笨，正确数字是", rand_num)
	}
}
