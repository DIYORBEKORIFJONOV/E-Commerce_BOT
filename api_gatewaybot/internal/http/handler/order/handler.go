package orderhandler

import (
	"context"
	interface17 "ecommercebot/internal/interface"
	models "ecommercebot/internal/model"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	U interface17.OrderService
	C context.Context
}

func (u *OrderHandler) CreateOrder(c *gin.Context) {
	var req models.CreateOrderReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(403, err)
		return
	}
	res, err := u.U.CreateOrder(u.C, &req)
	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(201, res)
}

func (u *OrderHandler) GetOrders(c *gin.Context) {
	var req models.GetAllOrdersReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(403, err)
		return
	}
	res, err := u.U.GetOrders(u.C, &req)
	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(201, res)
}

func (u *OrderHandler) OrderCompleted(c *gin.Context) {
	var req models.UpdateOrderReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(403, err)
		return
	}
	res, err := u.U.OrderCompleted(u.C, &req)
	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(201, res)
}

func (u *OrderHandler) AddProduct2Cart(c *gin.Context) {
	var req models.AddProducts2Cart
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(403, err)
		return
	}
	res, err := u.U.AddProduct2Cart(u.C, &req)
	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(201, res)
}

func (u *OrderHandler) GetCart(c *gin.Context) {
	var req models.GetCartReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(403, err)
		return
	}
	res, err := u.U.GetCart(u.C, &req)
	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(201, res)
}

func (u *OrderHandler) UpdateCart(c *gin.Context) {
	var req models.UpdateCartReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(403, err)
		return
	}
	res, err := u.U.UpdateCart(u.C, &req)
	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(201, res)
}

func (u *OrderHandler) DeleteCart(c *gin.Context) {
	var req models.GetCartReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(403, err)
		return
	}
	res, err := u.U.DeleteCart(u.C, &req)
	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(201, res)
}

func (u *OrderHandler) DeletDeleteProductsFromCarteCart(c *gin.Context) {
	var req models.DeleteProductsfromCartReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(403, err)
		return
	}
	res, err := u.U.DeleteProductsFromCart(u.C, &req)
	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(201, res)
}
