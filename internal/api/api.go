package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mood_tg_bot/pb/storagepb"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartBot() {
	InitKafkaWriter()
	initApi()

	SendPhoto(888558026, "Бот запущен!")
	handleResponses()
}

var client *http.Client

func initApi() {
	client = &http.Client{Timeout: 300 * time.Second}
}

func handleResponses() {
	var LastUpdateID int
	for {
		upds, err := GetUpdates(fmt.Sprint(LastUpdateID))
		if err != nil {
			panic(err)
		}

		for _, upd := range upds {

			callbackData := ""
			if upd.CallbackQuery != nil {
				callbackData = upd.CallbackQuery.Data
			}

			var chatId int
			if upd.CallbackQuery != nil {
				chatId = upd.CallbackQuery.MsgInfo.Chat.Id
			} else {
				chatId = upd.MsgInfo.Chat.Id
			}

			_, isCategory := emotionCategories[callbackData]
			_, isEmotion := emotionsCategoriesMap[callbackData]
			isMessage := upd.MsgInfo != nil
			isBackSymbol := callbackData == backSymbol

			switch {
			case isMessage || isBackSymbol || (!isCategory && !isEmotion):
				err = SendEmotionCategories(chatId)
			case isCategory:
				err = sendEmotionsMessage(chatId, callbackData)
			case isEmotion:
				category := emotionsCategoriesMap[callbackData]
				err = SendMoodToKafka(chatId, callbackData, category)
				if err != nil {
					panic(err)
				}
				err = SendMessage(chatId, moodWasSuccesfullyAddedText, nil)
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

func SendEmotionCategories(chatId int) error {
	return SendPhoto(chatId, SuggetCheckEmotionText)
}

func sendEmotionsMessage(chatId int, emotion string) error {
	switch emotion {
	case "Радость":
		return SendMessage(chatId, chooseYourEmotion, joyEmotionButtons)
	case "Грусть":
		return SendMessage(chatId, chooseYourEmotion, sadnessEmotionButtons)
	case "Злость":
		return SendMessage(chatId, chooseYourEmotion, angerEmotionButtons)
	case "Страх":
		return SendMessage(chatId, chooseYourEmotion, fearEmotionButtons)
	case "Спокойствие":
		return SendMessage(chatId, chooseYourEmotion, calmnessEmotionButtons)
	}

	return nil
}

func GetUpdates(offset string) (updates []update, err error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")
	uri := BASE_URL + getUpdatesMethod

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

func SendMessage(chatId int, message string, btns [][]inlineKeyboardButton) error {
	sendingMessage := sentMessage{
		ChatId: chatId,
		Text:   message,
		ReplyMarkup: &replyMarkup{
			btns,
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

func SendPhoto(chatId int, caption string) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}
	photoID := os.Getenv("PHOTO_ID")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}
	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")

	sendingPhoto := sentPhoto{
		ChatId:  chatId,
		Photo:   photoID,
		Caption: caption,
		ReplyMarkup: &replyMarkup{
			InlineKeyboard: emotionCategoryButtons,
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

func GetAllChatIDsFromGRPC() ([]int64, error) {
	conn, err := grpc.NewClient("storage:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := storagepb.NewStorageServiceClient(conn)
	res, err := client.GetChatIDs(context.Background(), &storagepb.Empty{})
	return res.ChatIDs, err
}
