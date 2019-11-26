package main

import "fmt"

//在go里面，我们一般不直接使用数组，因为想要不产生值传递我们还需要使用指针。

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func changeArray(arr [5]int) {
	arr[0] = 100
	//值传递，传参前会拷贝整个数组。
}

func realChangeArray(arr *[5]int) {
	fmt.Println((*arr)[0])
	fmt.Println("打印取地址符", arr[0])
	arr[0] = 100
	//使用指针来实现引用传递。
}

func main() {
	var array1 [5]int
	array2 := [3]int{1, 2, 3}
	array3 := [...]int{1, 2, 3, 4, 5}
	var grid [4][5]int
	// array3[6] = 10
	fmt.Println(len(array3), cap(array3))
	// printArray(array1)
	// // printArray(array2)
	// printArray(array3)

	// changeArray(array3)

	realChangeArray(&array3)

	fmt.Println(array1, array2, array3)
	fmt.Println(grid)

	for key := range array1 {
		fmt.Println(key)
	}
	for key, value := range array2 {
		fmt.Println(key, value)
	}
	for _, value := range array3 {
		fmt.Println(value)
	}
}
