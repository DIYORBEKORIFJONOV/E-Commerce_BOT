package orderservice

import (
	"context"
	"fmt"

	models "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/order"
	clientgrpcserver "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/client_grpc_server"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/order/adjust/adjustrequest"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/order/adjust/adjustresponse"
	orderproduct "github.com/diyorbek/E-Commerce_BOT/api_gateway/pkg/protos/gen/order"
)

type OrderServiceforbot struct {
	O   clientgrpcserver.ServiceClient
	Req *adjustrequest.AdjustRequest
	Res *adjustresponse.AdjustResponse
}

func NewOrderService(req *adjustrequest.AdjustRequest, 
	res *adjustresponse.AdjustResponse, O *clientgrpcserver.ServiceClient)  *OrderServiceforbot{
	return &OrderServiceforbot{
		Req: req,
		Res: res,
		O:	*O,
	}
}


func (u *OrderServiceforbot) CreateOrder(ctx context.Context, req *models.CreateOrderReq) (*models.Order, error) {
	new_req, err := u.Req.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.OrderService().CreateOrder(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.CreateOrder(ctx, res)
	if err != nil {
		return nil, err
	}

	u.O.OrderService().DeleteCart(ctx,&orderproduct.DeleteCartReq{
		UserId: req.UserID,
	})
	return new_res, nil
}
func (u *OrderServiceforbot) GetOrders(ctx context.Context, req *models.GetAllOrdersReq) (*models.GetAllOrdersRes, error) {
	new_req, err := u.Req.GetOrders(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := u.O.OrderService().GetAllOrders(ctx, new_req)
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
	res, err := u.O.OrderService().OrderCompleted(ctx, new_req)
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
	if req.Quantity <= 0 {
		return nil,fmt.Errorf("quantity must be positive number")
	}

	cart,err := u.O.OrderService().GetCart(ctx,&orderproduct.GetCartReq{
		UserId: req.UserID,
	})
	if err != nil {
		return nil,err
	}

	for _, p  := range cart.Products {
		if p.ProductId == req.ProductID {
			u.O.OrderService().UpdateCart(ctx,&orderproduct.UpdateCartReq{
				UserId: req.UserID,
				ProductId: req.ProductID,
				Quantity: req.Quantity + p.Quantity,
			})

			return &models.GeneralOrderResponse{
				Status: true,
				Message: "Product is successfuly added  to cart",
			},nil
		}
	}

	
	res, err := u.O.OrderService().AddProductsToCart(ctx, new_req)
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

	res, err := u.O.OrderService().GetCart(ctx, new_req)
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
	if req.Quantity == 0 {
		res,err :=u.O.OrderService().DeleteCartProducts(ctx,&orderproduct.DeleteProductsfromCartReq{
			UserId: req.UserID,
			ProductId: req.ProductID,
		})
		if err != nil {
			return nil,err
		}

		return u.Res.DeleteProductsFromCart(ctx,res)
	}
	if req.Quantity < 0 {
		return nil,fmt.Errorf("quantity must be posstive number")
	}
	new_req, err := u.Req.UpdateCart(ctx, req)
	if err != nil {
		return nil, err
	}

	res, err := u.O.OrderService().UpdateCart(ctx, new_req)
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
	res, err := u.O.OrderService().DeleteCart(ctx, &orderproduct.DeleteCartReq{UserId: new_req.UserId})
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
	res, err := u.O.OrderService().DeleteCartProducts(ctx, new_req)
	if err != nil {
		return nil, err
	}
	new_res, err := u.Res.DeleteProductsFromCart(ctx, res)
	if err != nil {
		return nil, err
	}
	return new_res, nil
}
