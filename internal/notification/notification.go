package notification

import (
	"context"
	"mood_tg_bot/pb/apipb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendCategoriesIn12And18() {
	for {
		now := time.Now()
		next := nextTriggerTime(now)
		duration := time.Until(next)

		time.Sleep(duration)

		sendEmotionCategories()
	}
}

func sendEmotionCategories() error {
	conn, err := grpc.NewClient("api:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := apipb.NewApiServiceClient(conn)
	_, err = client.SendEmotionCategories(context.Background(), &apipb.Empty{})
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
