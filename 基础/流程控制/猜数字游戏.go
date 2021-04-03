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
	const maxGuessNum = 5
	for {
		var guessNum int = 0
		for count := 1; count <= maxGuessNum; count++ {
			fmt.Print("请输入你猜的数值：")
			fmt.Scan(&guessNum)

			if guessNum > rand_num {
				fmt.Printf("你猜大了，你还有%d次机会\n", maxGuessNum-count)
			} else if guessNum < rand_num {
				fmt.Printf("你猜小了，你还有%d次机会\n", maxGuessNum-count)
			} else {
				fmt.Printf("你猜对了了，你真聪明,用了%d次就猜对了\n", count)
				break
			}
			if count == maxGuessNum {
				fmt.Println("你真笨，正确数字是", rand_num)
			}
		}
		var choice string
		fmt.Print("重新开始游戏吗？y/n: ")
		fmt.Scan(&choice)
		if choice != "y" {
			break
		}
	}
}
