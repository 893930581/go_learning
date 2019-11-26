package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("arr= %v, len=%d, cap=%d \n", arr, len(arr), cap(arr))
}

func main() {
	var s []int
	fmt.Println("making slice by append \n")
	for i := 0; i < 5; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	fmt.Println("making slice by []")
	s1 := []int{2, 3, 4, 5, 6}
	printSlice(s1)

	fmt.Println("making slice by len and cap \n")
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying Slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice \n")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Pop from front \n")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front, "shift")
	printSlice(s2)

	fmt.Println("Pop from back \n")
	back := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(back, "已被pop")
	printSlice(s2)
}
