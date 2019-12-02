package main

import "fmt"

func main() {
	//a := make([]int,10)
	//	for i:=0; i<10; i++ {
	//	go func(i int) {
	//		for {
	//			a[i]++
	//			runtime.Gosched()
	//			//fmt.Printf("HEllO FROM GO ROUTINUE %d\n",i)
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Millisecond)
	//fmt.Println(a)

	var a int = 0x35484c55
	fmt.Println(a)
	fmt.Printf("%X",0x35484c55)
}
