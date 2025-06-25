package storage

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/segmentio/kafka-go"
)

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

func StartKafkaConsumer() {
	waitForKafka()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "mood-events",
		GroupID: "storage-group",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Kafka error: %v", err)
			continue
		}

		var moodMsg Message
		err = json.Unmarshal(m.Value, &moodMsg)
		if err != nil {
			log.Printf("Kafka error: %v", err)
			continue
		}

		if moodMsg.Text == "/start" {
			addUser(moodMsg.ChatId)
			log.Printf("Пользователь добавлен: %d", moodMsg.ChatId)
			return
		}

		err = addMoodToDb(moodMsg.ChatId, moodMsg.Mood, moodMsg.Category)
		if err != nil {
			log.Printf("Ошибка записи в БД: %v", err)
		} else {
			log.Printf("Добавлено настроение для %d: %s", moodMsg.ChatId, moodMsg.Mood)
		}
	}
}

type Message struct {
	ChatId   int    `json:"chat_id"`
	Text     string `json:"text,omitempty"`
	Mood     string `json:"mood,omitempty"`
	Category string `json:"category,omitempty"`
}
