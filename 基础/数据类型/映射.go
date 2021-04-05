package main

import "fmt"

func main() {
	var cityMap map[string]string
	cityMap = make(map[string]string)
	cityMap["湖北"] = "武汉"
	cityMap["广东"] = "广州"
	cityMap["湖南"] = "长沙"
	cityMap["河南"] = "郑州"

	for city := range cityMap {
		fmt.Println(city, cityMap[city])
	}

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
}
