package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string = "innoweb.com我"
	fmt.Println(len(str))

	fmt.Println("字符串切割")
	str2 := str[0:7]
	fmt.Println(str2)

	fmt.Println("字符串比较")
	fmt.Println(strings.Compare(str,str2))
	// Compare 函数，用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。
	// 不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观。
	fmt.Println(strings.EqualFold("GO", "go"))
	// EqualFold 函数，计算 s 与 t 忽略字母大小写后是否相等。

	fmt.Println("是否包含字串")
	fmt.Println(strings.Contains(str,"inno"))
	fmt.Println(strings.Contains(str,"l n"))
	fmt.Println(strings.Count("cheese", "ee"))
	fmt.Println(strings.Count("fivevev", "vev"))
	// Count 是计算子串在字符串中出现的无重叠的次数
}