package main

import (
	"context"
	"fmt"
	"mood_tg_bot/internal/storage"
	"mood_tg_bot/pb/storagepb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	storagepb.UnimplementedStorageServiceServer
}

func (s *server) AddMood(context context.Context, req *storagepb.AddMoodRequest) (*storagepb.AddMoodResponse, error) {
	err := storage.AddMoodToDb(int(req.ChatId), req.Mood)
	if err != nil {
		return nil, err
	}
	return &storagepb.AddMoodResponse{Status: "Ok"}, nil
}

func main() {
	fmt.Println("Starting storage service...")

	storage.StartKafkaConsumer()

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	storagepb.RegisterStorageServiceServer(s, &server{})

	fmt.Println("gRPC сервер запущен на порту 50051...")
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
