package main

import (
	"context"
	"fmt"
	"mood_tg_bot/internal/api"
	"mood_tg_bot/pb/apipb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	apipb.UnimplementedApiServiceServer
}

func (s *server) SendEmotionCategories(ctx context.Context, req *apipb.SendEmotionCategoriesRequest) (*apipb.SendEmotionCategoriesResponse, error) {
	err := api.SendEmotionCategories(int(req.ChatId))
	if err != nil {
		return nil, err
	}

	return &apipb.SendEmotionCategoriesResponse{Status: "Ok"}, nil
}

func main() {
	go api.StartBot()

	listen, err := net.Listen("tcp", ":50052")
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
