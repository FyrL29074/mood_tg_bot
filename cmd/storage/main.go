package main

import (
	"fmt"
	"mood_tg_bot/internal/storage"
)

func main() {
	fmt.Println("Starting storage service...")

	storage.StartKafkaConsumer()
}
