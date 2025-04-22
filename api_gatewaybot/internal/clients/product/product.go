package productservice

import (
	productpb "ecommercebot/internal/protos/github.com/diyorbek/E-Commerce_BoT/product-service/pkg/protos/gen/product"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet() productpb.ProductServiceClient {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := productpb.NewProductServiceClient(conn)
	return client
}
