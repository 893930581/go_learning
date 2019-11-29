package main

import "fmt"

func adder() func(int) int /** 返回值是一个参数和返回值都为int的函数 */ {
	var sum int = 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func tealAdder(base int) iAdder {
	return func(v int) (int, iAdder){
		return base+v, tealAdder(base+v)
	}
}

func main(){
	b := tealAdder(0)

	for i := 1; i < 10; i++ {
		var k int
		k, b = b(i)
		fmt.Println(k)
	}

	a := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}