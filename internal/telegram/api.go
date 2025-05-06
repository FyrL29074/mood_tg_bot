package telegram

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

const (
	GET_UPDATES_METHOD              = "getUpdates"
	SEND_MESSAGE_METHOD             = "sendMessage"
	SUGGEST_CHECK_EMOTION_TEXT      = "Как ты себя сейчас чувствуешь?"
	MOOD_WAS_ADDED_SUCCESFULLY_TEXT = "Ваша эмоция была успешная сохранена!"
	GOOD_CALLBACK_DATA              = "good"
	NORMAL_CALLBACK_DATA            = "normal"
	BAD_CALLBACK_DATA               = "bad"
)

func StartBot() {
	handleResponses()
}

func handleResponses() {
	var LastUpdateID int

	for {
		updates, err := GetUpdates(fmt.Sprint(LastUpdateID))
		if err != nil {
			panic(err)
		}
		for _, update := range updates {
			if update.MsgInfo != nil {
				err = sendEmotionsMessage(update.MsgInfo.Chat.Id)
				fmt.Println("sent message: " + update.MsgInfo.Text)
			}
			if err != nil {
				panic(err)
			}

			if update.CallbackQuery != nil {
				err = AddMood(update.CallbackQuery.MsgInfo.Chat.Id, update.CallbackQuery.Data)
				if err != nil {
					panic(err)
				}
				err = SendMessage(update.CallbackQuery.MsgInfo.Chat.Id, MOOD_WAS_ADDED_SUCCESFULLY_TEXT)
				if err != nil {
					panic(err)
				}
				fmt.Println("sent message: " + update.CallbackQuery.Data)
			}
			if err != nil {
				panic(err)
			}

			LastUpdateID = update.UpdateId + 1
		}

		GetUpdates(fmt.Sprint(LastUpdateID)) // Костыль чтобы пометить последнее сообщение как отработанное

		time.Sleep(2 * time.Second)
	}
}

func sendEmotionsMessage(chatId int) error {
	btns := []inlineKeyboardButton{
		{Text: "Хорошо", CallbackData: GOOD_CALLBACK_DATA},
		{Text: "Нормально", CallbackData: NORMAL_CALLBACK_DATA},
		{Text: "Плохо", CallbackData: BAD_CALLBACK_DATA},
	}

	return SendMessageWithInlinedButtons(chatId, SUGGEST_CHECK_EMOTION_TEXT, btns)
}

func GetUpdates(offset string) (messages []userAction, err error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")
	uri := BASE_URL + GET_UPDATES_METHOD
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
	client := http.Client{Timeout: 3 * time.Second} // todo don't create client every time, create singleton once

	sendingMessage := sentMessage{
		ChatId:      chatId,
		Text:        message,
		ReplyMarkup: nil,
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")
	url := BASE_URL + SEND_MESSAGE_METHOD
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

func SendMessageWithInlinedButtons(chatId int, message string, btns []inlineKeyboardButton) error {
	client := http.Client{Timeout: 3 * time.Second}

	sendingMessage := sentMessage{
		ChatId: chatId,
		Text:   message,
		ReplyMarkup: &replyMarkup{
			[][]inlineKeyboardButton{
				btns,
			},
		},
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")

	url := BASE_URL + SEND_MESSAGE_METHOD
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
