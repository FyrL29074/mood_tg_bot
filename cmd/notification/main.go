package main

import (
	"fmt"
	"mood_tg_bot/internal/notification"
)

func main() {
	fmt.Println("Starting notification service...")
	go notification.SendCategoriesIn12And18()
	notification.SendStatisticsOnMonday()
}
