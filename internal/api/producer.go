package api

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func waitForKafka() {
	for {
		conn, err := net.Dial("tcp", "kafka:9092")
		if err == nil {
			_ = conn.Close()
			break
		}
		log.Println("Ожидание Kafka...")
		time.Sleep(3 * time.Second)
	}
}

func InitKafkaWriter() {
	waitForKafka()
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "mood-events",
		Balancer: &kafka.LeastBytes{},
	}
}

func SendMoodToKafka(chatId int, mood string, category string) error {
	msgStruct := MoodMessage{ChatId: chatId, Mood: mood, Category: category}
	value, err := json.Marshal(msgStruct)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte("chatId"),
		Value: value,
	}

	err = kafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	}
	return nil
}

type MoodMessage struct {
	ChatId   int    `json:"chat_id"`
	Mood     string `json:"mood"`
	Category string `json"category"`
}
