package main

/** logAgent 入口文件 */

import (
	"fmt"
	"github.com/go-ini/ini"
	"go_learning/src/logAgent/config"
	"go_learning/src/logAgent/kafka"
	"go_learning/src/logAgent/tailLog"
	"time"
)

//var wg sync.WaitGroup
var (
	cfg	= new(config.AppConfig)
)

func run(topic string) {
	// 1,读取日志写入kafka
	for {
		select {
			case line := <-tailLog.ReadLog():
				kafka.SendToKafka(topic,line.Text)
			default:
				time.Sleep(time.Second)
		}
	}
}

func main() {
	err := ini.MapTo(cfg,"./config/config.ini")
	if err != nil {
		fmt.Printf("Init Config fail, err:%v",err)
		return
	}
	fmt.Println(cfg)
	// 初始化Kafka。
	err = kafka.Init([]string{cfg.Address})
	if err != nil {
		fmt.Printf("Init kafka fail, err:%v",err)
		return
	}

	// 打开日志文件收集日志。
	err = tailLog.Init(cfg.FilePath)
	if err != nil {
		fmt.Printf("Init log tail fail, err:%v",err)
		return
	}
	//wg.Add(1)
	run(cfg.Topic)
	//wg.Wait()
}
