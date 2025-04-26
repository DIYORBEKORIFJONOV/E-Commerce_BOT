package interface17

import (
	"context"

	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/models"
	orderproduct "github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/protos"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order interface {
	CreateOrder(ctx context.Context, req *models.Order) error
	GetOrders(ctx context.Context, req *models.GetAllOrdersReq) (*models.GetAllOrdersRes, error)
	OrderCompleted(ctx context.Context, req *models.UpdateOrderReq) error
	AddProduct2Cart(ctx context.Context, req *models.AddProducts2Cart) error
	GetCart(ctx context.Context, req *models.GetCartReq) (*mongo.Cursor, error)
	UpdateCart(ctx context.Context, req *models.UpdateCartReq) error
	DeleteCart(ctx context.Context, req *models.GetCartReq) error
	DeleteProductsFromCart(ctx context.Context, req *models.DeleteProductsfromCartReq)error
	GetOrder(ctx context.Context, id string) (*models.Order, error)
}
type OrderService interface {
	CreateOrder(ctx context.Context, req *orderproduct.CreateOrderReq) (*orderproduct.Order, error)
	GeOrders(ctx context.Context, req *orderproduct.GetAllOrdersReq) (*orderproduct.GetAllOrdersRes, error)
	OrderCompleted(ctx context.Context, req *orderproduct.UpdateOrderReq) (*orderproduct.Order, error)
	AddProduct2Cart(ctx context.Context, req *orderproduct.AddProducts2Cart) (*orderproduct.GeneralOrderResponse, error)
	GetCart(ctx context.Context, req *orderproduct.GetCartReq) (*orderproduct.Cart, error)
	UpdateCart(ctx context.Context, req *orderproduct.UpdateCartReq) (*orderproduct.Cart, error)
	DeleteCart(ctx context.Context, req *orderproduct.GetCartReq) (*orderproduct.GeneralOrderResponse, error)
	DeleteProductsFromCart(ctx context.Context, req *orderproduct.DeleteProductsfromCartReq) (*orderproduct.Cart, error)
}
