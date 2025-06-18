package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := sendLocalPhoto(888558026, "assets/mood2.png", "тест, чтобы отправить фотку")
	if err != nil {
		panic(err)
	}
}

// Сначала нужно отправить картинку в базу tg, чтобы можно было отправлять ее по file_id, что экономнее
func sendLocalPhoto(chatId int, photoPath string, caption string) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env")
	}
	BASE_URL := os.Getenv("TELEGRAM_BASE_URL_WITH_TOKEN")
	url := BASE_URL + "sendPhoto"

	file, err := os.Open(photoPath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	_ = writer.WriteField("chat_id", fmt.Sprint(chatId))
	_ = writer.WriteField("caption", caption)

	part, err := writer.CreateFormFile("photo", photoPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Ответ Telegram:", string(body))

	return nil
}
