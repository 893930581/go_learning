package main

import (
	"fmt"
	"time"
)

func main (){
	now := time.Now()
	fmt.Printf("%v\n",now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	timeObj := time.Unix(now.Unix(), 0) //将时间戳转为时间格式
	fmt.Printf("%T\n",timeObj)

	fmt.Println(time.Now().Add(time.Duration(time.Minute*60)).Unix())
	fmt.Println()
	//tickDemo()

	fmt.Println(time.Now().Add(time.Duration(time.Second*60)).Format("2006|01|02|03|04"))

}

func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i)//每秒都会执行的任务
	}
}