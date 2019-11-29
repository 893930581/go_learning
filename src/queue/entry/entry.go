package main

import (
	"fmt"
	"go_learning/src/queue"
)

func main() {
	var a []int= []int{1}
	fmt.Print(a)
	var q queue.Queue= queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Print(q.Pop())
	fmt.Print(q.Pop())
	fmt.Print(q.IsEmpty())
	fmt.Print(q.Pop())
	fmt.Print(q.IsEmpty())
	q.Push("吃饭")
	fmt.Println(q.Pop())
}
