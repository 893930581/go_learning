package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func Init(kafkaServer []string) (err error) {
	/** 配置kafka*/
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 对应kafka的应答机制的三个状态码，这里我们等待所有的follower都同步完成之后producer再发送下一条。
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 以轮询的方式随机写入partion
	config.Producer.Return.Successes = true
	// 成功交付的信息再channel中返回

	client, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	// 这里接受一个ip地址的切片，相当于是一个集群
	if err != nil {
		fmt.Println("Producer closed, err:", err)
		return
	}
	return
}

func SendToKafka(topic, data string) {
	//构造消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:",err)
	}
	fmt.Printf("pid: %v	offset: %v\n", pid, offset)
}
