package main

import "fmt"

func caculate(a, b int, c string) int {
	var result int
	switch {
	case c == "+":
		result = a + b
	case c == "-":
		result = a - b
	default:
		fmt.Println("计算错误")
	}
	return result
}

func highCaculate(caculate func(int, int, string) int, d int) (a int, b string) {
	a = caculate(2, 4, "+")
	b = "吃啊翻"
	return
}

func main() {
	fmt.Println(highCaculate(caculate, 1))
}
