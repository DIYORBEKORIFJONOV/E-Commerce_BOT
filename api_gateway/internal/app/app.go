package app

import (
	"context"
	"log"
	"log/slog"

	htppapp "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/app/htpp"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/config"
	clientgrpcserver "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/client_grpc_server"
	minao1 "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/minao"
	orderservice "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/order"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/order/adjust/adjustrequest"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/order/adjust/adjustresponse"
	product_service "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/product"
	usecaseorder "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/order"
	productusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/product"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type App struct {
	HTTPApp *htppapp.App
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	clientGrpc, err := clientgrpcserver.NewSerevice(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	addJustRes := adjustresponse.NewAddJsutResponse()
	addJustReq := adjustrequest.NewAddReqeust(clientGrpc.OrderService(), clientGrpc, addJustRes)

	serviceOrder := orderservice.NewOrderService(addJustReq, addJustRes, &clientGrpc)
	orderServiceIml := usecaseorder.NewOrderService(serviceOrder)
	var minio_client *minio.Client
	minio_client, err = minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("admin", "secretpass", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := createBucket(minio_client, "products"); err != nil {
		log.Fatal(err)
	}
	minClient := minao1.NewFileStorage(cfg,minio_client)

	serviceProduct := product_service.NewProductReqService(&clientGrpc)
	serviceProductIml := productusecase.NewProductUsage(serviceProduct)

	server := htppapp.NewApp(logger, cfg.AppPort, orderServiceIml, minClient, serviceProductIml)
	return &App{
		HTTPApp: server,
	}
}

func createBucket(client *minio.Client, bucket string) error {
	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		return err
	}

	if !exists {
		return client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
	}
	return nil
}