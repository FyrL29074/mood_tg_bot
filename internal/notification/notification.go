package notification

import (
	"context"
	"mood_tg_bot/pb/apipb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendCategoriesIn12And18() {
	sendEmotionCategories()

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

	ten := time.Date(year, month, day, 10, 0, 0, 0, location)
	fifteen := time.Date(year, month, day, 15, 0, 0, 0, location)
	twenty := time.Date(year, month, day, 20, 0, 0, 0, location)

	if now.Before(ten) {
		return ten
	}
	if now.Before(fifteen) {
		return fifteen
	}
	if now.Before(twenty) {
		return twenty
	}

	return time.Date(year, month, day+1, 12, 0, 0, 0, location)
}
