package api

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func InitKafkaWriter() {
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "mood-events",
		Balancer: &kafka.LeastBytes{},
	}
}

func SendMoodToKafka(chatId int, mood string) error {
	msgStruct := MoodMessage{ChatId: chatId, Mood: mood}
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
	ChatId int    `json:"chat_id"`
	Mood   string `json:"mood"`
}
