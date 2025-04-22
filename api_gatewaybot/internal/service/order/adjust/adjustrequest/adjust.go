package adjustrequest

import (
	"context"
	models "ecommercebot/internal/model"
	productpb "ecommercebot/internal/protos/github.com/diyorbek/E-Commerce_BoT/product-service/pkg/protos/gen/product"
	"ecommercebot/internal/protos/orderproduct"
)

type AdjustRequest struct {
	O orderproduct.OrderServiceClient
	P productpb.ProductServiceClient
}

func (u *AdjustRequest) CreateOrder(ctx context.Context, req *models.CreateOrderReq) (*orderproduct.CreateOrderReq, error) {
	var cartproducts []*orderproduct.ProductOrder
	totalprice := 0
	for _, i := range req.CartID.Products {
		price, err := u.ProductClient(ctx, i.ProductID)
		if err != nil {
			return nil, err
		}
		var product = orderproduct.ProductOrder{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
			Price:     price,
		}
		totalprice += int(i.Quantity) * int(price)
		cartproducts = append(cartproducts, &product)
	}
	var Cart = orderproduct.Cart{
		Products:   cartproducts,
		TotalPrice: float32(totalprice),
		UserId:     req.UserID,
	}
	var location = orderproduct.Location{
		Longtitude: req.Coordination.Longitude,
		Latitude:   req.Coordination.Latitude,
	}
	var newreq = orderproduct.CreateOrderReq{
		CartId:        &Cart,
		Cordination:   &location,
		UserId:        req.UserID,
		Comment:       req.Comment,
		ContactNumber: req.ContactNumber,
	}
	return &newreq, nil
}
func (u *AdjustRequest) GetOrders(ctx context.Context, req *models.GetAllOrdersReq) (*orderproduct.GetAllOrdersReq, error) {
	return &orderproduct.GetAllOrdersReq{UserId: req.UserID}, nil
}
func (u *AdjustRequest) OrderCompleted(ctx context.Context, req *models.UpdateOrderReq) (*orderproduct.UpdateOrderReq, error) {
	return &orderproduct.UpdateOrderReq{OrderId: req.OrderID}, nil
}
func (u *AdjustRequest) AddProduct2Cart(ctx context.Context, req *models.AddProducts2Cart) (*orderproduct.AddProducts2Cart, error) {
	price, err := u.ProductClient(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}

	return &orderproduct.AddProducts2Cart{UserId: req.UserID, ProductId: req.ProductID, Price: price, Quantity: req.Quantity}, nil
}
func (u *AdjustRequest) GetCart(ctx context.Context, req *models.GetCartReq) (*orderproduct.GetCartReq, error) {
	return &orderproduct.GetCartReq{UserId: req.UserID}, nil
}
func (u *AdjustRequest) UpdateCart(ctx context.Context, req *models.UpdateCartReq) (*orderproduct.UpdateCartReq, error) {
	return &orderproduct.UpdateCartReq{UserId: req.UserID, ProductId: req.ProductID, Quantity: req.Quantity}, nil
}
func (u *AdjustRequest) DeleteCart(ctx context.Context, req *models.GetCartReq) (*orderproduct.GetCartReq, error) {
	return &orderproduct.GetCartReq{UserId: req.UserID}, nil
}
func (u *AdjustRequest) DeleteProductsFromCart(ctx context.Context, req *models.DeleteProductsfromCartReq) (*orderproduct.DeleteProductsfromCartReq, error) {
	return &orderproduct.DeleteProductsfromCartReq{UserId: req.UserID, ProductId: req.ProductID}, nil
}

func (u *AdjustRequest) ProductClient(ctx context.Context, id string) (float32, error) {
	res, err := u.P.GetAllProduct(ctx, &productpb.GetProductsReq{Field: "productId", Value: id})
	if err != nil {
		return 0, err
	}
	for _, i := range res.Product {
		return i.Price, nil
	}
	return 0, nil
}
