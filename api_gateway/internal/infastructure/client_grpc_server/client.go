package clientgrpcserver

import (
	"log"

	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/config"
	orderproduct "github.com/diyorbek/E-Commerce_BOT/api_gateway/pkg/protos/gen/order"
	productpb "github.com/diyorbek/E-Commerce_BOT/api_gateway/pkg/protos/gen/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	OrderService() orderproduct.OrderServiceClient
	ProductService() productpb.ProductServiceClient
	Close() error
}

type serviceClient struct {
	connection  []*grpc.ClientConn
	orderService orderproduct.OrderServiceClient
	productService  productpb.ProductServiceClient
}

func NewSerevice(cfg *config.Config) (ServiceClient,error) {
	// log.Fatal(cfg.AppPort,cfg.OrderServicePort,cfg.ProductServicePort)
	ColOrderService,err := grpc.NewClient(cfg.OrderServicePort,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		
		return nil,err
	}

	ColProdductService,err := grpc.NewClient(cfg.ProductServicePort,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil,err
	}


	return &serviceClient{
		productService: productpb.NewProductServiceClient(ColProdductService),
		orderService: orderproduct.NewOrderServiceClient(ColOrderService),
		connection: []*grpc.ClientConn{ColOrderService,ColOrderService},
	},nil
}


func (s *serviceClient)OrderService() orderproduct.OrderServiceClient{
	return s.orderService
}


func (s *serviceClient)ProductService() productpb.ProductServiceClient {
	return s.productService
}


func (s *serviceClient) Close() error {
	var err error
	for _, conn := range s.connection {
		if cer := conn.Close(); cer != nil {
			log.Println("Error while closing gRPC connection:", cer)
			err = cer
		}
	}
	return err
}


