package main

import "fmt"

import "strconv"

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	fmt.Println(n)
	return result
}

func while() {
	index := 9
	for index > 0 {
		index--
		fmt.Println(index)
	}

}

// func forever() {
// 	for {
// 		fmt.Println("这是一个死循环")
// 	}
// }

func main() {
	while()
	// forever()
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)
}
