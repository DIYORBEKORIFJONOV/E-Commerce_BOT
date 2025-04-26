package adjust

import (
	"context"

	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/models"
	orderproduct "github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/protos"
)

type Adjust struct {
	CA *Centralized
	C  context.Context
}

func (u *Adjust) CreateOrder(ctx context.Context, req *orderproduct.CreateOrderReq) (*orderproduct.Order, error) {
	var cartproducts []models.ProductOrder
	totalprice := 0
	for _, i := range req.CartId.Products {
		var product = models.ProductOrder{
			ProductID: i.ProductId,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		totalprice += int(i.Quantity) * int(i.Price)
		cartproducts = append(cartproducts, product)
	}
	var Cart = models.Cart{
		Products:   cartproducts,
		TotalPrice: float32(totalprice),
		UserID:     req.UserId,
	}
	var location = models.Location{
		Longitude: req.Cordination.Longtitude,
		Latitude:  req.Cordination.Latitude,
	}
	var newreq = models.CreateOrderReq{
		CartID:        Cart,
		Coordination:  location,
		UserID:        req.UserId,
		Comment:       req.Comment,
		ContactNumber: req.ContactNumber,
	}
	res, err := u.CA.CreateOrder(ctx, &newreq)
	if err != nil {
		return nil, err
	}
	var c []*orderproduct.ProductOrder
	t := 0

	for _, i := range res.CartID.Products {
		var product = orderproduct.ProductOrder{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, &product)
	}
	var finalcart = orderproduct.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserId:     req.UserId,
	}
	var l = orderproduct.Location{
		Longtitude: res.Coordination.Longitude,
		Latitude:   res.Coordination.Latitude,
	}

	return &orderproduct.Order{
		OrderId:       res.OrderID,
		CartId:        &finalcart,
		Cordination:   &l,
		Status:        res.Status,
		UserId:        res.UserID,
		Comment:       res.Comment,
		ContactNumber: res.ContactNumber,
		Createdat:     res.CreatedAt,
		Updated:       res.UpdatedAt,
	}, nil
}
func (u *Adjust) GeOrders(ctx context.Context, req *orderproduct.GetAllOrdersReq) (*orderproduct.GetAllOrdersRes, error) {
	res, err := u.CA.GeOrders(ctx, &models.GetAllOrdersReq{UserID: req.UserId})
	if err != nil {
		return nil, err
	}
	var orders []*orderproduct.Order
	for _, i := range res.Orders {
		var c []*orderproduct.ProductOrder
		t := 0

		for _, j := range i.CartID.Products {
			var product = orderproduct.ProductOrder{
				ProductId: j.ProductID,
				Quantity:  j.Quantity,
				Price:     j.Price,
			}
			t += int(j.Quantity) * int(j.Price)
			c = append(c, &product)
		}
		var finalcart = orderproduct.Cart{
			Products:   c,
			TotalPrice: float32(t),
			UserId:     req.UserId,
		}
		var l = orderproduct.Location{
			Longtitude: i.Coordination.Longitude,
			Latitude:   i.Coordination.Latitude,
		}
		var order = orderproduct.Order{
			OrderId:       i.OrderID,
			CartId:        &finalcart,
			Cordination:   &l,
			Status:        i.Status,
			UserId:        i.UserID,
			Comment:       i.Comment,
			ContactNumber: i.ContactNumber,
			Createdat:     i.CreatedAt,
			Updated:       i.UpdatedAt,
		}
		orders = append(orders, &order)
	}
	return &orderproduct.GetAllOrdersRes{Orders: orders}, nil
}
func (u *Adjust) OrderCompleted(ctx context.Context, req *orderproduct.UpdateOrderReq) (*orderproduct.Order, error) {
	res, err := u.CA.OrderCompleted(ctx, &models.UpdateOrderReq{OrderID: req.OrderId})
	if err != nil {
		return nil, err
	}
	var c []*orderproduct.ProductOrder
	t := 0

	for _, i := range res.CartID.Products {
		var product = orderproduct.ProductOrder{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, &product)
	}
	var finalcart = orderproduct.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserId:     res.UserID,
	}
	var l = orderproduct.Location{
		Longtitude: res.Coordination.Longitude,
		Latitude:   res.Coordination.Latitude,
	}

	return &orderproduct.Order{
		OrderId:       res.OrderID,
		CartId:        &finalcart,
		Cordination:   &l,
		Status:        res.Status,
		UserId:        res.UserID,
		Comment:       res.Comment,
		ContactNumber: res.ContactNumber,
		Createdat:     res.CreatedAt,
		Updated:       res.UpdatedAt,
	}, nil
}
func (u *Adjust) AddProduct2Cart(ctx context.Context, req *orderproduct.AddProducts2Cart) (*orderproduct.GeneralOrderResponse, error) {
	res, err := u.CA.AddProduct2Cart(ctx, &models.AddProducts2Cart{UserID: req.UserId, ProductID: req.ProductId, Quantity: req.Quantity, Price: req.Price})
	if err != nil {
		return nil, err
	}
	return &orderproduct.GeneralOrderResponse{Status: res.Status, Message: res.Message}, nil
}
func (u *Adjust) GetCart(ctx context.Context, req *orderproduct.GetCartReq) (*orderproduct.Cart, error) {
	res, err := u.CA.GetCart(ctx, &models.GetCartReq{UserID: req.UserId})
	if err != nil {
		return nil, err
	}
	var c []*orderproduct.ProductOrder
	t := 0

	for _, i := range res.Products {
		var product = orderproduct.ProductOrder{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, &product)
	}
	var finalcart = orderproduct.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserId:     res.UserID,
	}
	return &finalcart, nil
}
func (u *Adjust) UpdateCart(ctx context.Context, req *orderproduct.UpdateCartReq) (*orderproduct.Cart, error) {
	res, err := u.CA.UpdateCart(ctx, &models.UpdateCartReq{UserID: req.UserId, ProductID: req.ProductId, Quantity: req.Quantity})
	if err != nil {
		return nil, err
	}
	var c []*orderproduct.ProductOrder
	t := 0

	for _, i := range res.Products {
		var product = orderproduct.ProductOrder{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, &product)
	}
	var finalcart = orderproduct.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserId:     res.UserID,
	}

	return &finalcart, nil
}
func (u *Adjust) DeleteCart(ctx context.Context, req *orderproduct.GetCartReq) (*orderproduct.GeneralOrderResponse, error) {
	res, err := u.CA.DeleteCart(ctx, &models.GetCartReq{UserID: req.UserId})
	if err != nil {
		return nil, err
	}
	return &orderproduct.GeneralOrderResponse{Status: res.Status, Message: res.Message}, nil
}
func (u *Adjust) DeleteProductsFromCart(ctx context.Context, req *orderproduct.DeleteProductsfromCartReq) (*orderproduct.Cart, error) {
	res, err := u.CA.DeleteProductsFromCart(ctx, &models.DeleteProductsfromCartReq{UserID: req.UserId, ProductID: req.ProductId})
	if err != nil {
		return nil, err
	}
	var c []*orderproduct.ProductOrder
	t := 0

	for _, i := range res.Products {
		var product = orderproduct.ProductOrder{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, &product)
	}
	var finalcart = orderproduct.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserId:     res.UserID,
	}
	return &finalcart, nil
}
