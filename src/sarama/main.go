package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/Shopify/sarama"
)

func main() {
	/** 配置kafka*/
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 对应kafka的应答机制的三个状态码，这里我们等待所有的follower都同步完成之后producer再发送下一条。
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 以轮询的方式随机写入partion
	config.Producer.Return.Successes = true
	// 成功交付的信息再channel中返回

	/** 发送消息*/
	msg := &sarama.ProducerMessage{}
	// 构造消息结构体
	msg.Topic = "web_log"
	// 话题
	msg.Value = sarama.StringEncoder("this is a string !")
	// 消息字符串格式化
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	// 这里接受一个ip地址的切片，相当于是一个集群
	if err != nil {
		fmt.Println("Producer closed, err:", err)
	}
	defer client.Close()
	// 写入结束关闭连接
	pid, offset, err := client.SendMessage(msg)
	// 返回进程号和写入文件的偏移位置
	if err != nil {
		fmt.Println("send msg failed, err:",err)
	}
	fmt.Printf("pid: %v	offset: %v\n", pid, offset)
}