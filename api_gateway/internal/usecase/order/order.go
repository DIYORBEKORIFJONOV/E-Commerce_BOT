package usecaseorder

import (
	models "api_gateway/internal/entity/order"
	"context"
)

type orderService interface {
	CreateOrder(ctx context.Context, req *models.CreateOrderReq) (*models.Order, error)
	GetOrders(ctx context.Context, req *models.GetAllOrdersReq) (*models.GetAllOrdersRes, error)
	OrderCompleted(ctx context.Context, req *models.UpdateOrderReq) (*models.Order, error)
	AddProduct2Cart(ctx context.Context, req *models.AddProducts2Cart) (*models.GeneralOrderResponse, error)
	GetCart(ctx context.Context, req *models.GetCartReq) (*models.Cart, error)
	UpdateCart(ctx context.Context, req *models.UpdateCartReq) (*models.Cart, error)
	DeleteCart(ctx context.Context, req *models.GetCartReq) (*models.GeneralOrderResponse, error)
	DeleteProductsFromCart(ctx context.Context, req *models.DeleteProductsfromCartReq) (*models.Cart, error)
}

type OrderUseCaseIml struct {
	order orderService
}

func NewOrderService(order orderService) *OrderUseCaseIml {
	return &OrderUseCaseIml{order: order}
}

func (o *OrderUseCaseIml) CreateOrder(ctx context.Context, req *models.CreateOrderReq) (*models.Order, error) {
	return o.order.CreateOrder(ctx, req)
}

func (o *OrderUseCaseIml) GetOrders(ctx context.Context, req *models.GetAllOrdersReq) (*models.GetAllOrdersRes, error) {
	return o.order.GetOrders(ctx, req)
}

func (o *OrderUseCaseIml) OrderCompleted(ctx context.Context, req *models.UpdateOrderReq) (*models.Order, error) {
	return o.order.OrderCompleted(ctx, req)
}

func (o *OrderUseCaseIml) AddProduct2Cart(ctx context.Context, req *models.AddProducts2Cart) (*models.GeneralOrderResponse, error) {
	return o.order.AddProduct2Cart(ctx, req)
}

func (o *OrderUseCaseIml) GetCart(ctx context.Context, req *models.GetCartReq) (*models.Cart, error) {
	return o.order.GetCart(ctx, req)
}

func (o *OrderUseCaseIml) UpdateCart(ctx context.Context, req *models.UpdateCartReq) (*models.Cart, error) {
	return o.order.UpdateCart(ctx, req)
}

func (o *OrderUseCaseIml) DeleteCart(ctx context.Context, req *models.GetCartReq) (*models.GeneralOrderResponse, error) {
	return o.order.DeleteCart(ctx, req)
}

func (o *OrderUseCaseIml) DeleteProductsFromCart(ctx context.Context, req *models.DeleteProductsfromCartReq) (*models.Cart, error) {
	return o.order.DeleteProductsFromCart(ctx, req)
}
