package main

import (
	"fmt"
	"mood_tg_bot/internal/api"
)

func main() {
	updates, err := api.GetUpdates("0")
	if err != nil {
		panic(err)
	}

	idForClear := updates[len(updates)-1].Id + 1
	api.GetUpdates(fmt.Sprint(idForClear))
}
