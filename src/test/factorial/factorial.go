package main

import "fmt"

func getFactorial(num int) int {
return 1

}

func main() {
	var a rune = '爱'
	fmt.Printf("%T",a)
	var a1 *int
	fmt.Println(a1)
	var a2 = new(int)
	fmt.Println(a2)
	b := make([]int,5,10)
	c := make(map[string]int,10)

	d := []int{1,2,3,5}
	e := d
	e[0] = 1000
	fmt.Println(d)
	fmt.Println(b)
	fmt.Print(c)
}