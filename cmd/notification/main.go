package main

import (
	"fmt"
	"mood_tg_bot/internal/notification"
)

func main() {
	fmt.Println("Starting notification service...")
	notification.SendCategoriesIn12And18()
}
