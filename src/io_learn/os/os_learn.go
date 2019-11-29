package main

import "fmt"

func main(){
	a := []int{1,2,3,4}
	b := make([]int,5)
	n := copy(b,a)
	fmt.Println(n)
	fmt.Println(a)
	fmt.Println(b)
}