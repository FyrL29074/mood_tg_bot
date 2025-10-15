package main

import (
	"context"
	"fmt"
	"mood_tg_bot/internal/api"
	"mood_tg_bot/pb/apipb"
	"net"

	"google.golang.org/grpc"
)

func main() {
	defer func() {
		api.SendMessage(888558026, "Backend пал...", nil)
	}()

	fmt.Println("Starting api service...")
	go api.StartBot()

	listen, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	apipb.RegisterApiServiceServer(s, &server{})

	fmt.Println("gRPC сервер api запущен на порту 50052...")
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

type server struct {
	apipb.UnimplementedApiServiceServer
}

func (s *server) SendEmotionCategories(ctx context.Context, req *apipb.Empty) (*apipb.Response, error) {
	chatIDs, err := api.GetAllChatIDsFromGRPC()
	if err != nil {
		return nil, err
	}

	for _, chatID := range chatIDs {
		err := api.SendEmotionCategories(int(chatID))
		if err != nil {
			return nil, err
		}
	}

	return &apipb.Response{Status: "Ok"}, nil
}

func (s *server) SendStatistics(ctx context.Context, req *apipb.Empty) (*apipb.Response, error) {
	chatIDs, err := api.GetAllChatIDsFromGRPC()
	if err != nil {
		return nil, err
	}

	for _, chatID := range chatIDs {
		stat, err := api.GetStatistics(int(chatID))
		if err != nil {
			return nil, err
		}

		api.SendMessage(int(chatID), api.FormatStatistics(stat), nil)
	}

	return &apipb.Response{Status: "Ok"}, nil
}
