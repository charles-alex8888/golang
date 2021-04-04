package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	return a / b, nil
}

func divi(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	/*
			 返回值， 自定义错误类型
		     errors.New() 引入包
			 fmt.Errof() 引入包
			无错误
	*/
	fmt.Println(div(1, 3))
	if v, err := div(1, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
}
