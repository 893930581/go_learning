package main

import (
	"fmt"
	"sync"
)

var x int = 0

var wg sync.WaitGroup
var lock sync.Mutex

func Add(){
	for i := 0;i<1000000;i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main(){
	wg.Add(5)
	go Add()
	go Add()
	go Add()
	go Add()
	go Add()
	wg.Wait()
	fmt.Println(x)
}
