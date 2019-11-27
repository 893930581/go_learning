package main

import "fmt"

var (
	a = 1
	b = "string"
)

var c = 2
var d string = "string2"

// e := 2

func variableZero() {
	var a string 
	var b int 
	fmt.Printf("%d %q\n", a, b)
}

func variableType() {
	var a string = "string"
	var b int = 6
	fmt.Println(a, b)
}

func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "string"
	fmt.Println(a, b, c, d)
}

func variableTypeDeductionShort() {
	a, b, c, d := 1, 2, true, "string"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("Hello world!")
	variableZero()
	variableType()
	variableTypeDeduction()
	variableTypeDeductionShort()
}
