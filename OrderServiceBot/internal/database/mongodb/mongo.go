package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	D *mongo.Collection //Order collection
	B *mongo.Collection // Cart collection
	C context.Context
}

func (u *Mongo) CreateOrder(ctx context.Context, order *models.Order) error {
	_, err := u.D.InsertOne(ctx, order)
	if err != nil {
		return fmt.Errorf("error while saving order: %v", err)
	}
	return nil
}

func (u *Mongo) GetOrders(ctx context.Context, order *models.GetAllOrdersReq) (*models.GetAllOrdersRes, error) {
	cursor, err := u.D.Find(ctx, bson.M{"userId": order.UserID})
	if err != nil {
		return nil, fmt.Errorf("error while getting all orders: %v", err)
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, fmt.Errorf("error decoding orders: %v", err)
	}

	return &models.GetAllOrdersRes{Orders: orders}, nil
}

func (u *Mongo) OrderCompleted(ctx context.Context, order *models.UpdateOrderReq) error {
	_, err := u.D.UpdateOne(ctx,
		bson.M{"orderId": order.OrderID},
		bson.M{"$set": bson.M{"status": "completed", "updatedat": time.Now().Format("2006-01-02 15:04:05")}})

	if err != nil {
		return fmt.Errorf("error while changing order status: %v", err)
	}
	return nil
}

func (u *Mongo) AddProduct2Cart(ctx context.Context, order *models.AddProducts2Cart) error {
	_, err := u.B.InsertOne(ctx, order)
	if err != nil {
		return fmt.Errorf("error while adding product to user's cart: %v", err)
	}
	return nil
}

func (u *Mongo) GetCart(ctx context.Context, order *models.GetCartReq) (*mongo.Cursor, error) {
	cursor, err := u.B.Find(ctx, bson.M{"userId": order.UserID})
	if err != nil {
		return nil, fmt.Errorf("error getting user's cart: %v", err)
	}
	return cursor, nil
}

func (u *Mongo) UpdateCart(ctx context.Context, req *models.UpdateCartReq) error {
	_, err := u.B.UpdateOne(ctx, bson.M{"userId": req.UserID, "productId": req.ProductID}, bson.M{"$set": bson.M{"quantity": req.Quantity, "updatedat": time.Now().Format("2006-01-02 15:04:05")}})
	if err != nil {
		return fmt.Errorf("updating cart product %v", err)
	}
	return nil
}

func (u *Mongo) DeleteCart(ctx context.Context, order *models.GetCartReq) error {
	_, err := u.B.DeleteOne(ctx, bson.M{"userId": order.UserID})
	if err != nil {
		return fmt.Errorf("error deleting cart: %v", err)
	}
	return nil
}
func (u *Mongo) DeleteProductsFromCart(ctx context.Context, req *models.DeleteProductsfromCartReq) error {
	// , "updatedat": time.Now().Format("2006-01-02 15:04:05")
	_, err := u.B.DeleteOne(ctx, bson.M{"userId": req.UserID, "productId": req.ProductID})
	if err != nil {
		return fmt.Errorf("error deleting product from cart: %v", err)
	}
	return nil
}

func (u *Mongo) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	if err := u.D.FindOne(ctx, bson.M{"orderId": id}).Decode(&order); err != nil {
		return nil, err
	}
	return &order, nil
}
