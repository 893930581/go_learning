package main

import "fmt"

func goR() <- chan int {
	ch := make(chan int)
	go func() {
		ch <- 5
		close(ch)
	}()
	return ch
}

func main() {
	for c := range goR() {
		fmt.Println(c)
	}
}
