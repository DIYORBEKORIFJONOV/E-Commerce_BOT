package handler

import (
	"log"
	"net/http"
	"strconv"

	productentity "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/product"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/minao"
	productusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/product"

	"github.com/gin-gonic/gin"
)

// ProductHandler хранит зависимости для product-эндпоинтов
type ProductHandler struct {
	product    *productusecase.ProductUseCaseIml
	minioPhoto *minao1.FileStorage
}

func NewProductHandler(
	productUC *productusecase.ProductUseCaseIml,
	minioClient *minao1.FileStorage,
) *ProductHandler {
	return &ProductHandler{
		product:    productUC,
		minioPhoto: minioClient,
	}
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Создаёт новый продукт на основе JSON-запроса
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      productentity.CreateProductReq  true  "Product payload"
// @Success      201      {object}  productentity.Product          "Created"
// @Failure      400      {object}  productentity.ErrorResponse    "Bad Request"
// @Failure      500      {object}  productentity.ErrorResponse    "Internal Server Error"
// @Router       /products [post]
func (p *ProductHandler) CreateProduct(c *gin.Context) {
	var req productentity.CreateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	mainProduct, err := p.product.CreateProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mainProduct)
}

// AddModel godoc
// @Summary      Добавляет модель товара с фото
// @Description  Принимает form-data: все поля AddModelReq + файл photo
// @Tags         products
// @Accept       multipart/form-data
// @Produce      json
// @Param        mainproductId formData string true  "MainProductId"
// @Param        description    formData string false "Description"
// @Param        colour         formData string false "Colour"
// @Param        size           formData int    false "Size"
// @Param        price          formData number false "Price"
// @Param        quantity       formData int    false "Quantity"
// @Param        photo          formData file   true  "Photo file"
// @Success      201 {object} productentity.Product  "Созданный товар"
// @Failure      400 {object} productentity.ErrorResponse          "Некорректный запрос"
// @Failure      500 {object} productentity.ErrorResponse          "Ошибка сервера"
// @Router       /products/addmodel [post]
func (p *ProductHandler) AddModel(c *gin.Context) {
	var req productentity.AddModelReq

	req.MainProductId = c.PostForm("mainproductId")
	req.Description = c.PostForm("description")
	req.Colour = c.PostForm("colour")

	if sizeStr := c.PostForm("size"); sizeStr != "" {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: "invalid size"})
			return
		}
		req.Size = int32(size)
	}

	if priceStr := c.PostForm("price"); priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: "invalid price"})
			return
		}
		req.Price = float32(price)
	}

	if qtyStr := c.PostForm("quantity"); qtyStr != "" {
		qty, err := strconv.Atoi(qtyStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: "invalid quantity"})
			return
		}
		req.Quantity = int32(qty)
	}

	if p.minioPhoto == nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: "MinIO client is not initialized"})
		return
	}

	fileHeader, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: "photo is required"})
		return
	}
	objectName, err := p.minioPhoto.UploadFile(fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	req.PhotoURL = objectName

	product, err := p.product.AddModel(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// UpdateProductName godoc
// @Summary      Update product name
// @Description  Обновляет название продукта по ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.UpdateNameReq         true  "Update payload"
// @Success      200      {object}  productentity.GeneralResponseProduct "Updated"
// @Failure      400      {object}  productentity.ErrorResponse         "Bad Request"
// @Failure      500      {object}  productentity.ErrorResponse         "Internal Server Error"
// @Router       /products/name [put]
func (p *ProductHandler) UpdateProductName(c *gin.Context) {
	var req productentity.UpdateNameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	generalRes, err := p.product.UpdateName(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, generalRes)
}

// GetMainProduct godoc
// @Summary      Get list of main products
// @Description  Получает список основных продуктов. Поддерживает фильтрацию и пагинацию.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        field  query     string  false  "Filter/sort field"  example(name)
// @Param        value   query     string     false  "Page number"        example(1)
// @Success      200    {array}   productentity.ProductMain
// @Failure      500    {object}  productentity.ErrorResponse
// @Router       /products/main [get]
func (p *ProductHandler) GetMainProduct(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	mainProducts, err := p.product.GetMainProduct(c.Request.Context(), field, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, mainProducts)
}

// GetAllProduct godoc
// @Summary      Список продуктов
// @Description  Получить список продуктов с фильтрацией и пагинацией
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        field   query     string  false  "Поле для фильтра (name, category)"
// @Param        value   query     string  false  "Значение для фильтра"
// @Param        page    query     int     false  "Номер страницы"   default(1)
// @Param        limit   query     int     false  "Размер страницы"  default(10)
// @Success      200     {object}  productentity.GetProductsRes   "Список продуктов"
// @Failure      400     {object}  productentity.ErrorResponse     "Неверный запрос"
// @Failure      500     {object}  productentity.ErrorResponse     "Внутренняя ошибка сервера"
// @Router       /products [get]
func (p *ProductHandler) GetAllProduct(c *gin.Context) {
	var req productentity.GetProductsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := p.product.GetAllProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}

	if len(resp.Products) == 0 {
		c.JSON(http.StatusOK, resp)
		return
	}

	if p.minioPhoto == nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	for i, prod := range resp.Products {
		log.Println(prod.PhotoURL)
		signedURL, err := p.minioPhoto.GetFile(prod.PhotoURL)
		if err != nil {
			c.Error(err)
			signedURL = ""
		}
		resp.Products[i].PhotoURL = signedURL
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateProduct godoc
// @Summary      Partially update product
// @Description  Обновляет продукт (multipart/form-data). Фото обновляется, если есть.
// @Tags         products
// @Accept       multipart/form-data
// @Produce      json
// @Param        id           formData string true  "Product ID"
// @Param        name         formData string false "Name"
// @Param        category     formData string false "Category"
// @Param        description  formData string false "Description"
// @Param        photo        formData file   false "Photo file"
// @Success      200          {object} productentity.Product        "Updated"
// @Failure      400          {object} productentity.ErrorResponse  "Bad Request"
// @Failure      500          {object} productentity.ErrorResponse  "Internal Server Error"
// @Router       /products [patch]
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	var req productentity.UpdateProductReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}

	if fileHeader, err := c.FormFile("photo"); err == nil {
		if p.minioPhoto == nil {
			c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: "MinIO client is not initialized"})
			return
		}
		objectName, err := p.minioPhoto.UploadFile(fileHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
			return
		}
		req.PhotoURL = objectName
	}

	product, err := p.product.UpdateProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary      Delete an existing product
// @Description  Удаляет продукт по ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.DeleteProductReq      true  "Delete payload"
// @Success      200      {object}  productentity.GeneralResponseProduct "Deleted"
// @Failure      400      {object}  productentity.ErrorResponse          "Bad Request"
// @Failure      500      {object}  productentity.ErrorResponse          "Internal Server Error"
// @Router       /products [delete]
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	var req productentity.DeleteProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	res, err := p.product.DeleteProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Создаёт категорию
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.CreateCategoryReq   true  "Category payload"
// @Success      200      {object}  productentity.GeneralResponseProduct "Created"
// @Failure      400      {object}  productentity.ErrorResponse          "Bad Request"
// @Failure      500      {object}  productentity.ErrorResponse          "Internal Server Error"
// @Router       /products/categories [post]
func (p *ProductHandler) CreateCategory(c *gin.Context) {
	var req productentity.CreateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	generalRes, err := p.product.CreateCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, generalRes)
}

// GetAllCategory godoc
// @Summary      Retrieve all categories
// @Description  Возвращает список всех категорий
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {array}   productentity.GetcategoriesRes    "Categories list"
// @Failure      500  {object}  productentity.ErrorResponse       "Internal Server Error"
// @Router       /products/categories [get]
func (p *ProductHandler) GetAllCategory(c *gin.Context) {
	categories, err := p.product.GetAllCategory(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// UpdateCategory godoc
// @Summary      Update an existing category
// @Description  Обновляет категорию
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.UpdateCategoryReq   true  "Update payload"
// @Success      200      {object}  productentity.GeneralResponseProduct "Updated"
// @Failure      400      {object}  productentity.ErrorResponse          "Bad Request"
// @Failure      500      {object}  productentity.ErrorResponse          "Internal Server Error"
// @Router       /products/categories [put]
func (p *ProductHandler) UpdateCategory(c *gin.Context) {
	var req productentity.UpdateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	generalRes, err := p.product.UpdateCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, generalRes)
}

// DeleteCategory godoc
// @Summary      Delete a category
// @Description  Удаляет категорию по ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.DeleteCategoryReq   true  "Delete payload"
// @Success      200      {object}  productentity.GeneralResponseProduct "Deleted"
// @Failure      400      {object}  productentity.ErrorResponse          "Bad Request"
// @Failure      500      {object}  productentity.ErrorResponse          "Internal Server Error"
// @Router       /products/categories [delete]
func (p *ProductHandler) DeleteCategory(c *gin.Context) {
	var req productentity.DeleteCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	generalRes, err := p.product.DeleteCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, generalRes)
}