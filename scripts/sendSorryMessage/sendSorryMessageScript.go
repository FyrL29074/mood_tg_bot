package main

import "mood_tg_bot/internal/api"

func main() {
	sorryMsg := "Бот был временно недоступен, приношу извинения за предоставленные неудобства. Сейчас бот снова в строю!"

	for _, id := range getAllUserID() {
		api.SendMessage(id, sorryMsg, nil)
	}
}

func getAllUserID() []int {
	return []int{1033135181, 888558026}
}
