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
		next := nextCategoriesTriggerTime(now)
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

func nextCategoriesTriggerTime(now time.Time) time.Time {
	location := time.FixedZone("UTC+3", 3*60*60)
	now = now.In(location)
	year, month, day := now.Date()

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

func SendStatisticsOnMonday() {
	sendStatistics()

	for {
		now := time.Now()
		next := nextStatisticsTriggerTime(now)
		duration := time.Until(next)

		time.Sleep(duration)

		sendStatistics()
	}
}

func nextStatisticsTriggerTime(now time.Time) time.Time {
	location := time.FixedZone("UTC+3", 3*60*60)
	now = now.In(location)

	daysUntilMonday := (int(time.Monday) - int(now.Weekday()) + 7) % 7
	if daysUntilMonday == 0 {
		daysUntilMonday = 7
	}

	nextMonday := now.AddDate(0, 0, daysUntilMonday)
	year, month, day := nextMonday.Date()

	return time.Date(year, month, day, 10, 0, 0, 0, location)
}

func sendStatistics() error {
	conn, err := grpc.NewClient("api:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := apipb.NewApiServiceClient(conn)
	_, err = client.SendStatistics(context.Background(), &apipb.Empty{})
	return err
}
