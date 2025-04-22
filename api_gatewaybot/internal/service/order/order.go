package order

import (
	"context"
	models "ecommercebot/internal/model"
	"ecommercebot/internal/protos/orderproduct"
	"ecommercebot/internal/service/order/adjust/adjustrequest"
	"ecommercebot/internal/service/order/adjust/adjustresponse"
)

type OrderServiceforbot struct {
	O   orderproduct.OrderServiceClient
	Req *adjustrequest.AdjustRequest
	Res *adjustresponse.AdjustResponse
}

func (u *OrderServiceforbot) CreateOrder(ctx context.Context, req *models.CreateOrderReq) (*models.Order, error) {
	new_req, err := u.Req.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.CreateOrder(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.CreateOrder(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
func (u *OrderServiceforbot) GetOrders(ctx context.Context, req *models.GetAllOrdersReq) (*models.GetAllOrdersRes, error) {
	new_req, err := u.Req.GetOrders(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.GetAllOrders(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.GetOrders(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
func (u *OrderServiceforbot) OrderCompleted(ctx context.Context, req *models.UpdateOrderReq) (*models.Order, error) {
	new_req, err := u.Req.OrderCompleted(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.OrderCompleted(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.OrderCompleted(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
func (u *OrderServiceforbot) AddProduct2Cart(ctx context.Context, req *models.AddProducts2Cart) (*models.GeneralOrderResponse, error) {
	new_req, err := u.Req.AddProduct2Cart(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.AddProductsToCart(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.AddProduct2Cart(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
func (u *OrderServiceforbot) GetCart(ctx context.Context, req *models.GetCartReq) (*models.Cart, error) {
	new_req, err := u.Req.GetCart(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.GetCart(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.GetCart(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
func (u *OrderServiceforbot) UpdateCart(ctx context.Context, req *models.UpdateCartReq) (*models.Cart, error) {
	new_req, err := u.Req.UpdateCart(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.UpdateCart(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.UpdateCart(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
func (u *OrderServiceforbot) DeleteCart(ctx context.Context, req *models.GetCartReq) (*models.GeneralOrderResponse, error) {
	new_req, err := u.Req.DeleteCart(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.DeleteCart(ctx, &orderproduct.DeleteCartReq{UserId: new_req.UserId})
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.DeleteCart(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
func (u *OrderServiceforbot) DeleteProductsFromCart(ctx context.Context, req *models.DeleteProductsfromCartReq) (*models.Cart, error) {
	new_req, err := u.Req.DeleteProductsFromCart(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.DeleteCartProducts(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.DeleteProductsFromCart(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
