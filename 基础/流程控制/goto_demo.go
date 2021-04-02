package main

import "fmt"

func main() {
	sum := 0
	i := 1
FORSTART:
	if i > 100 {
		goto FOREND
	}
	sum += i
	i++
	goto FORSTART
FOREND:
	fmt.Println(sum)
BREAKEND:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 9 {
				break BREAKEND
			}
			fmt.Println(i, j)
		}
	}
}
