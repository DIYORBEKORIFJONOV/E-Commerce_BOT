package interface17

import (
	"context"
	models "ecommercebot/internal/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *models.CreateOrderReq) (*models.Order, error)
	GetOrders(ctx context.Context, req *models.GetAllOrdersReq) (*models.GetAllOrdersRes, error)
	OrderCompleted(ctx context.Context, req *models.UpdateOrderReq) (*models.Order, error)
	AddProduct2Cart(ctx context.Context, req *models.AddProducts2Cart) (*models.GeneralOrderResponse, error)
	GetCart(ctx context.Context, req *models.GetCartReq) (*models.Cart, error)
	UpdateCart(ctx context.Context, req *models.UpdateCartReq) (*models.Cart, error)
	DeleteCart(ctx context.Context, req *models.GetCartReq) (*models.GeneralOrderResponse, error)
	DeleteProductsFromCart(ctx context.Context, req *models.DeleteProductsfromCartReq) (*models.Cart, error)
}
