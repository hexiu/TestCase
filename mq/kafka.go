/*
 * @Date: 2021-09-23 14:44:09
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-09-23 18:09:52
 * @FilePath: /TestCase/mq/kafka.go
 */
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	sarama.Logger = log.New(os.Stdout, "[samara]", log.LstdFlags)

	msg := &sarama.ProducerMessage{}
	msg.Topic = "hello"
	msg.Partition = int32(-1)
	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.ByteEncoder("你好, 世界!")
	fmt.Println(config.Metadata.Full, config.Metadata.Retry.Max)
	// config.Metadata.Full = false
	client, err := sarama.NewClient([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Println("Err:", err)
		return
	}
	defer client.Close()
	producer, err := sarama.NewSyncProducerFromClient(client) // 同步
	if err != nil {
		panic(err)
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("Failed to produce message: ", err)
	}
	log.Printf("partition=%d, offset=%d\n", partition, offset)

}
