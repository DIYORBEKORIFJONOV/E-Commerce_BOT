package adjust

import (
	"context"
	"fmt"
	"time"

	interface17 "github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/interface"
	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/models"
	"github.com/google/uuid"
)

type Centralized struct {
	M interface17.Order
}

func (u *Centralized) CreateOrder(ctx context.Context, req *models.CreateOrderReq) (*models.Order, error) {
	idorder := uuid.NewString()
	var order = models.Order{
		OrderID:       idorder,
		CartID:        req.CartID,
		Coordination:  req.Coordination,
		Status:        "In Progress",
		UserID:        req.UserID,
		Comment:       req.Comment,
		ContactNumber: req.ContactNumber,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := u.M.CreateOrder(ctx, &order); err != nil {
		return nil, err
	}
	return &order, nil
}
func (u *Centralized) GeOrders(ctx context.Context, req *models.GetAllOrdersReq) (*models.GetAllOrdersRes, error) {
	res, err := u.M.GetOrders(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Centralized) OrderCompleted(ctx context.Context, req *models.UpdateOrderReq) (*models.Order, error) {
	if err := u.M.OrderCompleted(ctx, req); err != nil {
		return nil, err
	}
	res, err := u.M.GetOrder(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Centralized) AddProduct2Cart(ctx context.Context, req *models.AddProducts2Cart) (*models.GeneralOrderResponse, error) {
	if err := u.M.AddProduct2Cart(ctx, req); err != nil {
		return nil, err
	}

	return &models.GeneralOrderResponse{Status: true, Message: "Product is succesfully added to Cart"}, nil
}

func (u *Centralized) GetCart(ctx context.Context, req *models.GetCartReq) (*models.Cart, error) {
	cursor, err := u.M.GetCart(ctx, req)
	if err != nil {
		return nil, err
	}
	var products []models.AddProducts2Cart
	if err := cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("failed to decode cart items: %w", err)
	}
	var cart models.Cart
	var total float32

	for _, p := range products {
		cart.Products = append(cart.Products, models.ProductOrder{
			ProductID: p.ProductID,
			Quantity:  p.Quantity,
			Price:     p.Price,
		})
		total += float32(p.Quantity) * p.Price
	}

	cart.TotalPrice = total
	cart.UserID = req.UserID
	return &cart, nil
}
func (u *Centralized) UpdateCart(ctx context.Context, req *models.UpdateCartReq) (*models.Cart, error) {
	if err := u.M.UpdateCart(ctx, req); err != nil {
		return nil, fmt.Errorf("updating product in the cart error %w", err)
	}
	cursor, err := u.M.GetCart(ctx, &models.GetCartReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	var products []models.AddProducts2Cart
	if err := cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("failed to decode cart items: %w", err)
	}
	var cart models.Cart
	var total float32

	for _, p := range products {
		cart.Products = append(cart.Products, models.ProductOrder{
			ProductID: p.ProductID,
			Quantity:  p.Quantity,
			Price:     p.Price,
		})
		total += float32(p.Quantity) * p.Price
	}

	cart.TotalPrice = total
	cart.UserID = req.UserID
	return &cart, nil
}
func (u *Centralized) DeleteCart(ctx context.Context, req *models.GetCartReq) (*models.GeneralOrderResponse, error) {
	if err := u.M.DeleteCart(ctx, req); err != nil {
		return nil, fmt.Errorf("deleting cart error %w", err)
	}
	return &models.GeneralOrderResponse{Status: true, Message: "succesfully deleted"}, nil
}
func (u *Centralized) DeleteProductsFromCart(ctx context.Context, req *models.DeleteProductsfromCartReq) (*models.Cart, error) {
	if err := u.M.DeleteProductsFromCart(ctx, req); err != nil {
		return nil, fmt.Errorf("updating product in the cart error %w", err)
	}
	cursor, err := u.M.GetCart(ctx, &models.GetCartReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	var products []models.AddProducts2Cart
	if err := cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("failed to decode cart items: %w", err)
	}
	var cart models.Cart
	var total float32

	for _, p := range products {
		cart.Products = append(cart.Products, models.ProductOrder{
			ProductID: p.ProductID,
			Quantity:  p.Quantity,
			Price:     p.Price,
		})
		total += float32(p.Quantity) * p.Price
	}

	cart.TotalPrice = total
	cart.UserID = req.UserID
	return &cart, nil
}
