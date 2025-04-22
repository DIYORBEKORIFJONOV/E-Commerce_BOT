package adjustresponse

import (
	"context"
	models "ecommercebot/internal/model"
	"ecommercebot/internal/protos/orderproduct"
)

type AdjustResponse struct {
	C context.Context
}

func (u *AdjustResponse) CreateOrder(ctx context.Context, res *orderproduct.Order) (*models.Order, error) {
	var c []models.ProductOrder
	t := 0

	for _, i := range res.CartId.Products {
		var product = models.ProductOrder{
			ProductID: i.ProductId,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, product)
	}
	var finalcart = models.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserID:     res.UserId,
	}
	var l = models.Location{
		Longitude: res.Cordination.Longtitude,
		Latitude:  res.Cordination.Latitude,
	}

	return &models.Order{
		OrderID:       res.OrderId,
		CartID:        finalcart,
		Coordination:  l,
		Status:        res.Status,
		UserID:        res.UserId,
		Comment:       res.Comment,
		ContactNumber: res.ContactNumber,
		CreatedAt:     res.Createdat,
		UpdatedAt:     res.Updated,
	}, nil
}
func (u *AdjustResponse) GetOrders(ctx context.Context, res *orderproduct.GetAllOrdersRes) (*models.GetAllOrdersRes, error) {
	var orders []models.Order
	for _, i := range res.Orders {
		var c []models.ProductOrder
		t := 0

		for _, j := range i.CartId.Products {
			var product = models.ProductOrder{
				ProductID: j.ProductId,
				Quantity:  j.Quantity,
				Price:     j.Price,
			}
			t += int(j.Quantity) * int(j.Price)
			c = append(c, product)
		}
		var finalcart = models.Cart{
			Products:   c,
			TotalPrice: float32(t),
			UserID:     i.UserId,
		}
		var l = models.Location{
			Longitude: i.Cordination.Longtitude,
			Latitude:  i.Cordination.Latitude,
		}
		var order = models.Order{
			OrderID:       i.OrderId,
			CartID:        finalcart,
			Coordination:  l,
			Status:        i.Status,
			UserID:        i.UserId,
			Comment:       i.Comment,
			ContactNumber: i.ContactNumber,
			CreatedAt:     i.Createdat,
			UpdatedAt:     i.Updated,
		}
		orders = append(orders, order)
	}

	return &models.GetAllOrdersRes{Orders: orders}, nil
}
func (u *AdjustResponse) OrderCompleted(ctx context.Context, res *orderproduct.Order) (*models.Order, error) {
	var c []models.ProductOrder
	t := 0

	for _, i := range res.CartId.Products {
		var product = models.ProductOrder{
			ProductID: i.ProductId,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, product)
	}
	var finalcart = models.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserID:     res.UserId,
	}
	var l = models.Location{
		Longitude: res.Cordination.Longtitude,
		Latitude:  res.Cordination.Latitude,
	}

	return &models.Order{
		OrderID:       res.OrderId,
		CartID:        finalcart,
		Coordination:  l,
		Status:        res.Status,
		UserID:        res.UserId,
		Comment:       res.Comment,
		ContactNumber: res.ContactNumber,
		CreatedAt:     res.Createdat,
		UpdatedAt:     res.Updated,
	}, nil
}
func (u *AdjustResponse) AddProduct2Cart(ctx context.Context, res *orderproduct.GeneralOrderResponse) (*models.GeneralOrderResponse, error) {
	return &models.GeneralOrderResponse{Status: res.Status, Message: res.Message}, nil
}
func (u *AdjustResponse) GetCart(ctx context.Context, res *orderproduct.Cart) (*models.Cart, error) {
	var c []models.ProductOrder
	t := 0

	for _, i := range res.Products {
		var product = models.ProductOrder{
			ProductID: i.ProductId,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, product)
	}
	var finalcart = models.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserID:     res.UserId,
	}
	return &finalcart, nil
}
func (u *AdjustResponse) UpdateCart(ctx context.Context, res *orderproduct.Cart) (*models.Cart, error) {
	var c []models.ProductOrder
	t := 0

	for _, i := range res.Products {
		var product = models.ProductOrder{
			ProductID: i.ProductId,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, product)
	}
	var finalcart = models.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserID:     res.UserId,
	}
	return &finalcart, nil
}
func (u *AdjustResponse) DeleteCart(ctx context.Context, res *orderproduct.GeneralOrderResponse) (*models.GeneralOrderResponse, error) {
	return &models.GeneralOrderResponse{Status: res.Status, Message: res.Message}, nil
}
func (u *AdjustResponse) DeleteProductsFromCart(ctx context.Context, res *orderproduct.Cart) (*models.Cart, error) {
	var c []models.ProductOrder
	t := 0

	for _, i := range res.Products {
		var product = models.ProductOrder{
			ProductID: i.ProductId,
			Quantity:  i.Quantity,
			Price:     i.Price,
		}
		t += int(i.Quantity) * int(i.Price)
		c = append(c, product)
	}
	var finalcart = models.Cart{
		Products:   c,
		TotalPrice: float32(t),
		UserID:     res.UserId,
	}
	return &finalcart, nil
}
