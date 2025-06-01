// file: internal/http/handler/order.go
package handler

import (
	_ "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/app/docs"
	models "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/order"
	usecaseorder "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	U *usecaseorder.OrderUseCaseIml
}

func NewOrderHandler(u *usecaseorder.OrderUseCaseIml) *OrderHandler {
	return &OrderHandler{U: u}
}

// @title Artisan Connect
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host hurmomarketshop.duckdns.org
// @BasePath        /
// @schemes         https
// @securityDefinitions.apiKey ApiKeyAuth
// @in              header
// @name            Authorization

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Создаёт новый заказ
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        order body models.CreateOrderReq true "Order payload"
// @Security ApiKeyAuth
// @Success      201   {object} models.Order         "Created"
// @Failure      403   {object} models.ErrorResponse "Bad Request"
// @Router       /orders/create [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req models.CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.U.CreateOrder(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetOrders godoc
// @Summary      Retrieve all orders
// @Description  Возвращает список заказов
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        userId query string false "Filter by customer ID"
// @Security ApiKeyAuth
// @Success      200    {object} models.GetAllOrdersRes "List"
// @Failure      400    {object} models.ErrorResponse    "Bad Request"
// @Failure      500    {object} models.ErrorResponse    "Internal Server Error"
// @Router       /orders/getall [get]
func (h *OrderHandler) GetOrders(c *gin.Context) {
	var req models.GetAllOrdersReq
	req.UserID = c.Query("userId")
	res, err := h.U.GetOrders(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// OrderCompleted godoc
// @Summary      Mark an order as completed
// @Description  Помечает заказ как завершённый
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        status body models.UpdateOrderReq true "Update payload"
// @Security ApiKeyAuth
// @Success      200    {object} models.Order         "Updated"
// @Failure      403    {object} models.ErrorResponse "Bad Request"
// @Router       /orders/completed [put]
func (h *OrderHandler) OrderCompleted(c *gin.Context) {
	var req models.UpdateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.U.OrderCompleted(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddProduct2Cart godoc
// @Summary      Add products to cart
// @Description  Добавляет товары в корзину
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        items body models.AddProducts2Cart true "Items payload"
// @Security ApiKeyAuth
// @Success      200   {object} models.GeneralOrderResponse "Updated"
// @Failure      403   {object} models.ErrorResponse         "Bad Request"
// @Router       /carts/add/product [post]
func (h *OrderHandler) AddProduct2Cart(c *gin.Context) {
	var req models.AddProducts2Cart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.U.AddProduct2Cart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetCart godoc
// @Summary      Get cart details
// @Description  Получает содержимое корзины
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        userId query string true "User ID"
// @Security ApiKeyAuth
// @Success      200    {object} models.Cart         "Contents"
// @Failure      400    {object} models.ErrorResponse "Bad Request"
// @Failure      500    {object} models.ErrorResponse "Internal Server Error"
// @Router       /carts [get]
func (h *OrderHandler) GetCart(c *gin.Context) {
	var req models.GetCartReq
	req.UserID = c.Query("userId")
	res, err := h.U.GetCart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateCart godoc
// @Summary      Update cart contents
// @Description  Обновляет корзину
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        update body models.UpdateCartReq true "Update payload"
// @Success      200    {object} models.Cart         "Updated"
// @Security ApiKeyAuth
// @Failure      403    {object} models.ErrorResponse "Bad Request"
// @Router       /carts [put]
func (h *OrderHandler) UpdateCart(c *gin.Context) {
	var req models.UpdateCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.U.UpdateCart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteCart godoc
// @Summary      Delete a user's cart
// @Description  Очищает корзину
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        user body models.GetCartReq true "User payload"
// @Success      200  {object} models.GeneralOrderResponse "Deleted"
// @Security ApiKeyAuth
// @Failure      403  {object} models.ErrorResponse         "Bad Request"
// @Router       /carts [delete]
func (h *OrderHandler) DeleteCart(c *gin.Context) {
	var req models.GetCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.U.DeleteCart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteProductsFromCart godoc
// @Summary      Remove specific products from cart
// @Description  Удаляет указанные товары из корзины
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        items body models.DeleteProductsfromCartReq true "Items payload"
// @Security ApiKeyAuth
// @Success      200   {object} models.Cart         "Updated"
// @Failure      403   {object} models.ErrorResponse "Bad Request"
// @Router       /carts/product [delete]
func (h *OrderHandler) DeleteProductsFromCart(c *gin.Context) {
	var req models.DeleteProductsfromCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.U.DeleteProductsFromCart(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusForbidden, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
