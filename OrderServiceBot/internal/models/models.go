package models

type Location struct {
	Longitude string `json:"longtitude" bson:"longtitude"`
	Latitude  string `json:"latitude" bson:"latitude"`
}

type ProductOrder struct {
	ProductID string  `json:"productId" bson:"productId"`
	Quantity  int64   `json:"quantity" bson:"quantity"`
	Price     float32 `json:"price" bson:"price"`
}

type Cart struct {
	Products   []ProductOrder `json:"products" bson:"products"`
	TotalPrice float32        `json:"totalPrice" bson:"totalPrice"`
	UserID     string         `json:"userId" bson:"userId"`
}

type Order struct {
	OrderID       string   `json:"orderId" bson:"orderId"`
	CartID        Cart     `json:"cartId" bson:"cartId"`
	Coordination  Location `json:"cordination" bson:"cordination"`
	Status        string   `json:"status" bson:"status"`
	UserID        string   `json:"userId" bson:"userId"`
	Comment       string   `json:"comment" bson:"comment"`
	ContactNumber string   `json:"contactNumber" bson:"contactNumber"`
	CreatedAt     string   `json:"createdat" bson:"createdat"`
	UpdatedAt     string   `json:"updatedat" bson:"updated"`
}

type CreateOrderReq struct {
	CartID        Cart     `json:"cartId" bson:"cartId"`
	Coordination  Location `json:"cordination" bson:"cordination"`
	UserID        string   `json:"userId" bson:"userId"`
	Comment       string   `json:"comment" bson:"comment"`
	ContactNumber string   `json:"contactNumber" bson:"contactNumber"`
}

type GetAllOrdersReq struct {
	UserID string `json:"userId" bson:"userId"`
}

type GetAllOrdersRes struct {
	Orders []Order `json:"orders" bson:"orders"`
}

type UpdateOrderReq struct {
	OrderID string `json:"orderId" bson:"orderId"`
}

type GeneralOrderResponse struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
}

type AddProducts2Cart struct {
	UserID    string  `json:"userId" bson:"userId"`
	ProductID string  `json:"productId" bson:"productId"`
	Quantity  int64   `json:"quantity" bson:"quantity"`
	Price     float32 `json:"price" bson:"price"`
}

type GetCartReq struct {
	UserID string `json:"userId" bson:"userId"`
}

type UpdateCartReq struct {
	UserID    string `json:"userId" bson:"userId"`
	ProductID string `json:"productId" bson:"productId"`
	Quantity  int64  `json:"quantity" bson:"quantity"`
}

type DeleteProductsfromCartReq struct {
	UserID    string `json:"userId" bson:"userId"`
	ProductID string `json:"productId" bson:"productId"`
}
