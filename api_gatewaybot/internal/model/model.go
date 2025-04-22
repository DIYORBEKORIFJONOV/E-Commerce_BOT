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

type GetMainProductReq struct {
    Field string `json:"field" bson:"field"`
    Value string `json:"value" bson:"value"`
}

type MainProduct struct {
    MainProductID string `json:"mainProductId" bson:"mainProductId"`
    Category      string `json:"category" bson:"category"`
    Name          string `json:"name" bson:"name"`
}

type GetMainProductRes struct {
    Products []MainProduct `json:"products" bson:"products"`
    Count    int64         `json:"count" bson:"count"`
}

type UpdateNameReq struct {
    Name    string `json:"name" bson:"name"`
    NewName string `json:"newname" bson:"newname"`
}

type Product struct {
    ProductID     string  `json:"productId" bson:"productId"`
    Description   string  `json:"description" bson:"description"`
    Colour        string  `json:"colour" bson:"colour"`
    Size          int32   `json:"size" bson:"size"`
    Price         float32 `json:"price" bson:"price"`
    Quantity      int32   `json:"quantity" bson:"quantity"`
    CreatedAt     string  `json:"createdat" bson:"createdat"`
    UpdatedAt     string  `json:"updatedat" bson:"updatedat"`
    PhotoURL      string  `json:"photourl" bson:"photourl"`
    MainProductID string  `json:"mainProductId" bson:"mainProductId"`
}

type CreateProductReq struct {
    Name     string `json:"name" bson:"name"`
    Category string `json:"category" bson:"category"`
}

type AddModelReq struct {
    MainProductID string  `json:"mainProductId" bson:"mainProductId"`
    Description   string  `json:"description" bson:"description"`
    Colour        string  `json:"colour" bson:"colour"`
    Size          int32   `json:"size" bson:"size"`
    Price         float32 `json:"price" bson:"price"`
    Quantity      int32   `json:"quantity" bson:"quantity"`
    PhotoURL      string  `json:"photourl" bson:"photourl"`
}

type GetProductsReq struct {
    Field string `json:"field" bson:"field"`
    Value string `json:"value" bson:"value"`
    Page  int64  `json:"page" bson:"page"`
    Limit int64  `json:"limit" bson:"limit"`
}

type GetProductsRes struct {
    Product []Product `json:"product" bson:"product"`
    Count   int64     `json:"count" bson:"count"`
}

type UpdateProductReq struct {
    ProductID   string  `json:"productId" bson:"productId"`
    Description string  `json:"description" bson:"description"`
    Colour      string  `json:"colour" bson:"colour"`
    Size        int32   `json:"size" bson:"size"`
    Price       float32 `json:"price" bson:"price"`
    Quantity    int32   `json:"quantity" bson:"quantity"`
    PhotoURL    string  `json:"photoUrl" bson:"photoUrl"`
}

type DeleteProductReq struct {
    ProductID string `json:"productId" bson:"productId"`
    IsDeleted bool   `json:"is_deleted" bson:"is_deleted"`
}

type GeneralResponseProduct struct {
    Status  bool   `json:"status" bson:"status"`
    Message string `json:"message" bson:"message"`
}

type CreateCategoryReq struct {
    Category string `json:"category" bson:"category"`
}

type GetCategoriesReq struct{}

type GetcategoriesRes struct {
    Category []CreateCategoryReq `json:"category" bson:"category"`
    Count    int64               `json:"count" bson:"count"`
}

type UpdateCategoryReq struct {
    Category    string `json:"category" bson:"category"`
    NewCategory string `json:"newcategory" bson:"newcategory"`
}

type DeleteCategoryReq struct {
    Category string `json:"category" bson:"category"`
}
