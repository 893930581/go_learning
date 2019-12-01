package main

import "fmt"

func con(d []int) (z []int){
	return d
}

func sum(nums ...interface{}) (){
	fmt.Printf("%T,%T,%T,%T",nums...)
	fmt.Println(nums...)
}

func variableLengthArray() {
	s1 := append([]int{},10)
	s1 = append(s1,10)
	s1 = append(s1,10)
	fmt.Println(s1,"len:",len(s1),"cap:",cap(s1))
}

func deletEleFromSlice() {
	s1 := []int{1,2,3,4,5}
	s2 := append(s1[:3],[]int{1,2,3,2}...)
	printSlice(s2)
}

func printSlice (arr []int){
	fmt.Printf("arr= %v, len=%d, cap=%d \n", arr, len(arr), cap(arr))
}

type a []int

func (arr *a) pop() {
	*arr = (*arr)[0:len(*arr)-1]
}

func unshift(arr []int){
	arr = arr[1:]
	fmt.Print(arr)
}

func shift(arr []int,a int){
	arr = append([]int{a},arr...)
}

func main (){
	//a := []int{1,2,3,1,1}
	//fmt.Printf("%T",con(a))
	//fmt.Println(con(a))
	sum(1,1,"dasd",3)
	variableLengthArray()
	deletEleFromSlice()
	var b a = []int{1,2,3,1}
	d := []int{1,2,3}
	b.pop()
	printSlice(b)
	unshift(d)
	printSlice(d)

	m1 := map[interface{}]interface{} {
		"name": 1,
		1:	1,
	}
	fmt.Println(m1)

	for k,v :=range m1 {
		fmt.Println(k)
		fmt.Println(v)
	}

	if value, ok := m1["name"]; ok {
		fmt.Println(value)
	}else{
		fmt.Println("键不存在。")
	}
	delete(m1,"name")

	var t2 string = "123吴"

	fmt.Println([]byte(t2))
	fmt.Println([]rune(t2))
	for _,ch := range []rune(t2){
		fmt.Printf("%T",ch)
	}
}