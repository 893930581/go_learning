package main

import "fmt"

func swap(a, b *int) {
	fmt.Println(a)
	*a, *b = *b, *a
	//*x表示地址
}

func swap2(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	a, b := 3, 4

	// var a int = &a

	// fmt.Println(&a,*1)

	swap(&a, &b)
	//&x取地址
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
