package main

import (
	"context"
	"fmt"
	"mood_tg_bot/internal/storage"
	"mood_tg_bot/pb/storagepb"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting storage service...")
	storage.InitDb()
	go storage.StartKafkaConsumer()

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	storagepb.RegisterStorageServiceServer(s, &server{})

	fmt.Println("gRPC сервер api запущен на порту 50051...")
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

type server struct {
	storagepb.UnimplementedStorageServiceServer
}

func (s *server) GetChatIDs(ctx context.Context, req *storagepb.Empty) (res *storagepb.SendChatIDsResponse, err error) {
	chatIDs, err := storage.GetAllUsersFromDB()
	if err != nil {
		return nil, err
	}

	return &storagepb.SendChatIDsResponse{Status: "OK", ChatIDs: chatIDs}, nil
}

func (s *server) GetStatistics(ctx context.Context, req *storagepb.GetStatisticsRequest) (*storagepb.StatisticsResponse, error) {
	stat, err := storage.GetStatistics(int(req.ChatId))
	if err != nil {
		return nil, err
	}

	statPb := statPbMap(*stat)

	return &storagepb.StatisticsResponse{Categories: statPb}, nil
}

func statPbMap(stat storage.Statistics) []*storagepb.Category {
	var statPb []*storagepb.Category

	for categoryName, category := range stat.Categories {
		var emotions []*storagepb.Emotion
		for emotion, count := range category.Emotions {
			emotions = append(emotions, &storagepb.Emotion{Name: emotion, Count: int32(count)})
		}

		statPb = append(statPb, &storagepb.Category{Name: categoryName, Emotions: emotions})
	}

	return statPb
}
