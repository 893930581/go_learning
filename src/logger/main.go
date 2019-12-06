package main

import (
	"go_learning/src/logger/console"
	"time"
)

var log mylogger.Logger

func main(){
	//log := mylogger.NewConsoleLogger("debug")
	//for i:=3;i>0;i--{
	//	log.Error("这是一个error级别的错误")
	//	log.Fatal("这是一个fatal级别的错误")
	//	log.Debug("这是一条debug日志||id：%d||name：%s",1772,"管理员A")
	//	time.Sleep(3)
	//}

	log := mylogger.NewFIleLogger("error","./","test.log",10*1024)
	defer log.CloseFIle()
	for i:=3;i>0;i--{
		log.Error("这是一个error级别的错误")
		log.Fatal("这是一个fatal级别的错误")
		log.Debug("这是一条debug日志||id：%d||name：%s",1772,"管理员A")
		time.Sleep(3)
	}
}