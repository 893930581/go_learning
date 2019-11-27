package main

import "fmt"

func calcurate(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	default:
		panic(op)
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {
	//可以不加条件，将条件写在case后面。
	case score < 60:
		g = "F"
	case score > 60:
		g = "A"
	default:
		g = "明天叫你家长来"
	}
	return g
}

func main() {
	fmt.Println(calcurate(1, 2, "+"))
	fmt.Println(grade(78))
}
