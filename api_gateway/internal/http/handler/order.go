package handler

import (
	_ "api_gateway/internal/app/docs"
	models "api_gateway/internal/entity/order"
	usecaseorder "api_gateway/internal/usecase/order"
	"net/http"

	"github.com/gin-gonic/gin"
)


type OrderHandler struct {
	U usecaseorder.OrderUseCaseIml
}

func NewOrderHandler(u *usecaseorder.OrderUseCaseIml) *OrderHandler {
	return &OrderHandler{
		U: *u,
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description CreateOrder accepts order details and creates a new order in the system.
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body models.CreateOrderReq true "Order request payload"
// @Success 201 {object} models.Order "Order created successfully"
// @Failure 403 {object} models.ErrorResponse "Invalid request or creation failed"
// @Router /orders/create [post]
func (u *OrderHandler) CreateOrder(c *gin.Context) {
	var req models.CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := u.U.CreateOrder(c.Request.Context(), &req)
	if err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(201, res)
}
// GetOrders godoc
// @Summary Retrieve all orders
// @Description GetOrders returns a list of orders based on the provided filters.
// @Tags Orders
// @Accept json
// @Produce json
// @Param userId query string false "Filter by customer ID"
// @Success 200 {object} models.GetAllOrdersRes "List of orders"
// @Failure 400 {object} models.ErrorResponse "Invalid query parameters"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /orders/getall [get]
func (u *OrderHandler) GetOrders(c *gin.Context) {
    var req models.GetAllOrdersReq

	req.UserID = c.Query("userId")
	
    res, err := u.U.GetOrders(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, res)
}

// OrderCompleted godoc
// @Summary Mark an order as completed
// @Description OrderCompleted updates the status of an order to completed.
// @Tags Orders
// @Accept json
// @Produce json
// @Param status body models.UpdateOrderReq true "Order completion payload"
// @Success 200 {object} models.Order "Order marked completed"
// @Failure 403 {object} models.ErrorResponse "Invalid request or update failed"
// @Router /orders/completed [put]
func (u *OrderHandler) OrderCompleted(c *gin.Context) {
	var req models.UpdateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := u.U.OrderCompleted(c.Request.Context(), &req)
	if err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, res)
}

// AddProduct2Cart godoc
// @Summary Add products to cart
// @Description AddProduct2Cart adds one or more products to a user's shopping cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param items body models.AddProducts2Cart true "Products to add to cart"
// @Success 200 {object} models.GeneralOrderResponse "Updated cart contents"
// @Failure 403 {object} models.ErrorResponse "Invalid request or addition failed"
// @Router /carts/add/product [post]
func (u *OrderHandler) AddProduct2Cart(c *gin.Context) {
	var req models.AddProducts2Cart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := u.U.AddProduct2Cart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, res)
}

// GetCart godoc
// @Summary Get cart details
// @Description GetCart retrieves the contents of a user's cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param userId query string true "User ID"
// @Success 200 {object} models.Cart "Current cart contents"
// @Failure 400 {object} models.ErrorResponse "Invalid query parameters"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /carts [get]
func (u *OrderHandler) GetCart(c *gin.Context) {
    var req models.GetCartReq
	req.UserID= c.Query("userId")
    res, err := u.U.GetCart(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}

// UpdateCart godoc
// @Summary Update cart contents
// @Description UpdateCart modifies the quantities or items in a user's cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param update body models.UpdateCartReq true "Cart update payload"
// @Success 200 {object} models.Cart "Updated cart contents"
// @Failure 403 {object} models.ErrorResponse "Invalid request or update failed"
// @Router /carts [put]
func (u *OrderHandler) UpdateCart(c *gin.Context) {
	var req models.UpdateCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := u.U.UpdateCart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, res)
}

// DeleteCart godoc
// @Summary Delete a user's cart
// @Description DeleteCart clears all items from a user's cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param user body models.GetCartReq true "User cart request"
// @Success 200 {object} models.GeneralOrderResponse "Cart deleted successfully"
// @Failure 403 {object} models.ErrorResponse "Invalid request or deletion failed"
// @Router /carts [delete]
func (u *OrderHandler) DeleteCart(c *gin.Context) {
	var req models.GetCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := u.U.DeleteCart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, res)
}

// DeleteProductsFromCart godoc
// @Summary Remove specific products from cart
// @Description DeleteProductsFromCart removes specified products from a user's cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param items body models.DeleteProductsfromCartReq true "Products to remove from cart"
// @Success 200 {object} models.Cart "Updated cart contents"
// @Failure 403 {object} models.ErrorResponse "Invalid request or deletion failed"
// @Router /carts/product [delete]
func (u *OrderHandler) DeleteProductsFromCart(c *gin.Context) {
	var req models.DeleteProductsfromCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := u.U.DeleteProductsFromCart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(403, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, res)
}
