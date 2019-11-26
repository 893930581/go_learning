package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	default:
		return 0, fmt.Errorf(
			"unsupport operation: %s", op)
	}
	return result, nil
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func div3(a, b int) (z, x int) {
	//不推荐
	z = a / b
	x = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d,%d)", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	if result, error := eval(3, 5, "x"); result != 0 {
		fmt.Println(result)
	} else {
		fmt.Println("Error:", error)
	}
	q, r := div2(1, 2)
	fmt.Println(q, r)
	z, _ := div3(4, 5)
	fmt.Println(z)

	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sum(1, 2, 1, 2))
}
