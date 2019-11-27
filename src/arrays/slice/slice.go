package main

import "fmt"

func changEle(arr []int) {
	//使用数组切片代替指针
	arr[0] = 1000
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr[2:6]) //半开半闭原则
	fmt.Println(arr[:6])  //默认左边是0的半开半闭原则
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	//arrSlice1 :=
	// arrSlice2 := arr[:]
	changEle(arr[2:6])
	fmt.Println("新数组1", arr)
	changEle(arr[:])
	fmt.Println("新数组2", arr)
	//slice本身没有数据，他是对底层数组的一个视图。

	// fmt.Println("Reslice")

	// fmt.Println("Extending slice")

	// s1 := arr[2:6]
	// fmt.Println(s1)
	// s1 = s1[3:5]
	// fmt.Println(s1)

	arrs := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arrs[2:5] //2,3,4 (5,6,7)
	fmt.Println(s1, "len:", len(s1), "cap", cap(s1))
	s2 := s1[2:6] //4 (5,6,7)
	fmt.Println(s2, "len:", len(s2), "cap", cap(s2))
	s3 := s2[1:4] //(5,6,7)
	fmt.Println(s3, "len:", len(s3), "cap", cap(s3))

	s4 := append(s1, 10)
	fmt.Println(s4)
	fmt.Println(s1, "\tlen:", len(s1), "cap", cap(s1))
	fmt.Println(s4, "\tlen:", len(s4), "cap", cap(s4))
	fmt.Println(s1[0:6])
	fmt.Println(s4[0:6])
}
