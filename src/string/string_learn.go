package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var str string = "innoweb.com我"
	fmt.Println("字符串长度。")
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))

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
	fmt.Println(strings.Count("cheeseeee", "ee"))
	fmt.Println(strings.Count("fivevev", "vev"))
	// Count 是计算子串在字符串中出现的无重叠的次数

	fmt.Println(strings.ContainsAny("infailurex", "in"))
	fmt.Println(strings.Split("love,love",""))
	fmt.Println(strings.Split("love love love"," "))
	fmt.Println(strings.Split("a,b,c,d,e",","))
	fmt.Println(strings.Fields("sdad"))

	fmt.Println(strings.Join(strings.Split("a,b,c,d,e",","),""))

	var a string = "!!@@$$  "
	fmt.Println(strings.Trim(a,"! "))
	fmt.Println(strings.TrimLeft(a,"!@"))
	fmt.Println(strings.TrimRight(a,"$"))
	fmt.Println(strings.TrimPrefix(a,"$$"))

	fg := strconv.Itoa(21313)
	fmt.Println(fg)

}