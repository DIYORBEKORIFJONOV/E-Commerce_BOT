package orderservice

import (
	"ecommercebot/internal/protos/orderproduct"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet() orderproduct.OrderServiceClient {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := orderproduct.NewOrderServiceClient(conn)
	return client
}
