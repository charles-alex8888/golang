package main

import (
	"fmt"
	"strings"
)

func main() {
	// 比较
	fmt.Println(strings.Compare("abc", "ABC"))
	// 包含
	fmt.Println(strings.Contains("abc", "bc"))
	fmt.Println(strings.Contains("abc", "bd"))
	// 统计
	fmt.Println(strings.Count("abcAAcdaaavfgds", "A"))
	// 按空格分割
	fmt.Println(strings.Fields("abc\tetf\tghh\tdds"))
	// 自定义分割符
	fmt.Println(strings.Split("abcd;efght;ik", ";"))
	// 以什么开头
	fmt.Println(strings.HasPrefix("abvcd", "ab"))
	// 以什么结尾
	fmt.Println(strings.HasSuffix("abcde", "d"))
	// 在字符串中第一次出现的位置（从0开始）
	fmt.Println(strings.Index("abvcdcv", "v"))
	// 在字符串中第一次出现的位置（从0开始）
	fmt.Println(strings.LastIndex("abvcdcv", "v"))
	// 以某种符号连接
	fmt.Println(strings.Join([]string{"abc", "def", "lnm"}, ":"))
	// 重复
	fmt.Println(strings.Repeat("abc", 3))
	// 替换
	fmt.Println(strings.Replace("abcbcdefabg", "ab", "xx", 1))
	fmt.Println(strings.Replace("abcbcdefabg", "ab", "xxxx", -1))
	fmt.Println(strings.ReplaceAll("abcbcdefabg", "ab", "xxx"))
	// 大小写转换
	fmt.Println(strings.ToLower("abcAbcefg"))
	fmt.Println(strings.ToUpper("abcAbcefg"))
	fmt.Println(strings.Title("abcAbcefg"))
	// 去掉开始和结尾的
	fmt.Println(strings.Trim("abcdfsxabfgsabxxxbabaa", "a"))
	// 去前后的空格
	fmt.Println(strings.TrimSpace(" axv kks\n "))

}
