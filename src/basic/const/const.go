package main

import "fmt"

const a = "string"

//常量必须被赋初值，但是常量不需要。

func const1() {
	const (
		a    = 1
		b, c = 1, "txt"
	)
	const const2 = "chifan"
	fmt.Println(a, b, c)
}

func main() {
	const1()
}
