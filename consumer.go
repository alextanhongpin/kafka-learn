package main

import (
	"fmt"
	"log"

	"golang.org/x/xerrors"
	sarama "gopkg.in/Shopify/sarama.v1"
)

var (
	brokers = []string{"127.0.0.1:9092"}
	topic   = "go-transactions"
	topics  = []string{topic}
)

func main() {
	mainConsumer(0)
}

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	return conf
}

func newKafkaConsumer() sarama.Consumer {
	kafka, err := sarama.NewConsumer(brokers, newKafkaConfiguration())
	if err != nil {
		log.Fatal(xerrors.Errorf("kafka error: %w", err))
	}
	return kafka
}
func mainConsumer(partition int32) {
	kafka := newKafkaConsumer()
	defer kafka.Close()
	consumer, err := kafka.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-consumer.Errors():
			fmt.Printf("kafka error: %w\n", err)
		case msg := <-consumer.Messages():
			fmt.Println(string(msg.Value))
		}
	}

}
