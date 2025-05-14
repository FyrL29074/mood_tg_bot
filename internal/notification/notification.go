package notification

import (
	"context"
	"mood_tg_bot/pb/apipb"
	"time"

	"google.golang.org/grpc"
)

func SendCategoriesIn12And18() {
	for {
		now := time.Now()
		next := nextTriggerTime(now)
		duration := time.Until(next)

		time.Sleep(duration)

		// sendEmotionCategories(1033135181)
		sendEmotionCategoriesGRPC(888558026)
	}
}

func sendEmotionCategoriesGRPC(chatId int) error {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return nil
	}
	defer conn.Close()

	client := apipb.NewApiServiceClient(conn)
	_, err = client.SendEmotionCategories(context.Background(), &apipb.SendEmotionCategoriesRequest{
		ChatId: int64(chatId),
	})
	return err
}

func nextTriggerTime(now time.Time) time.Time {
	year, month, day := now.Date()
	location := now.Location()

	twelve := time.Date(year, month, day, 12, 0, 0, 0, location)
	eighteen := time.Date(year, month, day, 18, 0, 0, 0, location)

	if now.Before(twelve) {
		return twelve
	}
	if now.Before(eighteen) {
		return eighteen
	}

	return time.Date(year, month, day+1, 12, 0, 0, 0, location)
}
