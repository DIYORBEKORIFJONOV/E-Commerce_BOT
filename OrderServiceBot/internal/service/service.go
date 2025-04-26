package service

import (
	"context"

	interface17 "github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/interface"
	orderproduct "github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/protos"
)

type Service struct {
	orderproduct.UnimplementedOrderServiceServer
	A interface17.OrderService
}

func (u *Service) AddProductsToCart(ctx context.Context, req *orderproduct.AddProducts2Cart) (*orderproduct.GeneralOrderResponse, error) {
	res, err := u.A.AddProduct2Cart(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) CreateOrder(ctx context.Context, req *orderproduct.CreateOrderReq) (*orderproduct.Order, error) {
	res, err := u.A.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) DeleteCart(ctx context.Context, req *orderproduct.DeleteCartReq) (*orderproduct.GeneralOrderResponse, error) {
	res, err := u.A.DeleteCart(ctx, &orderproduct.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) DeleteCartProducts(ctx context.Context, req *orderproduct.DeleteProductsfromCartReq) (*orderproduct.Cart, error) {
	res, err := u.A.DeleteProductsFromCart(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) GetAllOrders(ctx context.Context, req *orderproduct.GetAllOrdersReq) (*orderproduct.GetAllOrdersRes, error) {
	res, err := u.A.GeOrders(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) GetCart(ctx context.Context, req *orderproduct.GetCartReq) (*orderproduct.Cart, error) {
	res, err := u.A.GetCart(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) OrderCompleted(ctx context.Context, req *orderproduct.UpdateOrderReq) (*orderproduct.Order, error) {
	res, err := u.A.OrderCompleted(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) UpdateCart(ctx context.Context, req *orderproduct.UpdateCartReq) (*orderproduct.Cart, error) {
	res, err := u.A.UpdateCart(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
