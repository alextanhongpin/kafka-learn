package main

import (
	"encoding/json"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/xerrors"
	sarama "gopkg.in/Shopify/sarama.v1"
)

var (
	brokers = []string{"127.0.0.1:9092"}
	topic   = "go-transactions"
	topics  = []string{topic}
)

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	return conf
}
func newKafkaSyncProducer() sarama.SyncProducer {
	kafka, err := sarama.NewSyncProducer(brokers, newKafkaConfiguration())
	if err != nil {
		log.Fatal(xerrors.Errorf("kafka error: %w", err))
	}
	return kafka
}

type Event struct {
	AccID string
	Type  string
}
type CreateEvent struct {
	Event
	AccName string
}

func NewCreateAccountEvent(name string) CreateEvent {
	return CreateEvent{
		Event: Event{
			Type:  "CreateEvent",
			AccID: uuid.Must(uuid.NewV4()).String(),
		},
		AccName: name,
	}
}

func main() {
	kafka := newKafkaSyncProducer()
	event := NewCreateAccountEvent("hello world")
	err := sendMsg(kafka, event)
	if err != nil {
		log.Fatal(err)
	}
}

func sendMsg(kafka sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msgLog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(json)),
	}
	partition, offset, err := kafka.SendMessage(msgLog)
	if err != nil {
		return err
	}
	fmt.Printf("Message: %+v\n", event)
	fmt.Printf("Message is stored in partition %d, offset %d\n", partition, offset)
	return nil
}
