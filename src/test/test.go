package main
import "fmt"

func changArray (array []int) {
	array[0] = 100
}

func main(){
	array := []int{1,2,3,4}
	changArray(array)
	fmt.Println(array)
}