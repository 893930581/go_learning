package main

import "fmt"

func isPalindrome(str string) bool {
	strCodes := []rune(str)
	for i ,v := range strCodes{
		if i == len(strCodes)/2+1 {
			return true
		}
		if v != strCodes[len(strCodes)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var s string = "上海自来水来自海上"
	fmt.Println(isPalindrome(s))
	var a int = 7
	a = a/2
	fmt.Println(a)
}