package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func StartBot() {
	SendEmotionCategories(888558026)
	InitKafkaWriter()
	handleResponses()
}

func handleResponses() {
	var LastUpdateID int
	for {
		upds, err := GetUpdates(fmt.Sprint(LastUpdateID))
		if err != nil {
			panic(err)
		}

		for _, upd := range upds {
			msgText := upd.MsgInfo.Text
			_, isCategory := emotionCategories[msgText]
			_, isEmotion := emotions[msgText]

			switch {
			case checkIfMessage(upd) && !isCategory && !isEmotion:
				err = SendPhoto(upd.MsgInfo.Chat.Id)
			case checkIfMessage(upd) && isCategory:
				err = sendEmotionsMessage(upd.MsgInfo.Chat.Id, msgText)
			case checkIfMessage(upd) && isEmotion:
				err = SendMoodToKafka(upd.MsgInfo.Chat.Id, msgText)
				if err != nil {
					panic(err)
				}
				err = SendMessage(upd.MsgInfo.Chat.Id, moodWasSuccesfullyAddedText)
			}
			if err != nil {
				panic(err)
			}

			LastUpdateID = upd.Id + 1
		}
		GetUpdates(fmt.Sprint(LastUpdateID)) // Костыль чтобы пометить последнее сообщение как отработанное

		time.Sleep(2 * time.Second)
	}
}

func checkIfMessage(upd update) bool {
	return upd.MsgInfo != nil
}

func SendEmotionCategories(chatId int) error {
	return sendMessageWithReplyButtons(chatId, suggetCheckEmotionText, emotionCategoryButtons)
}

func sendEmotionsMessage(chatId int, emotion string) error {
	switch emotion {
	case "Радость":
		return sendMessageWithReplyButtons(chatId, chooseYourEmotion, joyEmotionButtons)
	case "Грусть":
		return sendMessageWithReplyButtons(chatId, chooseYourEmotion, sadnessEmotionButtons)
	case "Злость":
		return sendMessageWithReplyButtons(chatId, chooseYourEmotion, angerEmotionButtons)
	case "Страх":
		return sendMessageWithReplyButtons(chatId, chooseYourEmotion, fearEmotionButtons)
	case "Спокойствие":
		return sendMessageWithReplyButtons(chatId, chooseYourEmotion, calmnessEmotionButtons)
	}

	return nil
}

func GetUpdates(offset string) (updates []update, err error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")
	uri := BASE_URL + getUpdatesMethod
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("offset", offset)
	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var updatesResponse getUpdatesResponse
	err = json.Unmarshal(body, &updatesResponse)
	if err != nil {
		return nil, err
	}

	return updatesResponse.UserActions, nil
}

func SendMessage(chatId int, message string) error {
	client := http.Client{Timeout: 3 * time.Second} // TODO: don't create client every time, create singleton once

	sendingMessage := sentMessage{
		ChatId:      chatId,
		Text:        message,
		ReplyMarkup: nil,
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")
	url := BASE_URL + sendMessageMethod
	b, err := json.Marshal(sendingMessage)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func sendMessageWithReplyButtons(chatId int, message string, btns []replyKeyboardButton) error {
	client := http.Client{Timeout: 3 * time.Second}

	sendingMessage := sentMessage{
		ChatId: chatId,
		Text:   message,
		ReplyMarkup: &replyMarkup{
			ReplyKeyboardButton: [][]replyKeyboardButton{btns},
			ResizeKeyboard:      true,
			OneTimeKeyboard:     true,
		},
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")

	url := BASE_URL + sendMessageMethod
	b, err := json.Marshal(sendingMessage)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func SendPhoto(chatId int) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}
	photoID := os.Getenv("PHOTO_ID")

	caption := suggetCheckEmotionText

	client := http.Client{Timeout: 3 * time.Second}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}
	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")

	sendingPhoto := sentPhoto{
		ChatId:  chatId,
		Photo:   photoID,
		Caption: caption,
		ReplyMarkup: &replyMarkup{
			ReplyKeyboardButton: [][]replyKeyboardButton{emotionCategoryButtons},
			ResizeKeyboard:      true,
			OneTimeKeyboard:     true,
		},
	}

	url := BASE_URL + sendPhotoMethod
	b, err := json.Marshal(sendingPhoto)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// func sendMessageWithInlinedButtons(chatId int, message string, btns []inlineKeyboardButton) error {
// 	client := http.Client{Timeout: 3 * time.Second}

// 	sendingMessage := sentMessage{
// 		ChatId: chatId,
// 		Text:   message,
// 		ReplyMarkup: &replyMarkup{
// 			[][]inlineKeyboardButton{
// 				btns,
// 			},
// 		},
// 	}

// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Ошибка загрузки .env")
// 	}

// 	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")

// 	url := BASE_URL + sendMessageMethod
// 	b, err := json.Marshal(sendingMessage)
// 	if err != nil {
// 		panic(err)
// 	}

// 	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Accept", "application/json")

// 	_, err = client.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
