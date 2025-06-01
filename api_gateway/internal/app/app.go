package app

import (
	"context"
	redisCash "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/redis"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/auth"
	authusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/auth"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/until"
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

const (
	authURL = "https://notify.eskiz.uz/api/auth/login"
	smsURL  = "https://notify.eskiz.uz/api/message/sms/send"
)

type App struct {
	HTTPApp *htppapp.App
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	// Инициализация gRPC клиента
	clientGrpc, err := clientgrpcserver.NewSerevice(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	// Order Service
	addJustRes := adjustresponse.NewAddJsutResponse()
	addJustReq := adjustrequest.NewAddReqeust(clientGrpc.OrderService(), clientGrpc, addJustRes)
	serviceOrder := orderservice.NewOrderService(addJustReq, addJustRes, &clientGrpc)
	orderServiceIml := usecaseorder.NewOrderService(serviceOrder)

	minioClient, err := minio.New("hurmomarkershoppicture.duckdns.org", &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalf("Ошибка инициализации клиента MinIO: %v", err)
	}
	log.Println("✅ Клиент MinIO успешно инициализирован")

	log.Println("✅ Соединение с MinIO успешно")

	if err := createBucket(minioClient, "products"); err != nil {
		log.Fatalf("Ошибка при создании бакета: %v", err)
	}

	minClient := minao1.NewFileStorage(cfg, minioClient)

	// Product Service
	serviceProduct := product_service.NewProductReqService(&clientGrpc)
	serviceProductIml := productusecase.NewProductUsage(serviceProduct)
	phoneSetting := until.NewPhoneSetting(authURL, smsURL, cfg.EskizEmail, cfg.EskizSenderId, cfg.EskizPassword)
	redisSetting := redisCash.NewRedis(*cfg)
	serviceAuth := auth.NewAuthService(redisSetting, phoneSetting, clientGrpc)
	serviceIml := authusecase.NewAuthUseCaseIml(serviceAuth)
	// HTTP сервер
	server := htppapp.NewApp(logger, cfg.AppPort, orderServiceIml, minClient, serviceProductIml, *serviceIml)
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
		log.Printf("📦 Бакет %s не существует. Создаём...", bucket)
		return client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
	}
	log.Printf("📦 Бакет %s уже существует", bucket)
	return nil
}
