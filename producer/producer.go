package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic", 0)
	// if err != nil {
	// 	fmt.Println("error is here: " + err.Error())

	// 	return
	// }
	// conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	// conn.WriteMessages(kafka.Message{
	// 	Value: []byte("hello"),
	// })

	// Create a new Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	})

	// Write a message to the Kafka topic
	err1 := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("key"),
		Value: []byte("value"),
	})

	if err1 != nil {
		fmt.Println("Error writing message:", err1)
	}
}
