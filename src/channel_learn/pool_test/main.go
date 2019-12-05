package main

import (
	"fmt"
	"math/rand"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job *job
	result int64
}

var jobChan chan *job= make(chan *job,100)
var resultChan chan *result = make(chan *result,100)

func creater(jobChan chan <- *job) {
	for i:=0;i<10;i++ {
		randX := rand.Int63()
		j := &job{
			randX,
		}
		jobChan <- j
	}
}

func worker(id int,jobChan <- chan *job,resultChan chan <- *result) {
	for j := range jobChan{
		fmt.Printf("worker:%d start solve:%d\n",id,j.value)
		qNum := j.value
		sum := int64(0)
		for qNum > 0 {
			sum += qNum % 10
			qNum = qNum / 10
		}
		newResult := &result{
			job:    j,
			result: sum,
		}
		fmt.Printf("worker:%d end solve:%d\n", id,sum)
		resultChan <- newResult
	}
}

func main(){
	for i := 1;i <= 4 ;i++  {
		go worker(i,jobChan,resultChan)
	}
	go creater(jobChan)
	time.Sleep(time.Second*3)
}