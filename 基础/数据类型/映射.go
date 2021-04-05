package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 映射(map) 是存贮一系列的无序的key/value对
	var cityMap map[string]string
	cityMap = make(map[string]string)
	cityMap["湖北"] = "武汉"
	cityMap["广东"] = "广州"
	cityMap["湖南"] = "长沙"
	cityMap["河南"] = "郑州"

	for city := range cityMap {
		fmt.Println(city, cityMap[city])
	}

	// 判断map中是否存在某个key
	city, ok := cityMap["四川"]
	if ok {
		fmt.Println(city)
	} else {
		fmt.Println("error")
	}

	//
	subjectMap := map[string]string{"english": "80", "chiness": "90", "math": "95"}
	fmt.Printf("成绩表修改前数据如下，共%d条数据\n", len(subjectMap))
	for subject := range subjectMap {
		fmt.Println(subject, subjectMap[subject])
	}
	// 增加数据
	subjectMap["music"] = "80"
	fmt.Println("成绩表新增数据后-----")
	for subject := range subjectMap {
		fmt.Println(subject, subjectMap[subject])
	}
	// 删除数据
	delete(subjectMap, "english")
	fmt.Println("成绩表删除数据后-----")
	for subject := range subjectMap {
		fmt.Println(subject, subjectMap[subject])
	}
	// 创建一个值类型为map的map
	var users map[string]map[string]string
	users = map[string]map[string]string{"alex": {"age": "20", "sex": "male"}, "lucy": {"age": "20", "sex": "female"}}
	fmt.Println(users)
	for k, v := range users {
		fmt.Println(k, v)
	}
	// 统计数字出现的次数
	rand.Seed(time.Now().Unix())
	var numCounter [30]int
	for i := 0; i < 30; i++ {
		numCounter[i] = rand.Intn(10)
	}
	fmt.Println(numCounter)
	numMap := make(map[int]int)
	for _, num := range numCounter {
		numMap[num]++
	}
	fmt.Println(numMap)

}
