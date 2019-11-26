package main

import (
	"fmt"
)

func main() {
	var str string = "chen艾莉"
	fmt.Println(str)

	for i, b := range []byte(str) {
		fmt.Printf("(%d %X)", i, b)
		fmt.Println(b)
		//utf-8编码
	}
	fmt.Println("遍历utf-8编码")
	

	for i, b := range str {
		fmt.Printf("(%d %X)", i, b)
		fmt.Println(b)
	}

	fmt.Println("遍历utf-8转最后unicode")
	for i, b := range []rune(str) {
		fmt.Printf("(%d %X)", i, b)
		//unicode编码
	}
	// fmt.Println("unicode")
	// bytes := []byte(str)
	// for len(bytes) > 0 {
	// 	ch, size := utf8.DecodeRune(bytes)
	// 	bytes = bytes[size:]
	// 	fmt.Printf("%c ", ch)
	// }
}
