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

func (s *server) SendChatIDs(ctx context.Context, req *storagepb.Empty) (res *storagepb.SendChatIDsResponse, err error) {
	chatIDs, err := storage.GetAllUsersFromDB()
	if err != nil {
		return nil, err
	}

	return &storagepb.SendChatIDsResponse{Status: "OK", ChatIDs: chatIDs}, nil
}
