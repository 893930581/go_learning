package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("MARA")
		time.Sleep(time.Millisecond * 500)
		select {
		case <- ctx.Done():
			break LOOP
		default:
		}
	}
}

func f1(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
LOOP:
	for {
		fmt.Println("WULA")
		time.Sleep(time.Millisecond * 500)
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
							//返回取消的函数
							//context.WithDeadline()
							//context.WithTimeout()
							//取消绝对时间和相对时间
							//time.After()
	wg.Add(2)
	go f1(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
